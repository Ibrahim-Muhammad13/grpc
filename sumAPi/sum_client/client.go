package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Ibrahim-Muhammad13/sum/sumpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("hello form client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println("error with client connection")
	}
	defer cc.Close()
	c := sumpb.NewSumNumbersClient(cc)

	req := &sumpb.SumRequest{
		FirstNumber:  10,
		SecondNumber: 3,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		fmt.Println("error calling sum service")
	}
	log.Printf("the sum equals: %v", res)
}
