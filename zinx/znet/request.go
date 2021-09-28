package znet

import "zinx/ziface"

type Request struct {
	conn ziface.IConnection

	msg ziface.IMessage
}

// get current connection
func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

// get request msg data
func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

func (r *Request) GetMsgId() uint32 {
	return r.msg.GetMsgId()
}

func NewRequest(c ziface.IConnection, m ziface.IMessage) *Request {
	r := &Request {
		conn : c,
		msg : m,
	}
	return r
}