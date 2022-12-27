package main

import (
	"context"
	"fmt"
	"net"

	"github.com/Ibrahim-Muhammad13/sum/sumpb"

	"google.golang.org/grpc"
)

type server struct {
	sumpb.UnimplementedSumNumbersServer
}

func (*server) Sum(ctx context.Context, req *sumpb.SumRequest) (*sumpb.SumResponse, error) {
	fmt.Printf("sum service invoked %v", req)
	no1 := req.GetFirstNumber()
	no2 := req.GetSecondNumber()
	sum := no1 + no2
	res := &sumpb.SumResponse{
		Result: sum,
	}
	return res, nil
}
func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		fmt.Println("error while listing to port")
	}
	s := grpc.NewServer()
	sumpb.RegisterSumNumbersServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		fmt.Println("faild to serve")
	}
}
