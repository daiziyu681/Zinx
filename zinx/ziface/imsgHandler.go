package ziface

type IMsgHandle interface {
	DoMsgHandle(request IRequest)

	AddRouter(msgID uint32, router IRouter)
}
