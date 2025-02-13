package messages

import "github.com/YoungAgency/simplefix-go/fix"

type Builder interface {
	Items() fix.Items
	CalcBodyLength() int
	BodyLength() int
	BytesWithoutChecksum() []byte
	CheckSum() string
	BeginString() *fix.KeyValue
	MsgType() string
	ToBytes() ([]byte, error)
	BeginStringTag() string
	BodyLengthTag() string
	CheckSumTag() string
}

type PipelineBuilder interface {
	HeaderBuilder() HeaderBuilder
	Builder
}
