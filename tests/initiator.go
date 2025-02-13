package tests

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/YoungAgency/simplefix-go/storages/memory"

	simplefixgo "github.com/YoungAgency/simplefix-go"
	"github.com/YoungAgency/simplefix-go/fix"
	"github.com/YoungAgency/simplefix-go/session"
	fixgen "github.com/YoungAgency/simplefix-go/tests/fix44"
)

func RunNewInitiator(addr string, t *testing.T, settings *session.LogonSettings) (s *session.Session, handler *simplefixgo.DefaultHandler) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Fatalf("could not dial: %s", err)
	}

	handler = simplefixgo.NewInitiatorHandler(context.Background(), fixgen.FieldMsgType, 10)
	client := simplefixgo.NewInitiator(conn, handler, 10, time.Minute*50)

	testStorage := memory.NewStorage()

	s, err = session.NewInitiatorSession(
		handler,
		&pseudoGeneratedOpts,
		settings,
		testStorage,
		testStorage,
	)
	if err != nil {
		panic(err)
	}

	// logging messages:
	handler.HandleIncoming(simplefixgo.AllMsgTypes, func(msg []byte) bool {
		fmt.Println("incoming:", string(bytes.ReplaceAll(msg, fix.Delimiter, []byte("|"))))
		return true
	})
	handler.HandleOutgoing(simplefixgo.AllMsgTypes, func(msg simplefixgo.SendingMessage) bool {
		data, err := msg.ToBytes()
		if err != nil {
			panic(err)
		}
		fmt.Println("outgoing:", string(bytes.ReplaceAll(data, fix.Delimiter, []byte("|"))))
		return true
	})

	// todo move
	go func() {
		time.Sleep(time.Second * 10)
		fmt.Println("resending the request after 10 seconds")
		err := s.Send(fixgen.ResendRequest{}.New().SetFieldBeginSeqNo(2).SetFieldEndSeqNo(3))
		if err != nil {
			panic(err)
		}
	}()

	err = s.Run()
	if err != nil {
		t.Fatalf("could not run the session: %s", err)
	}

	go func() {
		err := client.Serve()
		if err != nil && !errors.Is(err, simplefixgo.ErrConnClosed) {
			panic(fmt.Errorf("could not serve the client: %s", err))
		}
	}()

	return s, handler
}
