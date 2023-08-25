package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID       uint      `gorm:"primarykey"`
	Segments []Segment `json:"-" gorm:"many2many:user_segments"`
}

type Segment struct {
	ID   uint   `gorm:"primarykey"`
	Slug string `json:"slug"`
}

type UserSegments struct {
	UserID    uint `json:"user_id" gorm:"primarykey"`
	SegmentID uint `json:"segment_id" gorm:"primarykey"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt // `gorm:"index"`
}
