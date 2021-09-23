package ziface

// define server interface
type Iserver interface {
	// start server
	Start()

	// stop server
	Stop()

	// run server
	Serve()
}