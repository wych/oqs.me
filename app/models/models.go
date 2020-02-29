package models

import (
	"time"
)

type OQSRecord struct {
	ID        uint `gorm:"primary_key"`
	Challenge string
	RealURL   string
}

type UserRecord struct {
	ID           uint `gorm:"primary_key"`
	CreatedAt    time.Time
	IP           string
	UA           string
	ForwardedFor string
	OQSRecordID  uint
	OQSRecord    OQSRecord `gorm:"association_autoupdate:false;association_autocreate:false"`
}
