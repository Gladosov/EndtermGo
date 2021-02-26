package main

import (
	"context"
	"fmt"
	"log"
	 "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.gogrpclang.org/grpc/status"
	calculatorpb "../calculatorpb"
)

func main() {
	fmt.Println("Hello I'm the client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()
	
	c:= calculatorpb.NewCalculatorServiceClient(cc)
	//doUnary(c)
	doErrorUnary(c)
	
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting unary rpc adding 7 and -34")
	req := &calculatorpb.SumRequest{
		Sum : &calculatorpb.Sum {
			X : 7,
			Y : -34,
		},
	}
	//fmt.Printf("Created client %f", c)

	res, err := c.Sum(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling calculatorpb's Sum RPC: %v", err)
	}
	log.Printf("Response from Sum %d", res.Result)
}

func doErrorUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("calling Squareroot")
	num := int32(-7)
	req := &calculatorpb.SquareRootRequest{
		Number : num,
	}
	//fmt.Printf("Created client %f", c)

	res, err := c.SquareRoot(context.Background(), req)

	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			fmt.Println(respErr.Message())
			fmt.Println(respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("negative number!")
			}
		} else {
			log.Fatalf("Error calling SquareRoot: %v", err)
		}
	} else {
		log.Printf("Response from SquareRoot of %v: %v",num, res.GetNumberRoot())
	}
	
}