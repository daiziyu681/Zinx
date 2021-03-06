package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
	"zinx/utils"
	"zinx/ziface"
)

type Connection struct {
	// belong to server
	TcpServer ziface.Iserver

	// current connect socket
	Conn *net.TCPConn

	// connect id
	ConnID uint32

	// connect state
	isClosed bool

	// channel
	ExitChan chan bool

	// message channel
	msgChan chan []byte

	// connection router
	MsgHandler ziface.IMsgHandle

	// connection property
	property map[string]interface{}

	// property lock
	propertyLock sync.RWMutex
}

func NewConnection(server ziface.Iserver, conn *net.TCPConn, connID uint32, msgHandler ziface.IMsgHandle) *Connection {
	c := &Connection{
		TcpServer:  server,
		Conn:       conn,
		ConnID:     connID,
		isClosed:   false,
		ExitChan:   make(chan bool, 1),
		msgChan:    make(chan []byte),
		MsgHandler: msgHandler,
		property:   make(map[string]interface{}),
	}

	c.TcpServer.GetConnMgr().Add(c)

	return c
}

func (c *Connection) StartReader() {
	fmt.Println("[Reader Goroutine is running...]")
	defer fmt.Println("[conn reader exit!], remote addr is ", c.RemoteAddr().String())
	defer c.Stop()

	for {
		// buf := make([]byte, utils.GlobalObject.MaxPackageSize)
		// _, err := c.Conn.Read(buf)
		// if err != nil {
		// 	fmt.Println("recv buff error : ", err)
		// 	continue
		// }

		// create data pack object
		dp := NewDataPack()

		dataHead := make([]byte, dp.GetHeadLen())

		if _, err := io.ReadFull(c.GetTCPConnection(), dataHead); err != nil {
			fmt.Println("read data head error : ", err)
			break
		}

		msg, err := dp.Unpack(dataHead)
		if err != nil {
			fmt.Println("unpack error : ", err)
			break
		}

		var data []byte
		if msg.GetDataLen() > 0 {
			data = make([]byte, msg.GetDataLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read data error : ", err)
				break
			}
		}
		msg.SetData(data)

		// get Request using tcp connection and reading data
		req := NewRequest(c, msg)

		if utils.GlobalObject.WorkerPoolSize > 0 {
			c.MsgHandler.SendMsgToTaskQueue(req)
		} else {
			go c.MsgHandler.DoMsgHandle(req)
		}
	}
}

func (c *Connection) StartWriter() {
	fmt.Println("[Writer Goroutine is running...]")
	defer fmt.Println("[conn writer exit!], remote addr is ", c.RemoteAddr().String())

	for {
		select {
		case data := <-c.msgChan:
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("Send data error, ", err)
				return
			}
		case <-c.ExitChan:
			return
		}
	}
}

// start link
func (c *Connection) Start() {
	fmt.Println("Conn start.. ConnID = ", c.ConnID)

	go c.StartReader()
	go c.StartWriter()

	c.TcpServer.CallOnConnStart(c)
}

// stop link
func (c *Connection) Stop() {
	fmt.Println("Conn stop.. ConnID = ", c.ConnID)

	if c.isClosed == true {
		return
	}
	c.isClosed = true

	c.TcpServer.CallOnConnStop(c)

	c.Conn.Close()

	c.ExitChan <- true

	c.TcpServer.GetConnMgr().Remove(c)

	close(c.ExitChan)
	close(c.msgChan)
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
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed == true {
		return errors.New("connection closed when send msg")
	}

	dp := NewDataPack()

	binaryMsg, err := dp.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		fmt.Println("pack error msg id : ", msgId)
		return errors.New("pack msg error")
	}

	c.msgChan <- binaryMsg

	return nil
}

func (c *Connection) SetProperty(key string, value interface{}) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	c.property[key] = value
}

func (c *Connection) GetProperty(key string) (interface{}, error) {
	c.propertyLock.RLock()
	defer c.propertyLock.RUnlock()

	if value, ok := c.property[key]; ok {
		return value, nil
	} else {
		return nil, errors.New("no property found")
	}
}

func (c *Connection) RemoveProperty(key string) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	delete(c.property, key)
}
