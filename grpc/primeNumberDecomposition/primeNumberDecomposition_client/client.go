package main

import (
	"context"
	"fmt"
	"log"
	"google.golang.org/grpc"
	primeNumberDecomposition "../primeNumberDecomposition_pb"
	"io"
)

func main() {
	fmt.Println("Hello I'm the client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()
	
	c:= primeNumberDecomposition.NewPrimeNumberDecompositionServiceClient(cc)
	doServerStreaming(c)
	
}


func doServerStreaming(c primeNumberDecomposition.PrimeNumberDecompositionServiceClient) {
	fmt.Println("Starting server streaming rpc")
	req := &primeNumberDecomposition.PrimeNumberDecompositionRequest {
		PrimeNumberDecomposition: &primeNumberDecomposition.PrimeNumberDecomposition {
			Num : 120,
		},
	}
	resStream, err := c.PrimeNumberDecompositionStream(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling PrimeNumberDecompositionStream RPC: %v", err)
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
		log.Printf("Response from PrimeNumberDecompositionStream: %v", msg.GetResult())
	}
}