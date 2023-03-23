package helper

import (
	"bytes"
	"encoding/binary"
	"log"
	"runtime/debug"

	"github.com/mr-tron/base58"
)

func Handle(err error) {
	if err != nil {
		log.Println("Ooh!!! Panic:")
		log.Println(string(debug.Stack()))
		// debug.PrintStack()
		log.Panic(err)
	}
}

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)

	}

	return buff.Bytes()
}

func Base58Encode(input []byte) []byte {
	encode := base58.Encode(input)

	return []byte(encode)
}

func Base58Decode(input []byte) []byte {
	decode, err := base58.Decode(string(input[:]))
	Handle(err)

	return decode
}
