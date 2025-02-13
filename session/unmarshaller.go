package session

import (
	"github.com/YoungAgency/simplefix-go/session/messages"
)

type Unmarshaller interface {
	Unmarshal(msg messages.Builder, d []byte) error
}
