package main

import (
	"fmt"
	"utils"
)

//type TaskParticle struct {
//preconds []uint8
//inputs   []interface{}
//execute  []func(...interface{}) interface{}
//}

type Task struct {
	Id string
}

type SimpleTask struct {
	Task
	ExecName string
	//exec func(...interface{}) interface{}
}

func (task *Task) Serialize(t utils.SerialType) []byte {
	return utils.Serialize(utils.JSON_SERIAL, task)
}

func (task *Task) UnSerialize(t utils.SerialType, bs []byte) {
	utils.UnSerialize(utils.JSON_SERIAL, bs, task)
}

func main() {
	t := &Task{"hello"}
	bs := t.Serialize(utils.JSON_SERIAL)
	fmt.Println(t)
	fmt.Println(bs)
	tt := &Task{}
	tt.UnSerialize(utils.JSON_SERIAL, bs)
	fmt.Println(tt)
}
