package main

import(
	"fmt"
	"zinx/znet"
	"zinx/ziface"
)

// ping test self define router
type PingRouter struct {
	znet.BaseRouter
}

func (this *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("call router PreHandle..")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping.."))
	if err != nil {
		fmt.Println("call back before ping error")
	}
}

func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("call router Handle..")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping..."))
	if err != nil {
		fmt.Println("call back ping ping error")
	}
}

func (this *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("call router PostHandle..")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping.."))
	if err != nil {
		fmt.Println("call back after ping error")
	}
}

func main() {
	s := znet.NewServer("[zinx V1.0]")
	s.AddRouter(&PingRouter{})
	s.Serve()
}