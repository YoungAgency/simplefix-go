package fix

import (
	"fmt"
	"strings"
)

// Item required method for basic FIX-item
type Item interface {
	ToBytes() []byte
	String() string
}

// Items array of Item elements
type Items []Item

// ToBytes convert Items to bytes
func (v Items) ToBytes() []byte {
	var msg [][]byte
	for _, item := range v {
		itemB := item.ToBytes()
		if itemB != nil {
			msg = append(msg, itemB)
		}
	}
	return joinBody(msg...)
}

func (v Items) String() string {
	var items []string
	for _, item := range v {
		itemStr := item.String()
		if itemStr != "" {
			items = append(items, itemStr)
		}
	}
	return fmt.Sprintf("{%s}", strings.Join(items, ", "))
}
