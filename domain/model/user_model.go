package model

import (
	"github.com/uptrace/bun"
)

type UserID string

type User struct {
	bun.BaseModel `bun:"table:user"`
	ID            UserID
	Name          string
	Email         string
	Password      string
	TimeStamp
}
