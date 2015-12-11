package main

import (
	"common"
	"fmt"
	"net"
)

const (
	TM_REG = 1 >> iota
	TM_HEATBEAT
	TW_REG
)

type Dtm struct {
	regAddr net.TCPAddr
	dtws    []Dtw
}

type Dtw struct {
	id    []byte
	tasks []Task
}

type Task struct {
	id   []byte
	exec func(...interface{}) interface{}
}

type Guard struct {
	regAddr net.TCPAddr
	dtms    []Dtm
	dtws    []Dtw
}

func main() {
	m := common.NewMsg(1, 10, "Hello World!")
	fmt.Println(m)
	bs, err := m.Marshal()
	fmt.Println(bs, err)
	n := new(common.Msg)
	fmt.Println(n)
	n.UnMarshal(bs)
	fmt.Println(n)
}
