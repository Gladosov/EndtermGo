package main

import (
	"context"
	"log"
	"net"
	"fmt"
	greetpb "../greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (*server) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {
	fmt.Printf("GreetEveryone function has invoked with a streaming request\n")
	
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream %v", err)
			return err
		}
		firstName := req.GetGreeting().GetFirstName()
		result := "Hello " + firstName + "! "
		senderErr := stream.Send(&greetpb.GreetEveryoneResponse {
			Result: result,
		})
		if senderErr != nil {
			log.Fatalf("Error while sending data to cclient: %v\n", err)
			return err
		}
	}
}

func (*server) GreetWithDeadline(ctx context.Context, req *greetpb.GreetWithDeadlineRequest) (*greetpb.GreetWithDeadlineResponse, error) {
	fmt.Printf("GreetWithDeadline invoked\n")
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			fmt.Println("The client canceled the request!")
			return nil, status.Error(codes.Canceled, "the client canceled the request")
		}
		time.Sleep(1 * time.Second)
	}
	firstname := req.GetGreeting().GetFirstName()

	result := "Hello " + firstname

	res := &greetpb.GreetWithDeadlineResponse{
		Result : result,
	}

	return res, nil
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

