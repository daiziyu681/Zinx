package znet

import (
	"zinx/ziface"
	"zinx/utils"
	"errors"
	"bytes"
	"encoding/binary"
)

type Datapack struct {}

func NewDataPack() *Datapack {
	return &Datapack {}
}

func (dp *Datapack) GetHeadLen() uint32 {
	// DataLen uint32(4 byte) + type uint32(4 byte) = 8 byte
	return 8
}

func (dp *Datapack) Pack(msg ziface.IMessage) ([]byte, error) {
	// create a byte buffer
	dataBuff := bytes.NewBuffer([]byte{})

	// little end storage
	// write datalen
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}

	// write msgType
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}

	// write data
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil
}

// read only message len and message type
func (dp *Datapack) Unpack(binaryData []byte) (ziface.IMessage, error) {
	// create a byte io reader
	dataBuff := bytes.NewReader(binaryData)

	msg := &Message {}

	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}

	// check message len
	if (utils.GlobalObject.MaxPackageSize > 0 && msg.DataLen > utils.GlobalObject.MaxPackageSize) {
		return nil, errors.New("too large msg data recv!")
	}

	return msg, nil
}