package main

import (
	"context"
	"log"

	pb "github.com/piyushpawar54/gRPC-learning/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("dogreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Clement",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greetings: %s\n", res.Result)
}
