package gamqp

import (
	"bytes"
	"io"
	"strings"
	"testing"
	"time"
)

func TestDecodeBoolean(t *testing.T) {
	cases := []struct {
		in    string
		value bool
		err   error
	}{
		{"\000", false, nil},
		{"\001", true, nil},
		{"", false, io.EOF},
	}
	for _, c := range cases {
		value, err := decodeBoolean(strings.NewReader(c.in))
		if err != nil {
			if err != c.err {
				t.Errorf("decodeLongString(%#v) == %#v, want %#v", c.in, err, c.err)
			}
		} else {
			if value != c.value {
				t.Errorf("decodeLongString(%#v) == %#v, want %#v", c.in, value, c.value)
			}
		}
	}
}

func TestDecodeByteArray(t *testing.T) {
	cases := []struct {
		in    string
		value []byte
		err   error
	}{
		{"\000\000\000\003abc", []byte{0x61, 0x62, 0x63}, nil},
		{"\000\000\000\003", []byte{0x00}, io.EOF},
		{"", []byte{0x00}, io.EOF},
	}
	for _, c := range cases {
		value, err := decodeByteArray(strings.NewReader(c.in))
		if err != nil {
			if err != c.err {
				t.Errorf("decodeByteArray(%#v) == %#v, want %#v", c.in, err, c.err)
			}
		} else {
			if bytes.Compare(value, c.value) != 0 {
				t.Errorf("decodeByteArray(%#v) == %#v, want %#v", c.in, value, c.value)
			}
		}
	}
}

func TestDecodeDouble(t *testing.T) {
	cases := []struct {
		in    string
		value float64
		err   error
	}{
		{"B\002\xa0_ \000\000\000", float64(1e+10), nil},
		{"", float64(0), io.EOF},
	}
	for _, c := range cases {
		value, err := decodeDouble(strings.NewReader(c.in))
		if err != nil {
			if err != c.err {
				t.Errorf("decodeDouble(%#v) == %#v, want %#v", c.in, err, c.err)
			}
		} else {
			if value != c.value {
				t.Errorf("decodeDouble(%#v) == %#v, want %#v", c.in, value, c.value)
			}
		}
	}
}

func TestDecodeFloat(t *testing.T) {
	cases := []struct {
		in    string
		value float32
		err   error
	}{
		{"@I\x0f\xd0", float32(3.14159), nil},
		{"", float32(0), io.EOF},
	}
	for _, c := range cases {
		value, err := decodeFloat(strings.NewReader(c.in))
		if err != nil {
			if err != c.err {
				t.Errorf("decodeFloat(%#v) == %#v, want %#v", c.in, err, c.err)
			}
		} else {
			if value != c.value {
				t.Errorf("decodeFloat(%#v) == %#v, want %#v", c.in, value, c.value)
			}
		}
	}
}

func TestDecodeLongInt(t *testing.T) {
	cases := []struct {
		in    string
		value uint32
		err   error
	}{
		{"\x7f\xff\xff\xff", 2147483647, nil},
		{"", 0, io.EOF},
	}
	for _, c := range cases {
		value, err := decodeLongInt(strings.NewReader(c.in))
		if err != nil {
			if err != c.err {
				t.Errorf("decodeLongInt(%#v) == %#v, want %#v", c.in, err, c.err)
			}
		} else {
			if value != c.value {
				t.Errorf("decodeLongInt(%#v) == %#v, want %#v", c.in, value, c.value)
			}
		}
	}
}

func TestDecodeLongLongInt(t *testing.T) {
	cases := []struct {
		in    string
		value uint64
		err   error
	}{
		{"\x7f\xff\xff\xff\xff\xff\xff\xf8", 9223372036854775800, nil},
		{"", 0, io.EOF},
	}
	for _, c := range cases {
		value, err := decodeLongLongInt(strings.NewReader(c.in))
		if err != nil {
			if err != c.err {
				t.Errorf("decodeLongLongInt(%#v) == %#v, want %#v", c.in, err, c.err)
			}
		} else {
			if value != c.value {
				t.Errorf("decodeLongLongInt(%#v) == %#v, want %#v", c.in, value, c.value)
			}
		}
	}
}

func TestDecodeLongString(t *testing.T) {
	cases := []struct {
		in, value string
		err       error
	}{
		{"\000\000\000\n0123456789", "0123456789", nil},
		{"\000\000\000\014abc123def456", "abc123def456", nil},
		{"aabc123", "", io.ErrUnexpectedEOF},
		{"", "", io.EOF},
	}
	for _, c := range cases {
		value, err := decodeLongString(strings.NewReader(c.in))
		if err != nil {
			if err != c.err {
				t.Errorf("decodeLongString(%s) == %#v, want %#v", c.in, err, c.err)
			}
		} else {
			if value != c.value {
				t.Errorf("decodeLongString(%s) == %s, want %s", c.in, value, c.value)
			}
		}
	}
}

func TestDecodeOctet(t *testing.T) {
	cases := []struct {
		in    string
		value byte
		err   error
	}{
		{"\xfe", byte(254), nil},
		{"", byte(0), io.EOF},
	}
	for _, c := range cases {
		value, err := decodeOctet(strings.NewReader(c.in))
		if err != nil {
			if err != c.err {
				t.Errorf("decodeOctet(%#v) == %#v, want %#v", c.in, err, c.err)
			}
		} else {
			if value != c.value {
				t.Errorf("decodeOctet(%#v) == %#v, want %#v", c.in, value, c.value)
			}
		}
	}
}

func TestDecodeShortInt(t *testing.T) {
	cases := []struct {
		in    string
		value uint16
		err   error
	}{
		{"\x7f\xff", 32767, nil},
		{"", 0, io.EOF},
	}
	for _, c := range cases {
		value, err := decodeShortInt(strings.NewReader(c.in))
		if err != nil {
			if err != c.err {
				t.Errorf("decodeShortInt(%#v) == %#v, want %#v", c.in, err, c.err)
			}
		} else {
			if value != c.value {
				t.Errorf("decodeShortInt(%#v) == %#v, want %#v", c.in, value, c.value)
			}
		}
	}
}

func TestDecodeShortString(t *testing.T) {
	cases := []struct {
		in, value string
		err       error
	}{
		{"\n0123456789", "0123456789", nil},
		{"\014abc123def456", "abc123def456", nil},
		{"aabc123", "", io.ErrUnexpectedEOF},
		{"", "", io.EOF},
	}
	for _, c := range cases {
		value, err := decodeShortString(strings.NewReader(c.in))
		if err != nil {
			if err != c.err {
				t.Errorf("decodeShortString(%#v) == %#v, want %#v", c.in, err, c.err)
			}
		} else {
			if value != c.value {
				t.Errorf("decodeShortString(%#v) == %#v, want %#v", c.in, value, c.value)
			}
		}
	}
}

func TestDecodeTimestamp(t *testing.T) {
	cases := []struct {
		in    string
		value time.Time
		err   error
	}{
		{"\x00\x00\x00\x00Ec)\x92", time.Unix(1164126610, 0), nil},
		{"", time.Unix(0, 0), io.EOF},
	}
	for _, c := range cases {
		value, err := decodeTimestamp(strings.NewReader(c.in))
		if err != nil {
			if err != c.err {
				t.Errorf("decodeTimestamp(%#v) == %#v, want %#v", c.in, err, c.err)
			}
		} else {
			if value != c.value {
				t.Errorf("decodeTimestamp(%#v) == %#v, want %#v", c.in, value, c.value)
			}
		}
	}
}
