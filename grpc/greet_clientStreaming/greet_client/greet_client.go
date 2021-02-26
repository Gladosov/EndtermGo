package main

import (
	"context"
	"fmt"
	"log"
	"google.golang.org/grpc"
	greetpb "../greetpb"
	"io"
	"time"
)

func main() {
	fmt.Println("Hello I'm the client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()
	
	c:= greetpb.NewGreetServiceClient(cc)

	//doServerStreaming(c)
	doClientStreaming(c)
	
}


func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting server streaming rpc")
	req := &greetpb.GreetManytimesRequest {
		Greeting: &greetpb.Greeting {
			FirstName : "Avery",
			LastName : "Wong",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GreetManyTimes RPC: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			//we reach end of stream
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream %v", err)
		}
		log.Printf("Response from GreetManyTimes: %v", msg.GetResult())
	}
}

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting client streaming rpc")

	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest {
			Greeting: &greetpb.Greeting {
				FirstName : "Avery",
				LastName : "Wong",
			},
		},
		&greetpb.LongGreetRequest {
			Greeting: &greetpb.Greeting {
				FirstName : "dan",
				LastName : "Wan",
			},
		},
		&greetpb.LongGreetRequest {
			Greeting: &greetpb.Greeting {
				FirstName : "max",
				LastName : "Lee",
			},
		},
		&greetpb.LongGreetRequest {
			Greeting: &greetpb.Greeting {
				FirstName : "Bane",
				LastName : "Anderson",
			},
		},
	}
	stream, err := c.LongGreet(context.Background()) 
	if err != nil {
		log.Fatalf("error while calling LongGreet RPC: %v", err)
	}

	for _, req := range requests {
		fmt.Println("Sending req %v", req)
		stream.Send(req)
		time.Sleep(100 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while recieving response from LongGreet &v", err)
	}
	fmt.Printf("LongGreet Response: %v\n",res) 
}