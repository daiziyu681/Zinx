package ziface

// pack message and unpack message data
// aims to solve tcp stick package problem

type IDatapack interface {
	/*
		TLV format : type, len, value
	*/
	GetHeadLen() uint32

	Pack(msg IMessage) ([]byte, error)

	Unpack([]byte) (IMessage, error)
}