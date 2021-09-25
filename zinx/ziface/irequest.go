package ziface

type IRequest interface {
	// get current connection
	GetConnection() IConnection

	// get request msg data
	GetData() []byte
}