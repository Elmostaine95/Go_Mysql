package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:char(36);uniqueIndex;primaryKey;autoIncrement:false"`
	HotelID   uuid.UUID
	UserID    uuid.UUID
	RoomID    uuid.UUID
	Check_in  time.Time `json:"check_in"`
	Check_out time.Time `json:"check_out"`
	Hotel     Hotel     `gorm:"foreignKey:HotelID"`
	User      User      `gorm:"foreignKey:UserID"`
	Room      Room      `gorm:"foreignKey:RoomID"`
}
