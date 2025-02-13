package session

import (
	simplefixgo "github.com/YoungAgency/simplefix-go"
	"github.com/YoungAgency/simplefix-go/fix"
)

// MessageStorage is an interface providing a basic method for storing messages awaiting to be sent.
type MessageStorage interface {
	Save(storageID fix.StorageID, msg simplefixgo.SendingMessage, msgSeqNum int) error
	Messages(storageID fix.StorageID, msgSeqNumFrom, msgSeqNumTo int) ([]simplefixgo.SendingMessage, error)
}

type CounterStorage interface {
	GetNextSeqNum(storageID fix.StorageID) (int, error)
	GetCurrSeqNum(storageID fix.StorageID) (int, error)
	ResetSeqNum(storageID fix.StorageID) error
	SetSeqNum(storageID fix.StorageID, seqNum int) error
}
