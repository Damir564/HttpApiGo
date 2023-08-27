package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID       uint      `gorm:"primarykey"`
	Segments []Segment `json:"-" gorm:"many2many:user_segments;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Segment struct {
	ID   uint   `gorm:"primarykey"`
	Slug string `json:"slug" gorm:"uniqueIndex"`
}

type UserSegments struct {
	UserID    uint `gorm:"primarykey;"`
	SegmentID uint `gorm:"primarykey;"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt // `gorm:"index"`
}
