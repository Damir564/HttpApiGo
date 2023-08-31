package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Damir564/HttpApiGo/config"
	"github.com/Damir564/HttpApiGo/models"
	"github.com/gin-gonic/gin"
)

type BindMessage struct {
	SegmentsAdd    []string `json:"segmentsAdd"`
	SegmentsRemove []string `json:"segmentsRemove"`
	UserId         uint     `json:"user_id"`
	TTL            Duration `json:"ttl"`
}

type Duration struct {
	time.Duration
}

func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case float64:
		d.Duration = time.Duration(value)
		return nil
	case string:
		var err error
		d.Duration, err = time.ParseDuration(value)
		if err != nil {
			return err
		}
		return nil
	default:
		return errors.New("invalid duration")
	}
}

// Bind godoc
// @Summary Bind/Unbind User with Segments
// @Description (un)binds user with segments using segment slugs and user id
// @Tags user_segments
// @Produce json
// @Param segmentsAdd formData array false "array of segment's slugs that you want to bind"
// @Param segmentsRemove formData array false "array of segment's slugs that you want to unbind"
// @Param user_id formData int true "Id of the user you want to (un)bind with segments"
// @Success 200
// @Router /bind [post]
func Bind(c *gin.Context) {
	//
	var bindMessage BindMessage
	c.BindJSON(&bindMessage)
	fmt.Println(bindMessage.TTL)

	tmpUserSegments := make([]models.UserSegments, 0)
	config.DB.Model(&models.UserSegments{}).Find(&tmpUserSegments)

	userSegmentsSliceAdd := bindMessage.Add(c)
	if len(userSegmentsSliceAdd) != 0 {
		config.DB.Save(&userSegmentsSliceAdd)
	}
	userSegmentsSliceRemove := make([]models.UserSegments, 0)
	if len(bindMessage.SegmentsRemove) != 0 {
		userSegmentsSliceRemove = bindMessage.Remove(c)
	}
	// c.JSON(http.StatusOK, &bindMessage)
	c.JSON(http.StatusOK, gin.H{"Added": &userSegmentsSliceAdd, "Removed": &userSegmentsSliceRemove})
}

// func GetSegmentIdByName(s string) (segment models.Segment) {
// 	if err := config.DB.Where("slug = ?", s).First(&segment).Error; err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			fmt.Printf("incorrect name \"%s\" in SegmentsAdd\n", s)
// 		}
// 		// panic(err)
// 	}
// 	return segment
// }

func (bm *BindMessage) Add(c *gin.Context) []models.UserSegments {
	userSegmentsSlice := make([]models.UserSegments, 0)
	for _, v := range bm.SegmentsAdd {
		var segment models.Segment
		if r := config.DB.Model(&models.Segment{}).Where("slug = ?", v).Limit(1).Find(&segment); r.RowsAffected == 0 {
			fmt.Printf("incorrect name \"%s\" in segmentsAdd\n", v)
			c.String(http.StatusBadRequest, fmt.Sprintf("incorrect name \"%s\" in SegmentsAdd\n", v))
		} else {
			var userSegments models.UserSegments
			var user models.User
			if r := config.DB.Model(&models.User{}).Where("id = ?", bm.UserId).Limit(1).Find(&user); r.RowsAffected == 0 {
				fmt.Printf("unregistered user id \"%d\" in user_id\n", bm.UserId)
				c.String(http.StatusBadRequest, fmt.Sprintf("unregistered user id \"%d\" in user_id\n", bm.UserId))
			}
			if r := config.DB.Model(&models.UserSegments{}).Limit(1).Find(&userSegments, "user_id = ? AND segment_id = ?", bm.UserId, segment.ID); r.RowsAffected > 0 {
				fmt.Printf("binding for User \"%d\" and Segment \"%d\" \"%s\"already exist\n", bm.UserId, segment.ID, segment.Slug)
				c.String(http.StatusBadRequest, fmt.Sprintf("binding for User \"%d\" and Segment \"%d\" \"%s\"already exist\n", bm.UserId, segment.ID, segment.Slug))
			} else {
				userSegments.UserID = bm.UserId
				userSegments.SegmentID = segment.ID
				userSegments.TTL = bm.TTL.Duration
				userSegments.CreatedAt = time.Now()
				userSegmentsSlice = append(userSegmentsSlice, userSegments)
				var history models.History = models.History{UserID: bm.UserId,
					SegmentID:   segment.ID,
					SegmentSlug: segment.Slug,
					Operation:   "Add",
					Timestamp:   time.Now(),
				}
				config.DB.Model(&models.History{}).Save(&history)
			}
			fmt.Println(userSegments)
		}
	}
	return userSegmentsSlice
}

func (bm *BindMessage) Remove(c *gin.Context) []models.UserSegments {
	userSegmentsSlice := make([]models.UserSegments, 0)
	for _, v := range bm.SegmentsRemove {
		var userSegments models.UserSegments
		var segment models.Segment
		if r := config.DB.Model(&models.Segment{}).Where("slug = ?", v).Limit(1).Find(&segment); r.RowsAffected == 0 {
			fmt.Printf("incorrect name \"%s\" in SegmentsRemove\n", v)
			c.String(http.StatusBadRequest, fmt.Sprintf("incorrect name \"%s\" in SegmentsRemove\n", v))
		} else {
			var user models.User
			if r := config.DB.Model(&models.User{}).Where("id = ?", bm.UserId).Limit(1).Find(&user); r.RowsAffected == 0 {
				fmt.Printf("unregistered user id \"%d\" in user_id\n", bm.UserId)
				c.String(http.StatusBadRequest, fmt.Sprintf("unregistered user id \"%d\" in user_id\n", bm.UserId))
			} else if r := config.DB.Model(&models.UserSegments{}).Where("user_id = ? AND segment_id = ?", user.ID, segment.ID).Limit(1).Find(&userSegments); r.RowsAffected == 0 {
				fmt.Printf("error deleting segment %d from user %d\n", segment.ID, bm.UserId)
				c.String(http.StatusBadRequest, fmt.Sprintf("error deleting segment %d from user %d\n", segment.ID, bm.UserId))
			} else {
				userSegmentsSlice = append(userSegmentsSlice, userSegments)
				var history models.History = models.History{UserID: bm.UserId,
					SegmentID:   segment.ID,
					SegmentSlug: segment.Slug,
					Operation:   "Remove",
					Timestamp:   time.Now().UTC(),
				}
				config.DB.Model(&models.History{}).Save(&history)
				config.DB.Model(&models.UserSegments{}).Where("user_id = ? AND segment_id = ?", user.ID, segment.ID).Delete(&userSegments)
			}
		}
	}
	return userSegmentsSlice
}
