package main

import (
	"flag"
	"fmt"
	json "github.com/json-iterator/go"
	"github.com/wuranxu/mouse-client/api"
	"github.com/wuranxu/mouse-client/proto"
	tool "github.com/wuranxu/mouse-tool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"io/ioutil"
	"log"
	"net"
	"time"
)

var (
	ConfigFile = flag.String("config", "./config.json", "default config filepath")
)

func printBanner() {
	banner := "                                  \n   ____ ___  ____  __  __________ \n  / __ `__ \\/ __ \\/ / / / ___/ _ \\\n / / / / / / /_/ / /_/ (__  )  __/\n/_/ /_/ /_/\\____/\\__,_/____/\\___/ \n                                  "
	log.Println(banner)
}

func loadConfig(filepath string, v interface{}) error {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func main() {
	flag.Parse()
	// load config file
	var cfg MouseConfig
	err := loadConfig(*ConfigFile, &cfg)
	if err != nil {
		log.Fatal("config load failed, ", err)
	}

	// start grpc server
	server := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		Time:    8 * time.Second, // Ping the client if it is idle for 5 seconds to ensure the connection is still active
		Timeout: 2 * time.Second,
	}))
	mouse := &api.MouseServiceApi{}
	proto.RegisterMouseServiceServer(server, mouse)
	addr := "0.0.0.0:"
	if cfg.Port != 0 {
		addr = fmt.Sprintf("0.0.0.0:%d", cfg.Port)
	}
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	port := listen.Addr().(*net.TCPAddr).Port
	client, err := tool.NewEtcdClient(cfg.Etcd.Prefix, port, cfg.Etcd.Endpoints)
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
