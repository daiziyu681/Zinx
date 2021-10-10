package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

// ping test self define router
type PingRouter struct {
	znet.BaseRouter
}

func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("call ping Handle..")
	fmt.Println("recv from client: msgId = ", request.GetMsgId(), ", data = ", string(request.GetData()))

	if err := request.GetConnection().SendMsg(200, []byte("ping...ping...ping")); err != nil {
		fmt.Println(err)
	}
}

type HelloZinxRouter struct {
	znet.BaseRouter
}

func (this *HelloZinxRouter) Handle(request ziface.IRequest) {
	fmt.Println("call Hello Handle..")
	fmt.Println("recv from client: msgId = ", request.GetMsgId(), ", data = ", string(request.GetData()))

	if err := request.GetConnection().SendMsg(201, []byte("Hello, Welcome to Zinx!!")); err != nil {
		fmt.Println(err)
	}
}

func DoConnectionBegin(conn ziface.IConnection) {
	fmt.Println("===> DoConnectionBegin is called...")
	if err := conn.SendMsg(202, []byte("DoConnection Begin")); err != nil {
		fmt.Println(err)
	}

	fmt.Println("set conn property...")
	conn.SetProperty("Name", "lucky have you hua")
	conn.SetProperty("Github", "https://github.com/daiziyu681")
}

func DoConnectionLost(conn ziface.IConnection) {
	fmt.Println("===> DoConnectionLost is called...")
	fmt.Println("conn ID = ", conn.GetConnID(), " is Lost...")

	if name, err := conn.GetProperty("Name"); err == nil {
		fmt.Println("Name = ", name)
	}
	if github, err := conn.GetProperty("Github"); err == nil {
		fmt.Println("Github = ", github)
	}
}

func main() {
	s := znet.NewServer("[zinx V1.0]")

	s.SetOnConnStart(DoConnectionBegin)
	s.SetOnConnStop(DoConnectionLost)

	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloZinxRouter{})

	s.Serve()
}
