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
	conn net.Conn
}

func NewClient(net, addr string) *Client {
	cli := &Client{net: net, addr: addr, conn: nil}
	cli.getConn()
	return cli
}

func (c *Client) getConn() error {
	if c.conn != nil {
		c.conn.Close()
	}

	conn, err := net.Dial(c.net, c.addr)
	if err != nil {
		logs.Logger.Debugf("Get Client Conn Error: %v", err)
	} else {
		c.conn = conn
	}

	return err
}

func (c *Client) ReConnect() {
	c.getConn()
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) SendStrings(t byte, str ...string) error {
	buf := bytes.NewBuffer(make([]byte, 0))

	for _, v := range str {
		buf.WriteByte(t)
		buf.Write(utils.IntToBytes(int32(len(v))))
		buf.Write([]byte(v))
		n, err := c.conn.Write(buf.Bytes())
		if err != nil {
			logs.Logger.Errorf("client send error")
			return err
		}
		buf.Reset()
		logs.Logger.Debugf("send %v bytes", n)
	}

	return nil
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
