package gamqp

import (
	"encoding/binary"
	"io"
	"time"
)

func decodeBit(r io.Reader) (value bool, err error) {
	return
}

func decodeBoolean(r io.Reader) (value bool, err error) {
	var byteValue uint8
	err = binary.Read(r, binary.BigEndian, &byteValue)
	if err != nil {
		return
	}
	value = byteValue != 0
	return value, nil
}

func decodeByteArray(r io.Reader) (value []byte, err error) {
	return
}

func decodeDecimal(r io.Reader) (value Decimal, err error) {
	return
}

func decodeDouble(r io.Reader) (value float64, err error) {
	return
}

func decodeFieldArray(r io.Reader) (value []interface{}, err error) {
	return
}

func decodeFieldTable(r io.Reader) (value map[string]interface{}, err error) {
	return
}

func decodeFloat(r io.Reader) (value float32, err error) {
	return
}

func decodeLongInt(r io.Reader) (value int32, err error) {
	return
}

func decodeLongString(r io.Reader) (value string, err error) {
	return
}

func decodeLongLongInt(r io.Reader) (value int64, err error) {
	return
}

func decodeOctet(r io.Reader) (value byte, err error) {
	return
}

func decodeShortInt(r io.Reader) (value int16, err error) {
	return
}

func decodeShortShortInt(r io.Reader) (value int8, err error) {
	return
}

func decodeShortStr(r io.Reader) (value string, err error) {
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

func decodeTimestamp(r io.Reader) (value time.Time, err error) {
	return
}
