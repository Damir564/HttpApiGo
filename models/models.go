package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID       uint      `json:"id" gorm:"primarykey"`
	Segments []Segment `json:"-" gorm:"many2many:user_segments;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Segment struct {
	ID             uint   `gorm:"primarykey"`
	Slug           string `json:"slug" gorm:"uniqueIndex"`
	AutoPercentage uint   `json:"auto_percentage"`
}

type UserSegments struct {
	UserID    uint `gorm:"primarykey;column:user_id"`
	SegmentID uint `gorm:"primarykey;column:segment_id"`
	TTL       time.Duration
	CreatedAt time.Time `json:"created_at"`
}

type History struct {
	ID          uint      `json:"id" gorm:"primarykey;column:id"`
	UserID      uint      `json:"user_id" gorm:"column:user_id"`
	SegmentID   uint      `json:"segment_id" gorm:"column:segment_id"`
	SegmentSlug string    `json:"segment_slug" gorm:"column:segment_slug"`
	Operation   string    `json:"operation" gorm:"column:operation"`
	Timestamp   time.Time `json:"timestamp" gorm:"primarykey;column:timestamp"`
}

func (v *UserSegments) AfterFind(tx *gorm.DB) (err error) {
	currentTime := time.Now()
	fmt.Printf("CreatedAt: %v; ExpireTime: %v; CurrentTime: %v", v.CreatedAt, v.CreatedAt.Add(v.TTL), time.Now())
	if currentTime.After(v.CreatedAt.Add(v.TTL)) && v.TTL != 0 {
		err = tx.Delete(v).Error
	}
	// fmt.Println(v)
	return err
}

// type Duration struct {
// 	time.Duration
// }

// func (d Duration) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(d.String())
// }

// func (d *Duration) UnmarshalJSON(b []byte) error {
// 	var v interface{}
// 	if err := json.Unmarshal(b, &v); err != nil {
// 		return err
// 	}
// 	switch value := v.(type) {
// 	case float64:
// 		d.Duration = time.Duration(value)
// 		return nil
// 	case string:
// 		var err error
// 		d.Duration, err = time.ParseDuration(value)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	default:
// 		return errors.New("invalid duration")
// 	}
// }
