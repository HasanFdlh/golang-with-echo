// internal/model/user.go
package model

import (
	"fmt"
	"time"
)

type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Email     string    `gorm:"uniqueIndex" json:"email"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"-"`

	// JSON output formatted
	CreatedAtStr JSONTime `gorm:"-" json:"created_at"`
	UpdatedAtStr JSONTime `gorm:"-" json:"updated_at"`
}

type UserRequest struct {
	Name  string `json:"name" validate:"required,alphaSpace"`
	Email string `json:"email" validate:"required,email"`
}
