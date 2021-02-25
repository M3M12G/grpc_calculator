package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc_calculator/calculator/calculatorpb"
	"io"
	"log"
	"net"
	"time"
)

//Server with embedded UnimplementedGreetServiceServer
type Server struct {
	calculatorpb.UnimplementedCalculatorServiceServer
}

func(s *Server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error{
	fmt.Printf("PrimeNumberDecomposition service is invoked\n")
	prime_number := req.GetNumber()
	prime_numbers_factors := decomposePrimeNumber(prime_number)

	for i := 0; i < len(prime_numbers_factors); i++ {
		res := &calculatorpb.PrimeNumberResponse{Result: prime_numbers_factors[i]}
		if err := stream.Send(res); err != nil {
			log.Fatalf("error with responses: %v", err.Error())
		}
		time.Sleep(time.Second)
	}
	return nil
}

// private method for identification of prime number factors
func decomposePrimeNumber(num int32) []int32 {
	res_arr := []int32{}
	for{
		res_arr = append(res_arr, 2)
		num /= 2
		if num% 2 != 0{
			break
		}
	}
	var i int32 = 0
	for i = 3; i <= num*num; i+=2{
		for {
			res_arr = append(res_arr, 3)
			num /= i;
			if num% i != 0 {
				break
			}
		}
	}
	if num > 2 {
		res_arr = append(res_arr, num)
	}

	return res_arr
}


func(s *Server) ComputeAverage(stream calculatorpb.CalculatorService_ComputeAverageServer) error{
	fmt.Printf("ComputeAverage service is invoked\n")
	var sum int32 = 0
	var quantity int32 = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			var response = &calculatorpb.AvgNumberResponse{Res: float32(sum / quantity)}
			return stream.SendAndClose(response)
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}
		sum += req.GetNum()
		quantity++
	}
}


func main() {
	l, err := net.Listen("tcp", "0.0.0.0:7777")
	if err != nil {
		log.Fatalf("Failed to listen:%v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &Server{})
	log.Println("Server is running on port:7777")
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}