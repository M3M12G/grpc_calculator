package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_calculator/calculator/calculatorpb"
	"io"
	"log"
	"time"
)

func main() {
	fmt.Println("Calculator Service is ready...")
	conn, err := grpc.Dial("localhost:7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := calculatorpb.NewCalculatorServiceClient(conn)

	doPrimeNumberDecomposition(c)
	doComputeAverage(c)
}


func doPrimeNumberDecomposition(c calculatorpb.CalculatorServiceClient) {
	ctx := context.Background()
	request := &calculatorpb.PrimeNumberRequest{Number: 120}

	stream, err := c.PrimeNumberDecomposition(ctx, request)
	if err != nil {
		log.Fatalf("error while calling PrimeNumberDecomposition RPC %v", err)
	}
	defer stream.CloseSend()

LOOP:
	for {
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break LOOP
			}
			log.Fatalf("error while reciving from PrimeNumberDecomposition RPC %v", err)
		}
		log.Printf("response from PrimeNumberDecomposition:%v \n", res.GetResult())
	}

}

func doComputeAverage(c calculatorpb.CalculatorServiceClient) {
	ctx := context.Background()
	stream, err := c.ComputeAverage(ctx)
	numbers := []int32{2, 5, 7, 8, 9}
	if err != nil {
		log.Fatalf("error while calling ComputeAverage: %v", err)
	}
	for _, number := range numbers {
		stream.Send(&calculatorpb.AvgNumberRequest{Num: number})
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from ComputeAverage RPC: %v", err)
	}
	fmt.Printf("ComputeAverage RPC Response: %v\n", res.GetRes())
}