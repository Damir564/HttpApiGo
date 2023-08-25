package controller

import (
	"net/http"

	"github.com/Damir564/HttpApiGo/config"
	"github.com/Damir564/HttpApiGo/models"
	"github.com/gin-gonic/gin"
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
	config.DB.Where("id = ?", c.Param("id")).Delete(&segment)
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

func Bind(c *gin.Context) {
	var userSegments models.UserSegments
	c.BindJSON(&userSegments)
	config.DB.Save(&userSegments)
	c.JSON(http.StatusOK, &userSegments)
}
