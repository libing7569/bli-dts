package common

import (
	"encoding/json"
)

type Msg struct {
	T byte
	L int32
	V interface{}
}

func NewMsg(t byte, l int32, v interface{}) *Msg {
	return &Msg{T: t, L: l, V: v}
}

func (msg *Msg) Marshal() ([]byte, error) {
	return json.Marshal(msg)
}

func (msg *Msg) UnMarshal(msgBs []byte) error {
	return json.Unmarshal(msgBs, msg)
}
