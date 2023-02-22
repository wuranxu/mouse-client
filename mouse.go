package main

import (
	"flag"
	"fmt"
	"github.com/wuranxu/mouse-client/api"
	"github.com/wuranxu/mouse-client/proto"
	tool "github.com/wuranxu/mouse-tool"
	"google.golang.org/grpc"
	"log"
	"net"
	"strings"
)

var (
	Port      = flag.Int("port", 0, "the port mouse server started at, if you give 0, mouse will listen at random port, default is 0")
	KeyPrefix = flag.String("prefix", "mouse:node", "node prefix for etcd key, default is mouse:node")
	Endpoints = flag.String("endpoints", "127.0.0.1:2379", "etcd server endpoints, eg: 127.0.0.1:2379,127.0.0.1:12379 default is 127.0.0.1:2379")
)

func printBanner() {
	banner := "                                  \n   ____ ___  ____  __  __________ \n  / __ `__ \\/ __ \\/ / / / ___/ _ \\\n / / / / / / /_/ / /_/ (__  )  __/\n/_/ /_/ /_/\\____/\\__,_/____/\\___/ \n                                  "
	log.Println(banner)
}

func main() {
	flag.Parse()
	server := grpc.NewServer()
	mouse := &api.MouseServiceApi{}
	proto.RegisterMouseServiceServer(server, mouse)
	addr := "0.0.0.0:"
	if *Port != 0 {
		addr = fmt.Sprintf("0.0.0.0:%d", *Port)
	}
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	port := listen.Addr().(*net.TCPAddr).Port
	endpoints := strings.Split(*Endpoints, ",")
	client, err := tool.NewEtcdClient(*KeyPrefix, port, endpoints)
	if err != nil {
		log.Fatal("register to etcd error: ", err)
	}

	// set etcd client
	mouse.Client = client

	if err = client.Register(tool.Ready); err != nil {
		log.Fatal("register to etcd error: ", err)
	}
	defer client.Close()
	printBanner()
	log.Println("mouse server is listening at: ", port)
	server.Serve(listen)
}
