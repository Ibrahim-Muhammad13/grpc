package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/Ibrahim-Muhammad13/prime/pb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("hello form client ")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error not connect %v", err)
	}
	defer cc.Close()
	c := pb.NewPrimeServiceClient(cc)
	//doUnnary(c)
	doServerStreaming(c)
}

func doServerStreaming(c pb.PrimeServiceClient) {
	fmt.Println("starting streaming server")

	req := &pb.PrimeRequest{
		Number: 12,
	}
	resStream, err := c.Prime(context.Background(), req)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("error  %v", err)
		}
		log.Printf("respnse  %v", msg.GetPrimeFactor())
	}
}
