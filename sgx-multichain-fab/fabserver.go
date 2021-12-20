package main

import (
	"fabricLink/fabricLink"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	pb "kunmel.com/sgx-multichain/multichain_proto"
	"log"
	"net"
)

type MultichainServiceServer struct {
	pb.MultichainServiceServer
}

func NewServer() *MultichainServiceServer {
	return &MultichainServiceServer{
	}
}
func main() {
	lis,err := net.Listen("tcp", ":2222")
	if err != nil {
		log.Fatalln("cannot create a listener at the address")
		log.Fatalln(err)
	}
	creds, err := credentials.NewServerTLSFromFile("../key/test.pem", "../key/test.key")
	if err != nil {
		grpclog.Fatalf("Failed to generate credentials %v", err)
	}
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterMultichainServiceServer(grpcServer, NewServer())
	fmt.Println("connect created")
	log.Fatalln(grpcServer.Serve(lis))
}

func (s * MultichainServiceServer) NewFabTx (request *pb.NewFabRequest, stream pb.MultichainService_NewFabTxServer) error {
	statusCode, txID := fabricLink.FabricLink(request.CCName, request.CCArg, request.FcnName)
	response := pb.CommitResponse{StatusCode: statusCode, TxID: txID}
	err := stream.Send(&response)
	if err != nil {
		return err
	}
	return nil
}