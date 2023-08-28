package controller

import (
	"fmt"
	"net/http"

	"github.com/Damir564/HttpApiGo/config"
	"github.com/Damir564/HttpApiGo/models"
	"github.com/gin-gonic/gin"
)

type SegmentChange struct {
	Slug    string `json:"slug"`
	NewSlug string `json:"newSlug"`
}

func GetSegments(c *gin.Context) {
	segments := []models.Segment{}
	config.DB.Find(&segments)
	c.JSON(http.StatusOK, &segments)
}

func CreateSegment(c *gin.Context) {
	var segment models.Segment
	// c.BindJSON(&segment)
	if err := c.BindJSON(&segment); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	var tempSegment models.Segment
	if err := config.DB.Table("segments").Where("slug = ?", segment.Slug).First(&tempSegment).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Segment with such slug already exists"})
		return
	}
	config.DB.Create(&segment)
	if segment.AutoPercentage > 0 {
		getUserSequence(c, segment.AutoPercentage, segment.Slug)
		// c.JSON(http.StatusOK, gin.H{"UsersAmount:": getUserSequence(segment.AutoPercentage)})
	}
	// c.JSON(http.StatusOK, &segment)
}

func UpdateSegment(c *gin.Context) {
	var segmentChange SegmentChange
	var segment models.Segment
	c.BindJSON(&segmentChange)
	if err := config.DB.Where("slug = ?", segmentChange.Slug).First(&segment).Error; err != nil {
		fmt.Printf("segment with slug \"%s\" mot found\n", segment.Slug)
	} else {
		segment.Slug = segmentChange.NewSlug
		config.DB.Save(&segment)
		c.JSON(http.StatusOK, &segment)
	}
}

func DeleteSegment(c *gin.Context) {
	var segment models.Segment
	c.BindJSON(&segment)
	if err := config.DB.Where("slug = ?", segment.Slug).Delete(&segment).Error; err != nil {
		fmt.Printf("segment with slug \"%s\" mot found\n", segment.Slug)
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("segment with slug = %s", segment.Slug)})
}
