package controller

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/Damir564/HttpApiGo/config"
	"github.com/Damir564/HttpApiGo/models"
	"github.com/gin-gonic/gin"
)

// type History struct {
// 	UserID    uint
// 	SegmentID uint
// 	Segment   string
// 	Operation Operation
// 	Timestamp time.Time
// }

// type Data struct {
// 	UserID      uint           `json:"user_id" gorm:"column:user_id"`
// 	SegmentID   uint           `json:"segment_id" gorm:"column:segment_id"`
// 	SegmentSlug string         `json:"segmentSlug" gorm:"column:slug"`
// 	CreatedAt   time.Time      `json:"created_at" gorm:"column:created_at"`
// 	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"column:deleted_at"`
// }

type YearMonth struct {
	Year  uint `json:"year"`
	Month uint `json:"month"`
	// Day   uint `json:"day"`
}

// GetHistory godoc
// @Summary Get History
// @Description Gets records in user_segments for specific month-year. Creates .csv file in the project
// @Tags user_segments
// @Produce json
// @Param year formData int true "year of records"
// @Param month formData int true "month of records"
// @Success 200
// @Router /history [get]
func GetHistory(c *gin.Context) {
	var yearMonth YearMonth
	c.BindJSON(&yearMonth)
	if yearMonth.Year == 0 {
		c.JSON(http.StatusBadRequest, "key \"year\" not correct")
		return
	}
	if yearMonth.Month == 0 || yearMonth.Month > 12 {
		c.JSON(http.StatusBadRequest, "key \"month\" not correct")
		return
	}
	histories := make([]models.History, 0)

	tmpUserSegments := make([]models.UserSegments, 0)
	config.DB.Model(&models.UserSegments{}).Find(&tmpUserSegments)

	// config.DB.Model(&models.History{}).Find(&histories)
	// config.DB.Model(&models.UserSegments{}).Find(&userSegments)

	if r := config.DB.Model(&models.History{}).Where("(EXTRACT('Year' FROM timestamp) = ? AND EXTRACT('Month' FROM timestamp) = ?)",
		yearMonth.Year,
		yearMonth.Month).
		Find(&histories); r.RowsAffected == 0 {
		c.String(http.StatusOK, "no data in this period")
	} else {
		sort.Slice(histories, func(i, j int) bool {
			return histories[i].Timestamp.Before(histories[i].Timestamp)
		})
		fileName := fmt.Sprintf("%d-%d.csv", yearMonth.Year, yearMonth.Month)
		filePath := fmt.Sprintf("./%s", fileName)

		file, err := os.Create(filePath)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()
		headers := []string{"N", "UserID", "SegmentId", "SegmentSlug", "Operation", "Timestamp"}
		writer.Write(headers)
		for _, history := range histories {
			row := []string{
				strconv.Itoa(int(history.ID)),
				strconv.FormatUint(uint64(history.UserID), 10),
				strconv.FormatUint(uint64(history.SegmentID), 10),
				history.SegmentSlug,
				history.Operation,
				history.Timestamp.Format(time.RFC3339),
			}
			if err := writer.Write(row); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"Error:": err.Error()})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{"report_url": filePath})
	}
}
