package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Ibrahim-Muhammad13/avg/pb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("hello form clientfmt")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error not connect %v", err)
	}
	defer cc.Close()
	c := pb.NewAvgServercieClient(cc)
	doClientStreaming(c)
}

func doClientStreaming(c pb.AvgServercieClient) {
	//	fmt.Println("hclient")

	reqests := []*pb.AvgRequest{
		&pb.AvgRequest{
			Number: 1,
		},
		&pb.AvgRequest{
			Number: 2,
		},
		&pb.AvgRequest{
			Number: 3,
		},
		&pb.AvgRequest{
			Number: 4,
		},
	}

	stream, err := c.Avg(context.Background())
	if err != nil {
		log.Fatalf("%v", err)
	}
	for _, req := range reqests {
		fmt.Printf("Sending requests %v \n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Printf("response %v", res)
}
