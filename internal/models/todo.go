package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID        uint64
	UserID    uint64
	Title     string
	Detail    string
	Complete  bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
