package ziface

type IMessage interface {
	GetMsgId() uint32
	GetDataLen() uint32
	GetData() []byte

	SetMsgId(id uint32)
	SetDataLen(len uint32)
	SetData(data []byte)
}