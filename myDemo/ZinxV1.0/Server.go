package main

import "zinx/znet"

func main() {
	s := znet.NewServer("[zinx V1.0]")
	s.Serve()
}