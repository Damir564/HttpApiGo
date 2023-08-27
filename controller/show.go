package controller

import (
	"net/http"

	"github.com/Damir564/HttpApiGo/config"
	"github.com/Damir564/HttpApiGo/models"
	"github.com/gin-gonic/gin"
)

type UserActiveSegments struct {
	ID       uint             `json:"user_id"`
	Segments []models.Segment `json:"segments"`
}

func GetBinds(c *gin.Context) {
	var segmentIds []uint
	var user models.User
	c.BindJSON(&user)
	// userSegments := []models.UserSegments{}
	if err := config.DB.Table("user_segments").Where("user_id = ?", user.ID).Select("segment_id").Order("segment_id asc").Find(&segmentIds).Error; err != nil {
		c.JSON(http.StatusBadRequest, &segmentIds)
	} else {
		var userActiveSegments UserActiveSegments
		userActiveSegments.ID = user.ID
		if err := config.DB.Table("segments").Where("id IN ?", segmentIds).Find(&userActiveSegments.Segments).Error; err != nil {
			c.JSON(http.StatusBadRequest, &userActiveSegments.Segments)
		}
		c.JSON(http.StatusOK, &userActiveSegments)
	}
}
