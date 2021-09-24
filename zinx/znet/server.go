package znet

import (
	"fmt"
	"net"
	"errors"
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
}

func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	fmt.Println("[Conn Handle] CallBackToClient...")
	fmt.Printf("recv data %s\n", data)
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("Write back error : ", err)
		return errors.New("CallBackToClient error")
	}
	return nil
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

			dealConn := NewConnection(conn, cid, CallBackToClient)
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

/* 
	initialize server
*/
func NewServer(name string) ziface.Iserver {
	s := &Server {
		Name : name,
		IPVersion : "tcp4",
		IP : "0.0.0.0",
		Port : 8999,
	}

	return s
}
