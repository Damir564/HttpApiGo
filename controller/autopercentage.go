package controller

import (
	"net/http"

	"github.com/Damir564/HttpApiGo/config"
	"github.com/Damir564/HttpApiGo/models"
	"github.com/gin-gonic/gin"
)

func getUserSequence(c *gin.Context, percentage uint, segmentSlug string) {
	var allUsersAmount int64
	var users []models.User
	config.DB.Table("users").Count(&allUsersAmount)
	limitUserAmount := int((float64(percentage) / 100.0) * float64(allUsersAmount))
	if limitUserAmount == 0 {
		c.JSON(http.StatusBadRequest, "low percentage amount")
		return
	}
	if err := config.DB.Table("users").Order("random()").Limit(limitUserAmount).Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, &users)
		return
	}
	userSegmentsSlice := make([]models.UserSegments, 0)
	for _, v := range users {
		bindMessage := BindMessage{SegmentsAdd: []string{segmentSlug}, UserId: v.ID}
		// userSegmentsSlice := bindMessage.Add()
		userSegmentsSlice = append(userSegmentsSlice, bindMessage.Add(c)[0])
	}
	if len(userSegmentsSlice) != 0 {
		config.DB.Save(&userSegmentsSlice)
	}
	c.JSON(http.StatusOK, &userSegmentsSlice)
}
