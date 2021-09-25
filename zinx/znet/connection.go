package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
)

type Connection struct {
	// current connect socket
	Conn *net.TCPConn

	// connect id
	ConnID uint32

	// connect state
	isClosed bool

	// channel
	ExitChan chan bool

	// connection router
	Router ziface.IRouter
}

func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn : conn,
		ConnID : connID,
		isClosed : false,
		ExitChan : make(chan bool, 1),
		Router : router,
	}

	return c
}

func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running..")
	defer fmt.Println("connID = ", c.ConnID, " Reader is exit, remote addr is ", c.RemoteAddr().String())
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buff error : ", err)
			continue
		}

		// get Request using tcp connection and reading data
		req := NewRequest(c, buf)

		go func() {
			c.Router.PreHandle(req)
			c.Router.Handle(req)
			c.Router.PostHandle(req)
		}()
	}
}

// start link
func (c *Connection) Start() {
	fmt.Println("Conn start.. ConnID = ", c.ConnID)

	go c.StartReader()
}

// stop link
func (c *Connection) Stop() {
	fmt.Println("Conn stop.. ConnID = ", c.ConnID)

	if c.isClosed == true {
		return
	}
	c.isClosed = true

	c.Conn.Close()

	close(c.ExitChan)
}

// get socker conn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// get link id
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// get client ip port
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// send msg to remote client
func (c *Connection) Send(data []byte) error {
	return nil
}