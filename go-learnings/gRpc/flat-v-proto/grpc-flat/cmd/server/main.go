package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	sm "github.com/shiva0612/grpc-flat/api/models"

	flatbuffers "github.com/google/flatbuffers/go"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 3000))
	if err != nil {
		log.Fatalf("Falied to listen: %v", err)
	}

	codec := &flatbuffers.FlatbuffersCodec{}
	grpcServer := grpc.NewServer(grpc.ForceServerCodec(codec))
	sm.RegisterUserAPIServer(grpcServer, newServer())

	fmt.Println("starting the server...")
	getresponse()
	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
