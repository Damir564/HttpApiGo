package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Damir564/HttpApiGo/config"
	"github.com/Damir564/HttpApiGo/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BindMessage struct {
	SegmentsAdd    []string `json:"segmentsAdd"`
	SegmentsRemove []string `json:"segmentsRemove"`
	UserId         uint     `json:"user_id"`
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
	userSegmentsSlice := bindMessage.Add(c)
	if len(userSegmentsSlice) != 0 {
		config.DB.Save(&userSegmentsSlice)
	}
	if len(bindMessage.SegmentsRemove) != 0 {
		bindMessage.Remove()
	}
	// c.JSON(http.StatusOK, &bindMessage)
	c.JSON(http.StatusOK, &userSegmentsSlice)
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
		var userSegments models.UserSegments
		var segment models.Segment
		if err := config.DB.Where("slug = ?", v).First(&segment).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				fmt.Printf("incorrect name \"%s\" in SegmentsAdd\n", v)
				c.String(http.StatusBadRequest, fmt.Sprintf("incorrect name \"%s\" in SegmentsAdd\n", v))
			}
			// panic(err)
		} else {
			var user models.User
			if err := config.DB.Where("id = ?", bm.UserId).First(&user).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					fmt.Printf("unregistered user id \"%d\" in user_id\n", bm.UserId)
					c.String(http.StatusBadRequest, fmt.Sprintf("unregistered user id \"%d\" in user_id\n", bm.UserId))
				}
				// panic(err)
			} else if config.DB.Where("user_id = ? AND segment_id = ?", bm.UserId, segment.ID).First(&userSegments).Error == nil {
				fmt.Printf("binding for User \"%d\" and Segment \"%d\" \"%s\"already exist\n", bm.UserId, segment.ID, segment.Slug)
				c.String(http.StatusBadRequest, fmt.Sprintf("binding for User \"%d\" and Segment \"%d\" \"%s\"already exist\n", bm.UserId, segment.ID, segment.Slug))
			} else {
				userSegments.UserID = bm.UserId
				userSegments.SegmentID = segment.ID
				userSegmentsSlice = append(userSegmentsSlice, userSegments)
			}
		}
	}
	return userSegmentsSlice
}

func (bm *BindMessage) Remove() {
	for _, v := range bm.SegmentsRemove {
		var userSegments models.UserSegments
		var segment models.Segment
		if err := config.DB.Where("slug = ?", v).First(&segment).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				fmt.Printf("incorrect name \"%s\" in SegmentsRemove\n", v)
			}
			// panic(err)
		} else {
			var user models.User
			if err := config.DB.Where("id = ?", bm.UserId).First(&user).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					fmt.Printf("unregistered user id \"%d\" in user_id\n", bm.UserId)
				}
				// panic(err)
			} else if err := config.DB.Where("user_id = ? AND segment_id = ?", bm.UserId, segment.ID).Delete(&userSegments).Error; err != nil {
				fmt.Printf("error deleting segment %d from user %d\n", segment.ID, bm.UserId)
				// userSegmentsSlice = append(userSegmentsSlice, userSegments)
			}
		}
	}
}
