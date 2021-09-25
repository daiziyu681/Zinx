package znet

import "zinx/ziface"

type Request struct {
	conn ziface.IConnection

	data []byte
}

// get current connection
func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

// get request msg data
func (r *Request) GetData() []byte {
	return r.data
}

func NewRequest(c ziface.IConnection, buf []byte) *Request {
	r := &Request {
		conn : c,
		data : buf,
	}
	return r
}