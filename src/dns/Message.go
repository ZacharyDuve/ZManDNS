package dns

import (
	"errors"
	"fmt"
)

const MessageLength int = 512

type Message [MessageLength]byte

func NewMessage() Message {
	return Message{}
}

const (
	idHighByteIndex int = 0
	idLowByteIndex  int = 1
)

func (this Message) Id() uint16 {
	var id uint16 = (uint16(this[idHighByteIndex]) << 8)
	id |= uint16(this[idLowByteIndex])
	return id
}

type MessageType byte

const (
	Query    MessageType = 0
	Response MessageType = 1
)

const (
	messageTypeByteMask  byte = 0x80
	messageTypeByteIndex int  = 2
)

func (this Message) Type() MessageType {
	if this[messageTypeByteIndex]&messageTypeByteMask == 0 {
		return Query
	}
	return Response
}

type OPCode byte

const (
	StandardQuery  OPCode = 0
	InverseQuery   OPCode = 1
	ServerStatus   OPCode = 2
	OPCode_Unknown OPCode = 0xff
)

const (
	opCodeByteMask  byte = 0x78
	opCodeByteIndex int  = 2
)

func (this Message) OPCode() (OPCode, error) {
	opCodeValue := (this[opCodeByteIndex] & opCodeByteMask) >> 3

	if opCodeValue <= 2 {
		return OPCode(opCodeValue), nil
	}

	return OPCode_Unknown, errors.New(fmt.Sprintf("Unknown Opcode %d", opCodeValue))
}

const (
	aaByteIndex int  = 2
	aaByteMask  byte = 0x04
)

func (this Message) IsAuthorativeAnswer() bool {
	return this[aaByteIndex]&aaByteMask != 0
}

const (
	tcByteIndex int  = 2
	tcByteMask  byte = 0x02
)

func (this Message) IsTruncated() bool {
	return this[tcByteIndex]&tcByteMask != 0
}

const (
	rdByteIndex int  = 2
	rdByteMask  byte = 0x01
)

func (this Message) RecursionDesired() bool {
	return this[rdByteIndex]&rdByteMask != 0
}

const (
	raByteIndex int  = 3
	raByteMask  byte = 0x80
)

func (this Message) RecursionAvailable() bool {
	return this[raByteIndex]&raByteMask != 0
}

type ReturnCode byte

const (
	ReturnCodeNoError   ReturnCode = 0
	ReturnCodeNameError ReturnCode = 3
	ReturnCodeUnknown   ReturnCode = 0xff

	rcByteIndex int  = 3
	rcByteMask  byte = 0x0f
)

func (this Message) ReturnCode() (ReturnCode, error) {
	rc := this[rcByteIndex] & rcByteMask

	if rc != 0 && rc != 3 {
		return ReturnCodeUnknown, errors.New(fmt.Sprintf("Error Unknown Return code %d", rc))
	}

	return ReturnCode(rc), nil
}

func (this Message) getUint16FromIndexes(highIndex, lowIndex int) uint16 {
	v := uint16(this[highIndex]) << 8
	v |= uint16(this[lowIndex])
	return v
}

const (
	numQuestionsHighIndex int = 4
	numQuestionsLowIndex  int = 5
)

func (this Message) NumberQuestions() uint16 {
	return this.getUint16FromIndexes(numQuestionsHighIndex, numQuestionsLowIndex)
}

const (
	numAnswersHighIndex int = 6
	numAnswersLowIndex  int = 7
)

func (this Message) NumberAnswers() uint16 {
	return this.getUint16FromIndexes(numAnswersHighIndex, numAnswersLowIndex)
}

const (
	numAARecordsHighIndex = 8
	numAARecordsLowIndex  = 9
)

func (this Message) NumberAuthorativeAnswers() uint16 {
	return this.getUint16FromIndexes(numAARecordsHighIndex, numAARecordsLowIndex)
}

const (
	numAdditionalAnswersHighIndex = 10
	numAdditionalAnswersLowIndex  = 11
)

func (this Message) NumberAdditionalAnswers() uint16 {
	return this.getUint16FromIndexes(numAdditionalAnswersHighIndex, numAdditionalAnswersLowIndex)
}
