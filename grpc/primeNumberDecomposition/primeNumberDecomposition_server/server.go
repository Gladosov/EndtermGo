package main

import (
	"context"
	"log"
	"net"
	"fmt"
	primeNumberDecomposition "../primeNumberDecomposition_pb"
	"google.golang.org/grpc"
	"time"
)

type server struct {}

func (*server) primeNumberDecomposition(ctx context.Context, req *primeNumberDecomposition.PrimeNumberDecompositionRequest) (*primeNumberDecomposition.PrimeNumberDecompositionResponse, error) {
	//this is how we get information from request
	num := req.GetPrimeNumberDecomposition().GetNum();

	res := &primeNumberDecomposition.PrimeNumberDecompositionResponse{
		Result : num,
	}

	return res, nil
}

func PrimeDecompose(N int32) (res []int32) {
	k := int32(2)
	for N > 1 {
		if N % k == 0 {
			res = append(res, k)
			N = N / k
		} else {
			k = k + 1
		}
	}
	return res
}
 
func (*server) PrimeNumberDecompositionStream(req *primeNumberDecomposition.PrimeNumberDecompositionRequest, stream primeNumberDecomposition.PrimeNumberDecompositionService_PrimeNumberDecompositionStreamServer ) error {
	fmt.Printf("PrimeNumberDecompositionStream has been invoked")
	num := req.GetPrimeNumberDecomposition().GetNum()
	for _, prime := range PrimeDecompose(num) {
		res := &primeNumberDecomposition.PrimeNumberDecompositionResponse {
			Result: prime,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func main() {
	fmt.Println("Hello World")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	primeNumberDecomposition.RegisterPrimeNumberDecompositionServiceServer(s, &server{})
	
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
		
}