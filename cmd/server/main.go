package main

import (
	"log"
	"net"

	"github.com/Efrizal-m/lottery/internal/db"
	"github.com/Efrizal-m/lottery/internal/lottery"
	"github.com/Efrizal-m/lottery/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	database, err := db.NewDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	grpcServer := grpc.NewServer()
	lotteryService := lottery.NewLotteryServiceServer(database)
	pb.RegisterLotteryServiceServer(grpcServer, lotteryService)

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
