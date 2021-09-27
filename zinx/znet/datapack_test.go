package znet

import (
	"testing"
	"net"
	"fmt"
	"io"
)

//unit test : annotation GlobalObject.Reload() in utils/globalobj.go before run "go test"
func TestDataPack(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("server listen err : ", err)
		return
	}

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("server accept error : ", err)
			}

			go func(conn net.Conn) {
				dp := NewDataPack()
				for {
					headData := make([]byte, dp.GetHeadLen())
					_, err := io.ReadFull(conn, headData)
					if err != nil {
						fmt.Println("read head error")
						break
					}

					msgHead, err := dp.Unpack(headData)
					if err != nil {
						fmt.Println("server unpack error : ", err)
						return
					}
					if msgHead.GetDataLen() > 0 {
						msg := msgHead.(*Message)
						msg.Data = make([]byte, msg.GetDataLen())
						_, err := io.ReadFull(conn, msg.Data) 
						if err != nil {
							fmt.Println("server unpack data error : ", err)
							return
						}
						fmt.Println("---> recv MsgId : ", msg.Id, ", dataLen : ", msg.DataLen, ", Data = ", string(msg.Data))
					}
				}
			}(conn)
		}
	}()




	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client dial error : ", err)
		return
	}

	dp := NewDataPack()
	msg1 := &Message {
		Id : 1,
		DataLen : 4,
		Data : []byte {'z', 'i', 'n', 'x'},
	}
	sendData1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("client pack msg error : ", err)
		return
	}

	msg2 := &Message {
		Id : 2,
		DataLen : 7,
		Data : []byte {'n', 'i', 'h', 'a', 'o', '!', '!'},
	}
	sendData2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("client pack msg error : ", err)
		return
	}

	sendData1 = append(sendData1, sendData2...)

	conn.Write(sendData1)

	select {}
}