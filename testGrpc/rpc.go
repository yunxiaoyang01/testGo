package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"runtime"
	"testMgo/testGrpc/proto"
)

const (
	port = "41005"
)

type Data struct{}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//起服务
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	data:=&Data{}
	proto.RegisterSayHelloServiceServer(s,data)
	log.Printf("grpc server in: %s", port)
	s.Serve(lis)

}

func (t *Data) SayHello(ctx context.Context,in *proto.SayHelloRequest) (result *proto.SayHelloResponse, err error){
	return &proto.SayHelloResponse{
		Result:[] byte("hello :"+string(in.Name)),
	},nil
}

func (t *Data) GetResult( proto.SayHelloService_GetResultServer)(result *proto.Result,err error)  {
	return &proto.Result{
		Sum: 1,
		Avg: 1,
		Cnt: 100,
	}, nil
}