package models

import (
	"github.com/google/uuid"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:char(36);uniqueIndex;primaryKey"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
}
