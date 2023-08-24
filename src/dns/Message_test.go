package dns

import (
	"net"
	"testing"
)

func TestThatMessageWithIdInBytes0And1ReturnsMatchingId(t *testing.T) {
	data := make([]byte, 12)
	data[0] = 0xF0
	data[1] = 0x9C
	m, _ := NewMessage(&net.IPAddr{}, data)

	if uint16(m.Id()) != 0xF09C {
		t.Log(uint16(m.Id()), 0xF09C)
		t.Fail()
	}
}

func TestThatMessageWithResponseFlagSetTo1ReturnResponseForType(t *testing.T) {
	data := make([]byte, 12)
	data[2] = 0x80
	m, _ := NewMessage(&net.IPAddr{}, data)

	if m.Type() != Response {
		t.Fail()
	}
}

func TestThatMessageWithResponseFlagSetTo0ReturnQueryForType(t *testing.T) {
	data := make([]byte, 12)
	data[2] = 0x00

	m, _ := NewMessage(&net.IPAddr{}, data)
	if m.Type() != Query {
		t.Fail()
	}
}
