package db

import (
	"github.com/Efrizal-m/lottery-club/internal/lottery"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("lottery.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto migrate the LotteryEntry schema
	if err := db.AutoMigrate(&lottery.LotteryEntry{}); err != nil {
		return nil, err
	}

	return db, nil
}
