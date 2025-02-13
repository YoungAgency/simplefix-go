// Code generated by fixgen. DO NOT EDIT.

package fix44

import (
	"github.com/YoungAgency/simplefix-go/fix"
	"github.com/YoungAgency/simplefix-go/session/messages"
)

const MsgTypeLogon = "A"

type Logon struct {
	*fix.Message
}

func makeLogon() *Logon {
	msg := &Logon{
		Message: fix.NewMessage(FieldBeginString, FieldBodyLength, FieldCheckSum, FieldMsgType, beginString, MsgTypeLogon).
			SetBody(
				fix.NewKeyValue(FieldEncryptMethod, &fix.String{}),
				fix.NewKeyValue(FieldHeartBtInt, &fix.Int{}),
				fix.NewKeyValue(FieldRawDataLength, &fix.Int{}),
				fix.NewKeyValue(FieldRawData, &fix.String{}),
				fix.NewKeyValue(FieldResetSeqNumFlag, &fix.Bool{}),
				fix.NewKeyValue(FieldNextExpectedMsgSeqNum, &fix.Int{}),
				fix.NewKeyValue(FieldMaxMessageSize, &fix.Int{}),
				NewMsgTypesGrp().Group,
				fix.NewKeyValue(FieldTestMessageIndicator, &fix.Bool{}),
				fix.NewKeyValue(FieldUsername, &fix.String{}),
				fix.NewKeyValue(FieldPassword, &fix.String{}),
			),
	}

	msg.SetHeader(makeHeader().AsComponent())
	msg.SetTrailer(makeTrailer().AsComponent())

	return msg
}

func CreateLogon(encryptMethod string, heartBtInt int) *Logon {
	msg := makeLogon().
		SetEncryptMethod(encryptMethod).
		SetHeartBtInt(heartBtInt)

	return msg
}

func NewLogon() *Logon {
	m := makeLogon()
	return &Logon{
		fix.NewMessage(FieldBeginString, FieldBodyLength, FieldCheckSum, FieldMsgType, beginString, MsgTypeLogon).
			SetBody(m.Body()...).
			SetHeader(m.Header().AsComponent()).
			SetTrailer(m.Trailer().AsComponent()),
	}
}

func (logon *Logon) Header() *Header {
	header := logon.Message.Header()

	return &Header{header}
}

func (logon *Logon) HeaderBuilder() messages.HeaderBuilder {
	return logon.Header()
}

func (logon *Logon) Trailer() *Trailer {
	trailer := logon.Message.Trailer()

	return &Trailer{trailer}
}

func (logon *Logon) EncryptMethod() string {
	kv := logon.Get(0)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (logon *Logon) SetEncryptMethod(encryptMethod string) *Logon {
	kv := logon.Get(0).(*fix.KeyValue)
	_ = kv.Load().Set(encryptMethod)
	return logon
}

func (logon *Logon) HeartBtInt() int {
	kv := logon.Get(1)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(int)
}

func (logon *Logon) SetHeartBtInt(heartBtInt int) *Logon {
	kv := logon.Get(1).(*fix.KeyValue)
	_ = kv.Load().Set(heartBtInt)
	return logon
}

func (logon *Logon) RawDataLength() int {
	kv := logon.Get(2)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(int)
}

func (logon *Logon) SetRawDataLength(rawDataLength int) *Logon {
	kv := logon.Get(2).(*fix.KeyValue)
	_ = kv.Load().Set(rawDataLength)
	return logon
}

func (logon *Logon) RawData() string {
	kv := logon.Get(3)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (logon *Logon) SetRawData(rawData string) *Logon {
	kv := logon.Get(3).(*fix.KeyValue)
	_ = kv.Load().Set(rawData)
	return logon
}

func (logon *Logon) ResetSeqNumFlag() bool {
	kv := logon.Get(4)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(bool)
}

func (logon *Logon) SetResetSeqNumFlag(resetSeqNumFlag bool) *Logon {
	kv := logon.Get(4).(*fix.KeyValue)
	_ = kv.Load().Set(resetSeqNumFlag)
	return logon
}

func (logon *Logon) NextExpectedMsgSeqNum() int {
	kv := logon.Get(5)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(int)
}

func (logon *Logon) SetNextExpectedMsgSeqNum(nextExpectedMsgSeqNum int) *Logon {
	kv := logon.Get(5).(*fix.KeyValue)
	_ = kv.Load().Set(nextExpectedMsgSeqNum)
	return logon
}

func (logon *Logon) MaxMessageSize() int {
	kv := logon.Get(6)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(int)
}

func (logon *Logon) SetMaxMessageSize(maxMessageSize int) *Logon {
	kv := logon.Get(6).(*fix.KeyValue)
	_ = kv.Load().Set(maxMessageSize)
	return logon
}

func (logon *Logon) MsgTypesGrp() *MsgTypesGrp {
	group := logon.Get(7).(*fix.Group)

	return &MsgTypesGrp{group}
}

func (logon *Logon) SetMsgTypesGrp(noMsgTypes *MsgTypesGrp) *Logon {
	logon.Set(7, noMsgTypes.Group)

	return logon
}

func (logon *Logon) TestMessageIndicator() bool {
	kv := logon.Get(8)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(bool)
}

func (logon *Logon) SetTestMessageIndicator(testMessageIndicator bool) *Logon {
	kv := logon.Get(8).(*fix.KeyValue)
	_ = kv.Load().Set(testMessageIndicator)
	return logon
}

func (logon *Logon) Username() string {
	kv := logon.Get(9)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (logon *Logon) SetUsername(username string) *Logon {
	kv := logon.Get(9).(*fix.KeyValue)
	_ = kv.Load().Set(username)
	return logon
}

func (logon *Logon) Password() string {
	kv := logon.Get(10)
	v := kv.(*fix.KeyValue).Load().Value()
	return v.(string)
}

func (logon *Logon) SetPassword(password string) *Logon {
	kv := logon.Get(10).(*fix.KeyValue)
	_ = kv.Load().Set(password)
	return logon
}

// New is a plane message constructor
func (Logon) New() messages.LogonBuilder {
	return makeLogon()
}

// Build provides an opportunity to customize message during building outgoing message
func (Logon) Build() messages.LogonBuilder {
	return makeLogon()
}

func (logon *Logon) SetFieldHeartBtInt(heartBtInt int) messages.LogonBuilder {
	return logon.SetHeartBtInt(heartBtInt)
}

func (logon *Logon) SetFieldEncryptMethod(encryptMethod string) messages.LogonBuilder {
	return logon.SetEncryptMethod(encryptMethod)
}

func (logon *Logon) SetFieldPassword(password string) messages.LogonBuilder {
	return logon.SetPassword(password)
}

func (logon *Logon) SetFieldUsername(username string) messages.LogonBuilder {
	return logon.SetUsername(username)
}

func (logon *Logon) SetFieldResetSeqNumFlag(resetSeqNumFlag bool) messages.LogonBuilder {
	return logon.SetResetSeqNumFlag(resetSeqNumFlag)
}
