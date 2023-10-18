package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:char(36);uniqueIndex;primaryKey;autoIncrement:false"`
	HotelID  uuid.UUID
	Capacity int    `json:"capacity"`
	Prize    string `json:"prize"`
	Hotel    Hotel  `gorm:"foreignKey:HotelID"`
}
