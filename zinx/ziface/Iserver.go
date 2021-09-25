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
	AddRouter(router IRouter)
}