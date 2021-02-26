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
	"io"
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

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	fmt.Printf("LongGreet has been invoked\n")
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			fmt.Printf("Sending %v\n", result)
			return stream.SendAndClose(&greetpb.LongGreetResponse {
				Result : result,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream %v", err)
		}
		firstName := req.GetGreeting().GetFirstName()
		result += "Hello " + firstName + "! "
	}
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