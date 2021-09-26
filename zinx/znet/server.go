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
	// router
	Router ziface.IRouter
}

func (s* Server) Start() {
	fmt.Printf("[Start] Server Listenner at IP : %s, Port : %d, is starting\n", s.IP, s.Port)

	go func() {
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

			dealConn := NewConnection(conn, cid, s.Router)
			cid++

			go dealConn.Start()
		}
	}()

	
}

func (s* Server) Stop() {

}

func (s* Server) Serve() {
	s.Start()

	select{}
}

func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
	fmt.Println("add router successful!")
}

/* 
	initialize server
*/
func NewServer(name string) ziface.Iserver {
	s := &Server {
		Name : utils.GlobalObject.Name,
		IPVersion : "tcp4",
		IP : utils.GlobalObject.Host,
		Port : utils.GlobalObject.TcpPort,
		Router : nil,
	}

	return s
}
