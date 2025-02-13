package main

import (
	"bytes"
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/YoungAgency/simplefix-go/storages/memory"

	simplefixgo "github.com/YoungAgency/simplefix-go"
	"github.com/YoungAgency/simplefix-go/fix"
	"github.com/YoungAgency/simplefix-go/fix/encoding"
	"github.com/YoungAgency/simplefix-go/session"
	"github.com/YoungAgency/simplefix-go/session/messages"
	fixgen "github.com/YoungAgency/simplefix-go/tests/fix44"
)

func mustConvToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}

// TODO: move boilerplate to generator.
var pseudoGeneratedOpts = session.Opts{
	MessageBuilders: session.MessageBuilders{
		HeaderBuilder:        fixgen.Header{}.New(),
		TrailerBuilder:       fixgen.Trailer{}.New(),
		LogonBuilder:         fixgen.Logon{}.New(),
		LogoutBuilder:        fixgen.Logout{}.New(),
		RejectBuilder:        fixgen.Reject{}.New(),
		HeartbeatBuilder:     fixgen.Heartbeat{}.New(),
		TestRequestBuilder:   fixgen.TestRequest{}.New(),
		ResendRequestBuilder: fixgen.ResendRequest{}.New(),
	},
	Tags: &messages.Tags{
		MsgType:         mustConvToInt(fixgen.FieldMsgType),
		MsgSeqNum:       mustConvToInt(fixgen.FieldMsgSeqNum),
		HeartBtInt:      mustConvToInt(fixgen.FieldHeartBtInt),
		EncryptedMethod: mustConvToInt(fixgen.FieldEncryptMethod),
	},
	AllowedEncryptedMethods: map[string]struct{}{
		fixgen.EnumEncryptMethodNoneother: {},
	},
	SessionErrorCodes: &messages.SessionErrorCodes{
		InvalidTagNumber:            mustConvToInt(fixgen.EnumSessionRejectReasonInvalidtagnumber),
		RequiredTagMissing:          mustConvToInt(fixgen.EnumSessionRejectReasonRequiredtagmissing),
		TagNotDefinedForMessageType: mustConvToInt(fixgen.EnumSessionRejectReasonTagNotDefinedForThisMessageType),
		UndefinedTag:                mustConvToInt(fixgen.EnumSessionRejectReasonUndefinedtag),
		TagSpecialWithoutValue:      mustConvToInt(fixgen.EnumSessionRejectReasonTagspecifiedwithoutavalue),
		IncorrectValue:              mustConvToInt(fixgen.EnumSessionRejectReasonValueisincorrectoutofrangeforthistag),
		IncorrectDataFormatValue:    mustConvToInt(fixgen.EnumSessionRejectReasonIncorrectdataformatforvalue),
		DecryptionProblem:           mustConvToInt(fixgen.EnumSessionRejectReasonDecryptionproblem),
		SignatureProblem:            mustConvToInt(fixgen.EnumSessionRejectReasonSignatureproblem),
		CompIDProblem:               mustConvToInt(fixgen.EnumSessionRejectReasonCompidproblem),
		Other:                       mustConvToInt(fixgen.EnumSessionRejectReasonOther),
	},
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 9991))
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected")

	handlerFactory := simplefixgo.NewAcceptorHandlerFactory(fixgen.FieldMsgType, 10)

	exampleStorage := memory.NewStorage()

	server := simplefixgo.NewAcceptor(listener, handlerFactory, time.Second*5, func(handler simplefixgo.AcceptorHandler) {
		sess, err := session.NewAcceptorSession(
			&pseudoGeneratedOpts,
			handler,
			&session.LogonSettings{
				HeartBtInt:   30,
				LogonTimeout: time.Second * 30,
				HeartBtLimits: &session.IntLimits{
					Min: 5,
					Max: 60,
				},
			},
			func(request *session.LogonSettings) (err error) {
				fmt.Printf(
					"Logon passed for '%s' (%s)\n",
					request.Username,
					request.Password,
				)

				return nil
			},
			exampleStorage,
			exampleStorage,
		)
		if err != nil {
			panic(err)
		}

		_ = sess.Run()

		handler.HandleIncoming(simplefixgo.AllMsgTypes, func(msg []byte) bool {
			fmt.Println("incoming", string(bytes.ReplaceAll(msg, fix.Delimiter, []byte("|"))))
			return true
		})
		handler.HandleOutgoing(simplefixgo.AllMsgTypes, func(msg simplefixgo.SendingMessage) bool {
			data, err := msg.ToBytes()
			if err != nil {
				panic(err)
			}
			fmt.Println("outgoing", string(bytes.ReplaceAll(data, fix.Delimiter, []byte("|"))))
			return true
		})

		handler.HandleIncoming(fixgen.MsgTypeMarketDataRequest, func(msg []byte) bool {
			request := fixgen.NewMarketDataRequest()
			err := encoding.Unmarshal(request, msg)
			if err != nil {
				panic(err)
			}

			for _, relatedSymbolEntry := range request.RelatedSymGrp().Entries() {
				fmt.Println(relatedSymbolEntry.Instrument().Symbol())
			}

			return true
		})
	})

	panic(fmt.Errorf("the server was stopped: %s", server.ListenAndServe()))
}
