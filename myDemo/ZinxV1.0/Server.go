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

func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("call router Handle..")
	fmt.Println("recv from client: msgId = ", request.GetMsgId(), ", data = ", string(request.GetData()))

	if err := request.GetConnection().SendMsg(1, []byte("ping...ping...ping")); err != nil {
		fmt.Println(err)
	}
}

func main() {
	s := znet.NewServer("[zinx V1.0]")
	s.AddRouter(&PingRouter{})
	s.Serve()
}