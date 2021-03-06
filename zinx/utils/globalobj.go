package utils

import (
	"encoding/json"
	"io/ioutil"
	"zinx/ziface"
)

// configure zinx server via zinx.json

type GlobalObj struct {
	/*
		about Server
	*/
	TcpServer ziface.Iserver // server object

	Host string // listen IP

	TcpPort int // listen port

	Name string // server name

	/*
		about Zinx
	*/
	Version string // zinx version

	MaxConn int // server max connection

	MaxPackageSize uint32 // zinx max data package

	WorkerPoolSize uint32 // worker pool goroutine size

	MaxWorkerTaskLen uint32 // max task len in worker
}

var GlobalObject *GlobalObj

func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, g)
	if err != nil {
		panic(err)
	}
}

// provide a init func for this package
func init() {

	//default value
	GlobalObject = &GlobalObj{
		Name:             "ZinxServerApp",
		Version:          "V0.1",
		TcpPort:          8999,
		Host:             "0.0.0.0",
		MaxConn:          1000,
		MaxPackageSize:   4096,
		WorkerPoolSize:   10,
		MaxWorkerTaskLen: 1024,
	}

	GlobalObject.Reload()
}
