package dns

import "testing"

func TestThatMessageWithIdInBytes0And1ReturnsMatchingId(t *testing.T) {
	data := make([]byte, 12)
	data[0] = 0xF0
	m[1] = 0x9C
	m := NewMessage()

	if uint16(m.Id()) != 0xF09C {
		t.Log(uint16(m.Id()), 0xF09C)
		t.Fail()
	}
}

func TestThatMessageWithResponseFlagSetTo1ReturnResponseForType(t *testing.T) {
	m := NewMessage()
	m[2] = 0x80

	if m.Type() != Response {
		t.Fail()
	}
}

func TestThatMessageWithResponseFlagSetTo0ReturnQueryForType(t *testing.T) {
	m := NewMessage()
	m[2] = 0x00

	if m.Type() != Query {
		t.Fail()
	}
}
