package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Ibrahim-Muhammad13/prime/pb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedPrimeServiceServer
}

func (*server) Prime(req *pb.PrimeRequest, stream pb.PrimeService_PrimeServer) error {
	number := req.GetNumber()
	divisor := int64(2)
	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&pb.PrimeResponse{
				PrimeFactor: divisor,
			})
			number = number / 2
		} else {
			divisor++
		}

		// time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func main() {
	fmt.Println("hello world!")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("faild to listen %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPrimeServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("faild to serve %v", err)
	}

}
