package handlers

import (
	"github.com/akposiyefa/go-gin-migration/core/models"
	"github.com/akposiyefa/go-gin-migration/internal"
	"github.com/akposiyefa/go-gin-migration/utils"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	payload := models.UserPayload{}
	err := c.BindJSON(&payload)
	if err != nil {
		utils.WriteError(c, 400, "Sorry validation fails", false)
		return
	}
	password, _ := internal.HashPassword(payload.Password)
	user := models.User{
		Name:        payload.Name,
		Email:       payload.Email,
		PhoneNumber: payload.PhoneNumber,
		Password:    string(password),
	}
	result := internal.DB.Create(&user)
	if result.Error != nil {
		utils.WriteError(c, 400, "Sorry unable to create user account", false)
		return
	}
	utils.WriteSuccess(c, 201, "User created successfully", map[string]string{}, true)
}

func GetUsers(c *gin.Context) {
	var users []models.User
	result := internal.DB.Find(&users)
	if result.Error != nil {
		utils.WriteError(c, 400, "Sorry unable to fetch users", false)
		return
	}
	utils.WriteSuccess(c, 200, "Users fetched successfully", users, true)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	result := internal.DB.First(&user, "id = ?", id)
	if result.Error != nil {
		utils.WriteError(c, 400, "Sorry unable to get single users", false)
		return
	}
	utils.WriteSuccess(c, 200, "Users fetched successfully", user, true)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	payload := models.UserUpdatePayload{}
	err := c.BindJSON(&payload)
	if err != nil {
		utils.WriteError(c, 400, "Sorry validation fails", false)
		return
	}
	var user models.User
	internal.DB.First(&user, "id = ?", id)
	result := internal.DB.Model(&user).Updates(models.User{
		Name:        payload.Name,
		Email:       payload.Email,
		PhoneNumber: payload.PhoneNumber,
	})
	if result.Error != nil {
		utils.WriteError(c, 400, "Sorry unable to update profile", false)
		return
	}
	utils.WriteSuccess(c, 200, "User updated successfully", map[string]string{}, true)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	internal.DB.First(&user, "id = ?", id)
	result := internal.DB.Delete(&user, "id = ?", id)
	if result.Error != nil {
		utils.WriteError(c, 400, "Sorry unable to delete user profile", false)
		return
	}
	utils.WriteSuccess(c, 201, "User deleted successfully", map[string]string{}, true)
}
