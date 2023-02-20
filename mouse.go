package main

import (
	"fmt"
	"github.com/wuranxu/mouse-client/api"
	"github.com/wuranxu/mouse-client/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	server := grpc.NewServer()
	proto.RegisterMouseServiceServer(server, &api.MouseServiceApi{})
	listen, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("listen at: ", listen.Addr().(*net.TCPAddr).Port)

}
