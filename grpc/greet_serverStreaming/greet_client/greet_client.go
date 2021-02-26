package main

import (
	"context"
	"fmt"
	"log"
	"google.golang.org/grpc"
	greetpb "../greetpb"
	"io"
)

func main() {
	fmt.Println("Hello I'm the client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()
	
	c:= greetpb.NewGreetServiceClient(cc)

	doServerStreaming(c)
	
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