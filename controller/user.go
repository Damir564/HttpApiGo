package controller

import (
	"net/http"

	"github.com/Damir564/HttpApiGo/config"
	"github.com/Damir564/HttpApiGo/models"
	"github.com/gin-gonic/gin"
)

// GetUsers godoc
// @Summary Gets Users
// @Description Gets all users
// @Tags users
// @Produce json
// @Success 200
// @Router /users [get]
func GetUsers(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(http.StatusOK, &users)
}

// CreateUser godoc
// @Summary Creates User
// @Description Creates user with autoincrement primary key
// @Tags users
// @Produce json
// @Success 200
// @Router /user [post]
func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	// if err := c.BindJSON(&segment); err != nil {
	// 	panic(err)
	// }
	config.DB.Create(&user)
	c.JSON(http.StatusOK, &user)
}

// DeleteUser godoc
// @Summary Deletes User
// @Description Deletes user passing it's ID
// @Tags users
// @Produce json
// @Param id formData int true "user's id you want to delete"
// @Success 200
// @Router /user/:id [delete]
func DeleteUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).Delete(&user)
	c.JSON(http.StatusOK, &user)
}

// // UpdateUser godoc
// // @Summary Updates User
// // @Description Updates user passing it's ID
// // @Tags users
// // @Produce json
// // @Param id formData int true "user's id you want to update"
// // @Success 200
// // @Router /user [delete]
// func UpdateUser(c *gin.Context) {
// 	var user models.User
// 	config.DB.Where("id = ?", c.Param("id")).First(&user)
// 	c.BindJSON(&user)
// 	// if err := c.BindJSON(&segment); err != nil {
// 	// 	panic(err)
// 	// }
// 	config.DB.Save(&user)
// 	c.JSON(http.StatusOK, &user)
// }
