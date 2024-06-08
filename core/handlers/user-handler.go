package handlers

import (
	"net/http"

	"github.com/akposiyefa/go-gin-migration/core/models"
	"github.com/akposiyefa/go-gin-migration/core/response"
	"github.com/akposiyefa/go-gin-migration/internal"
	"github.com/akposiyefa/go-gin-migration/utils"
	"github.com/gin-gonic/gin"
)

// create user
func CreateUser(c *gin.Context) {

	payload := models.UserPayload{}
	err := c.BindJSON(&payload)

	if err != nil {
		utils.WriteError(c, http.StatusUnprocessableEntity, "Sorry validation fails", false)
		return
	}
	password, err := internal.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, "Sorry failed to hash password", false)
		return
	}

	user := models.User{
		Name:        payload.Name,
		Email:       payload.Email,
		PhoneNumber: payload.PhoneNumber,
		Password:    string(password),
	}
	result := internal.DB.Create(&user)
	if result.Error != nil {
		utils.WriteError(c, http.StatusBadRequest, "Sorry unable to create user account", false)
		return
	}
	resp := response.UserResponse{
		ID:          user.ID.String(),
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		DeletedAt:   user.DeletedAt.Time,
	}
	utils.WriteSuccess(c, http.StatusCreated, "User created successfully", resp, true)
}

// get users
func GetUsers(c *gin.Context) {
	var users []models.User
	result := internal.DB.Find(&users)
	if result.Error != nil {
		utils.WriteError(c, http.StatusBadRequest, "Sorry unable to fetch users", false)
		return
	}
	utils.WriteSuccess(c, http.StatusOK, "Users fetched successfully", users, true)
}

// get user
func GetUser(c *gin.Context) {

	id := c.Param("id")
	var user models.User

	result := internal.DB.First(&user, "id = ?", id)
	if result.Error != nil {
		utils.WriteError(c, http.StatusBadRequest, "Sorry unable to get single users", false)
		return
	}
	resp := response.UserResponse{
		ID:          user.ID.String(),
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		DeletedAt:   user.DeletedAt.Time,
	}
	utils.WriteSuccess(c, http.StatusOK, "Users fetched successfully", resp, true)
}

// delete user
func UpdateUser(c *gin.Context) {

	id := c.Param("id")
	payload := models.UserUpdatePayload{}
	err := c.BindJSON(&payload)

	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, "Sorry validation fails", false)
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
		utils.WriteError(c, http.StatusBadRequest, "Sorry unable to update profile", false)
		return
	}
	utils.WriteSuccess(c, http.StatusOK, "User updated successfully", map[string]string{}, true)
}

// delete users
func DeleteUser(c *gin.Context) {

	id := c.Param("id")
	var user models.User

	internal.DB.First(&user, "id = ?", id)
	result := internal.DB.Delete(&user, "id = ?", id)

	if result.Error != nil {
		utils.WriteError(c, http.StatusBadRequest, "Sorry unable to delete user profile", false)
		return
	}
	utils.WriteSuccess(c, http.StatusOK, "User deleted successfully", map[string]string{}, true)
}
