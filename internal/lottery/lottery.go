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

	// Create a new LotteryEntry and save it to the database
	entry := &LotteryEntry{
		UserID: req.GetUserId(),
		Win:    win,
	}
	if err := s.db.Create(entry).Error; err != nil {
		return nil, err
	}

	return &pb.LotteryResponse{
		Win:     win,
		Message: message,
	}, nil
}

func (s *LotteryServiceServer) GetLotteryEntries(ctx context.Context, req *pb.LotteryEntriesRequest) (*pb.LotteryEntriesResponse, error) {
	var entries []LotteryEntry
	if err := s.db.Where("user_id = ?", req.GetUserId()).Find(&entries).Error; err != nil {
		return nil, err
	}

	var responses []*pb.LotteryResponse
	for _, entry := range entries {
		responses = append(responses, &pb.LotteryResponse{
			Win:     entry.Win,
			Message: getMessage(entry.Win),
		})
	}

	return &pb.LotteryEntriesResponse{
		Entries: responses,
	}, nil
}

func getMessage(isWin bool) string {
	if isWin {
		return "Congratulations, you won the lottery!"
	}
	return "You lost the lottery"
}
