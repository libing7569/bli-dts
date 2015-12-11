package system

import (
	"bytes"
	"encoding/binary"
	"errors"
	"net"
	"os"

	"utils/logs"
)

const (
	BUF_MAX_SIZE = 1024

	MSG_TYPE_SIZE     = 1
	MSG_DATA_LEN_SIZE = 4
)

const (
	MSG_TYPE_A = 1 >> iota
	MSG_TYPE_B
	MSG_TYPE_C
)

type Msg struct {
	T byte
	L int32
	V interface{}
}

type Server struct {
	net         string
	addr        string
	rawHandler  func(net.Conn, chan<- []byte) error
	dataHandler func(<-chan []byte) error
}

func Bytes2Int32(bs []byte) int32 {
	var n int32
	buf := bytes.NewBuffer(bs)
	err := binary.Read(buf, binary.BigEndian, &n)
	ErrorHandler(err, true)
	return n
}

func IntToBytes(i interface{}) []byte {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, i)
	return buf.Bytes()
}

func NewServer(net, addr string,
	rawHandler func(conn net.Conn, c chan<- []byte) error,
	dataHandler func(c <-chan []byte) error) *Server {
	if rawHandler == nil {
		rawHandler = defaultRawHandler
	}

	if dataHandler == nil {
		dataHandler = defaultDataHandler
	}

	return &Server{net: net, addr: addr, rawHandler: rawHandler, dataHandler: dataHandler}
}

func ErrorHandler(err error, isExit bool) {
	if err != nil {
		logs.Logger.Debugf("server error: %v", err.Error())
		if isExit {
			logs.Logger.Errorf("server error to exit: %v", err.Error())
			os.Exit(1)
		}
	}
}

func checkMsgType(t byte) error {
	return nil
}

func scanBuf2ExtractMsg(buf *bytes.Buffer, c chan<- []byte) {

	for {
		bs := buf.Bytes()
		if len(bs) < MSG_TYPE_SIZE {
			break
		}

		err := checkMsgType(bs[0])
		ErrorHandler(err, false)

		if len(bs) < MSG_TYPE_SIZE+MSG_DATA_LEN_SIZE {
			break
		}

		dataLen := int(Bytes2Int32(bs[1 : 1+MSG_DATA_LEN_SIZE]))

		logs.Logger.Debugf("extract data length: %v", dataLen)
		if len(bs) < MSG_TYPE_SIZE+MSG_DATA_LEN_SIZE+dataLen {
			break
		}

		logs.Logger.Debugf("buffer scan before read: %v", buf.Bytes())
		bsTmp := make([]byte, dataLen)
		buf.Next(MSG_TYPE_SIZE + MSG_DATA_LEN_SIZE)
		buf.Read(bsTmp)
		logs.Logger.Debugf("buffer scan after read: %v", buf.Bytes())
		logs.Logger.Debugf("scan to send: %v", bsTmp)
		c <- bsTmp
	}
}

func defaultRawHandler(conn net.Conn, c chan<- []byte) (err error) {
	bs := make([]byte, BUF_MAX_SIZE)
	buf := bytes.NewBuffer(make([]byte, 0))

	for {
		nr, err := conn.Read(bs)
		logs.Logger.Debugf("read %v bytes", nr)
		logs.Logger.Debugf("read data: %v", bs[:nr])
		ErrorHandler(err, false)
		if err != nil {
			break
		}

		nw, err := buf.Write(bs[:nr])
		logs.Logger.Debugf("write %v bytes", nw)
		logs.Logger.Debugf("raw buffer: %v", buf.Bytes())
		ErrorHandler(err, false)
		scanBuf2ExtractMsg(buf, c)
	}

	return
}

func defaultDataHandler(c <-chan []byte) error {
	for {
		select {
		case bs, err := <-c:
			if err {
				//TODO:
				logs.Logger.Debugf("data: %v", string(bs))
			} else {
				return errors.New("conn closed")
			}
		}
	}
}

func (s *Server) Run() {
	listener, err := net.Listen(s.net, s.addr)
	ErrorHandler(err, true)
	for {
		conn, err := listener.Accept()
		if err != nil {
			logs.Logger.Errorf("server accept error: %v\n", err.Error())
			continue
		}

		go func(conn net.Conn) {
			defer func() {
				conn.Close()
				logs.Logger.Debug("client closed")
			}()
			c := make(chan []byte)
			go s.dataHandler(c)
			s.rawHandler(conn, c)
		}(conn)
	}
}
