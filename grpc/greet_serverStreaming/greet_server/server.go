package main

import (
	"context"
	"log"
	"net"
	"fmt"
	greetpb "../greetpb"
	"google.golang.org/grpc"
	"strconv"
	"time"
)

type server struct {}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	//this is how we get information from request
	firstname := req.GetGreeting().GetFirstName()

	result := "Hello " + firstname

	res := &greetpb.GreetResponse{
		Result : result,
	}

	return res, nil
}


func (*server) GreetManyTimes(req *greetpb.GreetManytimesRequest, stream greetpb.GreetService_GreetManyTimesServer ) error {
	fmt.Printf("GreetManyTimes has been invoked")
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello " + firstName + " number " + strconv.Itoa(i)
		res := &greetpb.GreetManytimesResponse {
			Result: result,
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

	greetpb.RegisterGreetServiceServer(s, &server{})
	
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
		
}