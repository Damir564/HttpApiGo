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

// GetSegments godoc
// @Summary Get segments array
// @Description Returns array of segments in JSON format
// @Tags segments
// @Produce json
// @Success 200
// @Router /segments [get]
func GetSegments(c *gin.Context) {
	segments := []models.Segment{}
	config.DB.Find(&segments)
	c.JSON(http.StatusOK, &segments)
}

// CreateSegment godoc
// @Summary Create Segment
// @Description Creates segment passing it's slug also has parameter for auto-binding
// @Tags segments
// @Produce json
// @Param slug formData string true "slug of the segment you want to create"
// @Param auto_percentage formData int false "percentage of users who will be automatically binded to these segments"
// @Success 200
// @Router /segment [post]
func CreateSegment(c *gin.Context) {
	var segment models.Segment
	// c.BindJSON(&segment)
	if err := c.BindJSON(&segment); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	var tempSegment models.Segment
	if r := config.DB.Model(&models.Segment{}).Where("slug = ?", segment.Slug).Limit(1).Find(&tempSegment); r.RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Segment with such slug already exists"})
		return
	}
	config.DB.Create(&segment)
	if segment.AutoPercentage > 0 {

		tmpUserSegments := make([]models.UserSegments, 0)
		config.DB.Model(&models.UserSegments{}).Find(&tmpUserSegments)

		getUserSequence(c, segment.AutoPercentage, segment.Slug)
		// c.JSON(http.StatusOK, gin.H{"UsersAmount:": getUserSequence(segment.AutoPercentage)})
	}
	// c.JSON(http.StatusOK, &segment)
}

// UpdateSegment godoc
// @Summary Update Segment
// @Description Updates segment slug with new slug
// @Tags segments
// @Produce json
// @Param slug formData string true "slug of the segment you want to update"
// @Param newSlug formData string true "new slug for the segment"
// @Success 200
// @Router /segment [put]
func UpdateSegment(c *gin.Context) {
	var segmentChange SegmentChange
	var segment models.Segment
	c.BindJSON(&segmentChange)
	if r := config.DB.Model(&models.Segment{}).Where("slug = ?", segmentChange.Slug).Limit(1).Find(&segment); r.RowsAffected == 0 {
		fmt.Printf("segment with slug \"%s\" mot found\n", segment.Slug)
	} else {
		segment.Slug = segmentChange.NewSlug
		config.DB.Save(&segment)
		c.JSON(http.StatusOK, &segment)
	}
}

// DeleteSegment godoc
// @Summary Delete Segment
// @Description Deletes segment passing it's slug
// @Tags segments
// @Produce json
// @Param slug formData string true "slug of the segment you want to delete"
// @Success 200
// @Router /segment [delete]
func DeleteSegment(c *gin.Context) {
	var segment models.Segment
	c.BindJSON(&segment)
	if err := config.DB.Where("slug = ?", segment.Slug).Delete(&segment).Error; err != nil {
		fmt.Printf("segment with slug \"%s\" mot found\n", segment.Slug)
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("segment with slug = %s", segment.Slug)})
}
