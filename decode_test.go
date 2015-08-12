package gamqp

import (
	"bytes"
	"strings"
	"testing"
)

func TestBoolean(t *testing.T) {
	cases := []struct {
		in    []byte
		value bool
	}{
		{[]byte{0x00}, false},
		{[]byte{0x01}, true},
	}
	for _, c := range cases {
		value, _ := boolean(bytes.NewReader(c.in))
		if value != c.value {
			t.Errorf("boolean(%v) == %t, want %t", c.in, value, c.value)
		}
	}
}

func TestShortStr(t *testing.T) {
	cases := []struct {
		in, value string
	}{
		{"\n0123456789", "0123456789"},
		{"\014abc123def456", "abc123def456"},
	}
	for _, c := range cases {
		value, _ := shortStr(strings.NewReader(c.in))
		if value != c.value {
			t.Errorf("shortStr(%s) == %s, want %s", c.in, value, c.value)
		}
	}

}
