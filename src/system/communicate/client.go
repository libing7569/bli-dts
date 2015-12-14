//Copyright 2015 Libing. All right reserved.

package communicate

import (
	"bytes"
	"fmt"
	"net"
	"time"
	"utils"

	"utils/logs"
)

type Client struct {
	net  string
	addr string
}

func NewClient(net, addr string) *Client {
	return &Client{net: net, addr: addr}
}

func (c *Client) SendStrings(t byte, str ...string) error {
	conn, err := net.Dial("tcp", "localhost:9999")
	buf := bytes.NewBuffer(make([]byte, 0))

	for _, v := range str {
		buf.WriteByte(t)
		buf.Write(utils.IntToBytes(int32(len(v))))
		buf.Write([]byte(v))
		n, err := conn.Write(buf.Bytes())
		if err != nil {
			logs.Logger.Errorf("client send error")
			return err
		}
		buf.Reset()
		logs.Logger.Debugf("send %v bytes", n)
	}

	return err
}

func main() {
	conn, err := net.Dial("tcp", "localhost:9999")
	fmt.Println(conn, err)
	buf := bytes.NewBuffer(make([]byte, 0))

	var l int32 = 10

	buf.WriteByte(1)
	buf.Write(utils.IntToBytes(l))
	buf.Write([]byte("{hello:99}"))

	var l2 int32 = 7

	for i := 0; i < 10; i++ {
		buf.WriteByte(2)
		buf.Write(utils.IntToBytes(l2))
		buf.Write([]byte("{t"))

		conn.Write(buf.Bytes())
		time.Sleep(5 * time.Second)

		buf.Reset()
		buf.Write([]byte("i:99}"))
		conn.Write(buf.Bytes())
		buf.Reset()
	}

	bs := make([]byte, 1024)
	n, err := conn.Read(bs)
	logs.Logger.Debug(string(bs[:n]), n, err)
}
