package ziface

// define server interface
type Iserver interface {
	// start server
	Start()

	// stop server
	Stop()

	// run server
	Serve()

	// add a router
	AddRouter(msgID uint32, router IRouter)

	// get conn manager
	GetConnMgr() IConnManager

	SetOnConnStart(func(connection IConnection))

	SetOnConnStop(func(connection IConnection))

	CallOnConnStart(connection IConnection)

	CallOnConnStop(connection IConnection)
}
