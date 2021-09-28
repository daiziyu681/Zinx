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

func main() {
	s := znet.NewServer("[zinx V1.0]")
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloZinxRouter{})
	s.Serve()
}
