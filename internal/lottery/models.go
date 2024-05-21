package lottery

import (
	"gorm.io/gorm"
)

// LotteryEntry represents a user's entry into the lottery.
type LotteryEntry struct {
	gorm.Model
	UserID string `gorm:"index"`
	Win    bool
}
