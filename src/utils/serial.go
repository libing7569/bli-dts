package utils

import (
	"bytes"
	"encoding/asn1"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
)

type SerialType byte

const (
	JSON_SERIAL = 1 >> iota
	ASN1_SERIAL
)

func Bytes2Int32(bs []byte) int32 {
	var n int32
	buf := bytes.NewBuffer(bs)
	err := binary.Read(buf, binary.BigEndian, &n)
	if err != nil {
		fmt.Fprintf(os.Stderr, "binary read error: ", err)
		os.Exit(1)
	}
	return n
}

func IntToBytes(i interface{}) []byte {
	buf := bytes.NewBuffer(make([]byte, 0))
	binary.Write(buf, binary.BigEndian, i)
	return buf.Bytes()
}

func Serialize(t SerialType, v interface{}) []byte {
	var bs []byte = nil
	var err error = nil
	switch t {
	case JSON_SERIAL:
		bs, err = json.Marshal(v)
	case ASN1_SERIAL:
		bs, err = asn1.Marshal(v)
	default:
		fmt.Fprintf(os.Stderr, "unkown serial type")
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "serial error: ", err)
		os.Exit(1)
	}

	return bs
}

func UnSerialize(t SerialType, bs []byte, v interface{}) {
	var err error = nil
	switch t {
	case JSON_SERIAL:
		fmt.Println("json serialize")
		err = json.Unmarshal(bs, v)
	case ASN1_SERIAL:
		_, err = asn1.Unmarshal(bs, v)
	default:
		fmt.Fprintf(os.Stderr, "unkown serial type")
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "unserial error: ", err)
		os.Exit(1)
	}
}
