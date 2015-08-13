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
	var length uint32
	err = binary.Read(r, binary.BigEndian, &length)
	if err != nil {
		return
	}
	bytes := make([]byte, length)
	if _, err = io.ReadFull(r, bytes); err != nil {
		return
	}
	return bytes, nil
}

func decodeDecimal(r io.Reader) (value Decimal, err error) {
	return
}

func decodeDouble(r io.Reader) (value float64, err error) {
	err = binary.Read(r, binary.BigEndian, &value)
	return
}

func decodeFieldArray(r io.Reader) (value []interface{}, err error) {
	return
}

func decodeFieldTable(r io.Reader) (value map[string]interface{}, err error) {
	return
}

func decodeFloat(r io.Reader) (value float32, err error) {
	err = binary.Read(r, binary.BigEndian, &value)
	return
}

func decodeLongInt(r io.Reader) (value uint32, err error) {
	err = binary.Read(r, binary.BigEndian, &value)
	return
}

func decodeLongString(r io.Reader) (value string, err error) {
	var length uint32
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

func decodeLongLongInt(r io.Reader) (value uint64, err error) {
	err = binary.Read(r, binary.BigEndian, &value)
	return
}

func decodeOctet(r io.Reader) (value byte, err error) {
	err = binary.Read(r, binary.BigEndian, &value)
	return
}

func decodeShortInt(r io.Reader) (value uint16, err error) {
	err = binary.Read(r, binary.BigEndian, &value)
	return
}

func decodeShortShortInt(r io.Reader) (value int8, err error) {
	err = binary.Read(r, binary.BigEndian, &value)
	return
}

func decodeShortString(r io.Reader) (value string, err error) {
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
	var epoch int64
	err = binary.Read(r, binary.BigEndian, &epoch)
	if err != nil {
		return
	}
	return time.Unix(epoch, 0), nil
}
