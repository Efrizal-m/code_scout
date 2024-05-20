package lottery

import (
	"context"
	"math/rand"
	"time"

	pb "github.com/Efrizal-m/lottery-club/pb"
	"gorm.io/gorm"
)

type LotteryServiceServer struct {
	pb.UnimplementedLotteryServiceServer
	db *gorm.DB
}

func NewLotteryServiceServer(db *gorm.DB) *LotteryServiceServer {
	return &LotteryServiceServer{db: db}
}

func (s *LotteryServiceServer) EnterLottery(ctx context.Context, req *pb.LotteryRequest) (*pb.LotteryResponse, error) {
	rand.Seed(time.Now().UnixNano())
	win := rand.Intn(2) == 0 // 50% chance to win
	message := "You lost the lottery"
	if win {
		message = "Congratulations, you won the lottery!"
	}
	return &pb.LotteryResponse{
		Win:     win,
		Message: message,
	}, nil
}
