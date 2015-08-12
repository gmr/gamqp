// Decode AMQP data types

package gamqp

import (
	"encoding/binary"
	"io"
)

func boolean(r io.Reader) (value bool, err error) {
	var byteValue uint8
	err = binary.Read(r, binary.BigEndian, &byteValue)
	if err != nil {
		return
	}
	value = byteValue != 0
	return value, nil
}

func shortStr(r io.Reader) (value string, err error) {
	var length uint8
	err = binary.Read(r, binary.BigEndian, &length)
	if err != nil {
		return
	}
	bytes := make([]byte, length)
	if _, err = io.ReadFull(r, bytes); err != nil {
		return
	}
	return string(bytes), nil
}
