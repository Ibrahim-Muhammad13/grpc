package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/Ibrahim-Muhammad13/avg/pb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedAvgServercieServer
}

func (*server) Avg(stream pb.AvgServercie_AvgServer) error {
	fmt.Println("server invoked")
	var result int64 = 0
	var i float64 = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Result: float64(result) / i,
			})
		}
		if err != nil {
			log.Fatalf("%v", err)
		}
		number := req.GetNumber()
		result += number
		i++
	}
}
func main() {
	fmt.Println("hello world!")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("faild to listen %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAvgServercieServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("faild to serve %v", err)
	}

}
