package main

import (
	"fmt"
	"io"
	"net"
	"time"
	"zinx/znet"
)

func main() {
	fmt.Println("client start")

	time.Sleep(1 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client start error, exit!")
		return
	}

	for {
		dp := znet.NewDataPack()

		binaryMsg, err := dp.Pack(znet.NewMsgPackage(1, []byte("zinx client test message")))
		if err != nil {
			fmt.Println("pack error : ", err)
			return
		}

		if _, err := conn.Write(binaryMsg); err != nil {
			fmt.Println("write error : ", err)
			return
		}

		binaryHead := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, binaryHead); err != nil {
			fmt.Println("read head error : ", err)
			break
		}

		msgHead, err := dp.Unpack(binaryHead)
		if err != nil {
			fmt.Println("client unpack msg error : ", err)
			break
		}

		msg := msgHead.(*znet.Message)

		if msgHead.GetDataLen() > 0 {
			msg.Data = make([]byte, msg.GetDataLen())
		}

		if _, err := io.ReadFull(conn, msg.Data); err != nil {
			fmt.Println("read msg data error : ", err)
			return
		}

		fmt.Println("recv server msg : Id = ", msg.Id, ", len = ", msg.DataLen, ", data = ", string(msg.Data))

		time.Sleep(1 * time.Second)
	}
}
