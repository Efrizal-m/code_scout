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

	// Prepare the request to enter the lottery
	enterReq := &pb.LotteryRequest{
		UserId: "user123", // Example user ID
	}

	// Set a timeout for the request
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Send the request and receive the response
	enterRes, err := client.EnterLottery(ctx, enterReq)
	if err != nil {
		log.Fatalf("could not enter lottery: %v", err)
	}

	// Print the enter lottery response
	log.Printf("Lottery Response: win=%v, message=%s", enterRes.GetWin(), enterRes.GetMessage())

	// Prepare the request to get lottery entries
	entriesReq := &pb.LotteryEntriesRequest{
		UserId: "user123",
	}

	// Send the request to get lottery entries and receive the response
	entriesRes, err := client.GetLotteryEntries(ctx, entriesReq)
	if err != nil {
		log.Fatalf("could not get lottery entries: %v", err)
	}

	// Print the lottery entries response
	for _, entry := range entriesRes.GetEntries() {
		log.Printf("Lottery Entry: win=%v, message=%s", entry.GetWin(), entry.GetMessage())
	}
}
