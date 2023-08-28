package controller

import (
	"net/http"

	"github.com/Damir564/HttpApiGo/config"
	"github.com/Damir564/HttpApiGo/models"
	"github.com/gin-gonic/gin"
)

// type UserActiveSegments struct {
// 	ID       uint             `json:"user_id"`
// 	Segments []models.Segment `json:"segments"`
// }

type UserRequest struct {
	ID uint `json:"user_id"`
}

func GetBinds(c *gin.Context) {
	//var segmentIds []uint
	var segments []models.Segment
	var user UserRequest
	c.BindJSON(&user)
	// userSegments := []models.UserSegments{}
	if err := config.DB.Table("user_segments").
		Where("user_id = ? AND deleted_at IS NULL", user.ID).
		Order("segment_id asc").Joins("join segments on segments.id = user_segments.segment_id").Select("segments.id", "segments.slug").
		Find(&segments).Error; err != nil {
		c.JSON(http.StatusBadRequest, &segments)
	} // } else {
	// 	c.JSON(http.StatusOK, &segments)
	// 	var userActiveSegments UserActiveSegments
	// 	userActiveSegments.ID = user.ID
	// 	if err := config.DB.Table("segments").Where("id IN ?", segmentIds).Find(&userActiveSegments.Segments).Error; err != nil {
	// 		c.JSON(http.StatusBadRequest, &userActiveSegments.Segments)
	// 	}
	// 	c.JSON(http.StatusOK, &userActiveSegments)
	// }
	c.JSON(http.StatusOK, &segments)
}
