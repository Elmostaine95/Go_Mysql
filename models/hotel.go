package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Hotel struct {
	gorm.Model
	ID    uuid.UUID `gorm:"type:char(36);uniqueIndex;primaryKey;autoIncrement:false"`
	Hotel_name string    `json:"hotel_name"`
	Location   string    `json:"location"`
}
