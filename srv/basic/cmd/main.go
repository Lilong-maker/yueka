package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	__ "yueka/proto"
	_ "yueka/srv/basic/init"
	"yueka/srv/handler/service"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	__.RegisterStreamGreeterServer(s, &service.Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
