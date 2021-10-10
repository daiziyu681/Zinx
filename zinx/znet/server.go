package znet

import (
	"fmt"
	"net"
	"zinx/utils"
	"zinx/ziface"
)

// server interface instance
type Server struct {
	// server name
	Name string
	// ip version
	IPVersion string
	// ip
	IP string
	// port
	Port int
	// msg handler module
	MsgHandler ziface.IMsgHandle
	// connection manager
	ConnMgr ziface.IConnManager
	// hook function called on connection start
	OnConnStart func(conn ziface.IConnection)
	// hook function called on connection stop
	OnConnStop func(conn ziface.IConnection)
}

func (s *Server) Start() {
	fmt.Printf("[Start] Server Listenner at IP : %s, Port : %d, is starting\n", s.IP, s.Port)

	go func() {
		s.MsgHandler.StartWorkerPool()

		// get tcp addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error : ", err)
			return
		}

		// listen
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen IP : ", s.IP, "error : ", err)
			return
		}

		fmt.Println("start zinx server ", s.Name, "successful, listening...")
		var cid uint32
		cid = 0

		// wait for connect
		for {
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept error : ", err)
				continue
			}

			if s.ConnMgr.Len() >= utils.GlobalObject.MaxConn {
				fmt.Println("====> The number of connection exceeds maximum!")
				conn.Close()
				continue
			}

			dealConn := NewConnection(s, conn, cid, s.MsgHandler)
			cid++

			go dealConn.Start()
		}
	}()

}

func (s *Server) Stop() {
	fmt.Println("[Stop] Zinx Server Name : ", s.Name)
	s.ConnMgr.ClearConn()
}

func (s *Server) Serve() {
	s.Start()

	select {}
}

func (s *Server) AddRouter(msgID uint32, router ziface.IRouter) {
	s.MsgHandler.AddRouter(msgID, router)
	fmt.Println("add router successful!")
}

func (s *Server) GetConnMgr() ziface.IConnManager {
	return s.ConnMgr
}

func (s *Server) SetOnConnStart(hookFunc func(connection ziface.IConnection)) {
	s.OnConnStart = hookFunc
}

func (s *Server) SetOnConnStop(hookFunc func(connection ziface.IConnection)) {
	s.OnConnStop = hookFunc
}

func (s *Server) CallOnConnStart(conn ziface.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println("----> Call OnConnStart()...")
		s.OnConnStart(conn)
	}
}

func (s *Server) CallOnConnStop(conn ziface.IConnection) {
	if s.OnConnStop != nil {
		fmt.Println("----> Call OnConnStop()...")
		s.OnConnStop(conn)
	}
}

/*
	initialize server
*/
func NewServer(name string) ziface.Iserver {
	s := &Server{
		Name:       utils.GlobalObject.Name,
		IPVersion:  "tcp4",
		IP:         utils.GlobalObject.Host,
		Port:       utils.GlobalObject.TcpPort,
		MsgHandler: NewMsgHandle(),
		ConnMgr:    NewConnManager(),
	}

	return s
}
