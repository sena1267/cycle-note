package model

import (
	"github.com/uptrace/bun"
)

type MedicineID string

type Medicine struct {
	bun.BaseModel `bun:"table:medicine"`
	ID            MedicineID `bun:"id,pk"`
	UserID        UserID
	Name          string
	Note          string
	TimeStamp
}

type Medicines []Medicine
