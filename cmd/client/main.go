package main

import (
	"context"
	"log"
	"time"

	"github.com/Efrizal-m/lottery-club/pb"
	"google.golang.org/grpc"
)

func main() {
	// Connect to the gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewLotteryServiceClient(conn)

	// Prepare the request
	req := &pb.LotteryRequest{
		UserId: "user123", // Example user ID
	}

	// Set a timeout for the request
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Send the request and receive the response
	res, err := client.EnterLottery(ctx, req)
	if err != nil {
		log.Fatalf("could not enter lottery: %v", err)
	}

	// Print the response
	log.Printf("Lottery Response: win=%v, message=%s", res.GetWin(), res.GetMessage())
}
