//Copyright 2015 Libing. All right reserved.
package main

import "system/communicate"

func main() {
	c := communicate.NewClient("tcp", "localhost:9999")
	c.SendStrings(1, "Hello world, hi li", "Hi ok", "{\"name\":\"hello\"}")

	//conn, err := net.Dial("tcp", "localhost:9999")
	//fmt.Println(conn, err)
	//buf := bytes.NewBuffer(make([]byte, 0))

	//var l int32 = 10

	//buf.WriteByte(1)
	//buf.Write(utils.IntToBytes(l))
	//buf.Write([]byte("{hello:99}"))

	//var l2 int32 = 7

	//for i := 0; i < 10; i++ {
	//buf.WriteByte(2)
	//buf.Write(utils.IntToBytes(l2))
	//buf.Write([]byte("{t"))

	//conn.Write(buf.Bytes())
	//time.Sleep(5 * time.Second)

	//buf.Reset()
	//buf.Write([]byte("i:99}"))
	//conn.Write(buf.Bytes())
	//buf.Reset()
	//}

	//bs := make([]byte, 1024)
	//n, err := conn.Read(bs)
	//fmt.Println(string(bs[:n]), n, err)
}
