package znet

import (
	"fmt"
	"net"
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

		// wait for connect
		for {
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept error : ", err)
				continue
			}

			// do something
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("recv buf error : ", err)
						continue
					}
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write back buf error : ", err)
						continue
					}
				}
			}()
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