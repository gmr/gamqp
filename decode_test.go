package gamqp

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestDecodeBoolean(t *testing.T) {
	cases := []struct {
		in    []byte
		value bool
		err   error
	}{
		{[]byte{0x00}, false, nil},
		{[]byte{0x01}, true, nil},
		{[]byte{}, false, io.EOF},
	}
	for _, c := range cases {
		value, err := decodeBoolean(bytes.NewReader(c.in))
		if err != nil {
			if err != c.err {
				t.Errorf("decodeBoolean(%s) == %#v, want %#v", c.in, err, c.err)
			}
		} else if value != c.value {
			t.Errorf("decodeBoolean(%v) == %t, want %t", c.in, value, c.value)
		}
	}
}

func TestDecodeShortStr(t *testing.T) {
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
		value, err := decodeShortStr(strings.NewReader(c.in))
		if err != nil {
			if err != c.err {
				t.Errorf("shortStr(%s) == %#v, want %#v", c.in, err, c.err)
			}
		} else {
			if value != c.value {
				t.Errorf("shortStr(%s) == %s, want %s", c.in, value, c.value)
			}
		}
	}

}
