package ziface

import "net"

type IConnection interface {
	// start link
	Start()

	// stop link
	Stop()

	// get socker conn
	GetTCPConnection() *net.TCPConn

	// get link id
	GetConnID() uint32

	// get client ip port
	RemoteAddr() net.Addr

	// send msg to remote client
	SendMsg(msgId uint32, data []byte) error
}

type HandleFunc func(*net.TCPConn, []byte, int) error