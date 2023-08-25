package main

import (
	"douyin_relation_service/consts"
	relation "douyin_relation_service/kitex_gen/relation/relationservice"
	"log"
	"net"

	server "github.com/cloudwego/kitex/server"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", consts.RelationServiceAddr)
	if err != nil {
		panic(err)
	}

	svr := relation.NewServer(new(RelationServiceImpl), server.WithServiceAddr(addr))
	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
