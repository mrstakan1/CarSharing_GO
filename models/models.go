package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"-"`
	IsAdmin   bool      `json:"is_admin"`
	CreatedAt time.Time `json:"created_at"`
}

type Car struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Make      string    `json:"make"`
	Model     string    `json:"model"`
	Year      int       `json:"year"`
	Location  string    `json:"location"`
	Available bool      `json:"available"`
	CreatedAt time.Time `json:"created_at"`
}

type Booking struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `json:"user_id"`
	CarID     uint      `json:"car_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	CreatedAt time.Time `json:"created_at"`

	User User `gorm:"foreignKey:UserID"`
	Car  Car  `gorm:"foreignKey:CarID"`
}
