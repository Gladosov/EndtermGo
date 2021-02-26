package main

import (
	"context"
	"log"
	"net"
	"fmt"
	calculatorpb "../calculatorpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc"
	"math"
)

type server struct {}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	//this is how we get information from request
	x := req.GetSum().GetX()
	y := req.GetSum().GetY()

	result := x + y

	res := &calculatorpb.SumResponse{
		Result : result,
	}

	return res, nil
}

func (*server) SquareRoot(c context.Context, req *calculatorpb.SquareRootRequest) (*calculatorpb.SquareRootResponse, error) {
	fmt.Println("Recieved SquareRoot RPC")
	number := req.GetNumber()
	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Recieved a negative number: %v", number),
		)
	}
	return &calculatorpb.SquareRootResponse{
		NumberRoot : math.Sqrt(float64(number)),
	}, nil
}

func main() {
	fmt.Println("Calculator")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	calculatorpb.RegisterCalculatorServiceServer(s, &server{})
	
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
		
}