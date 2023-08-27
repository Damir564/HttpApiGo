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

func GetUsers(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(http.StatusOK, &users)
}

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	// if err := c.BindJSON(&segment); err != nil {
	// 	panic(err)
	// }
	config.DB.Create(&user)
	c.JSON(http.StatusOK, &user)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).Delete(&user)
	c.JSON(http.StatusOK, &user)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).First(&user)
	c.BindJSON(&user)
	// if err := c.BindJSON(&segment); err != nil {
	// 	panic(err)
	// }
	config.DB.Save(&user)
	c.JSON(http.StatusOK, &user)
}

// Segments

func GetSegments(c *gin.Context) {
	segments := []models.Segment{}
	config.DB.Find(&segments)
	c.JSON(http.StatusOK, &segments)
}

func CreateSegment(c *gin.Context) {
	var segment models.Segment
	c.BindJSON(&segment)
	// if err := c.BindJSON(&segment); err != nil {
	// 	panic(err)
	// }
	config.DB.Create(&segment)
	c.JSON(http.StatusOK, &segment)
}

func DeleteSegment(c *gin.Context) {
	var segment models.Segment
	if err := config.DB.Where("id = ?", c.Param("id")).Delete(&segment).Error; err != nil {

	}
	c.JSON(http.StatusOK, &segment)
}

func UpdateSegment(c *gin.Context) {
	var segment models.Segment
	config.DB.Where("id = ?", c.Param("id")).First(&segment)
	c.BindJSON(&segment)
	// if err := c.BindJSON(&segment); err != nil {
	// 	panic(err)
	// }
	config.DB.Save(&segment)
	c.JSON(http.StatusOK, &segment)
}

// Bind

type BindMessage struct {
	SegmentsAdd    []string `json:"segmentsAdd"`
	SegmentsRemove []string `json:"segmentsRemove"`
	UserId         uint     `json:"user_id"`
}

func GetSegmentIdByName(s string) (segment models.Segment) {
	if err := config.DB.Where("slug = ?", s).First(&segment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Printf("incorrect name \"%s\" in SegmentsAdd\n", s)
		}
		// panic(err)
	}
	return segment
}

func (bm *BindMessage) Add() []models.UserSegments {
	userSegmentsSlice := make([]models.UserSegments, 0)
	for _, v := range bm.SegmentsAdd {
		var userSegments models.UserSegments
		var segment models.Segment
		if err := config.DB.Where("slug = ?", v).First(&segment).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				fmt.Printf("incorrect name \"%s\" in SegmentsAdd\n", v)
			}
			// panic(err)
		} else {
			var user models.User
			if err := config.DB.Where("id = ?", bm.UserId).First(&user).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					fmt.Printf("unregistered user id \"%d\" in user_id\n", bm.UserId)
				}
				// panic(err)
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

func Bind(c *gin.Context) {
	//
	var bindMessage BindMessage
	c.BindJSON(&bindMessage)
	userSegmentsSlice := bindMessage.Add()
	if len(userSegmentsSlice) != 0 {
		config.DB.Save(&userSegmentsSlice)
	}
	bindMessage.Remove()
	// c.JSON(http.StatusOK, &bindMessage)
	c.JSON(http.StatusOK, &userSegmentsSlice)
}

func GetBinds(c *gin.Context) {
	userSegments := []models.UserSegments{}
	config.DB.Find(&userSegments)
	c.JSON(http.StatusOK, &userSegments)
}
