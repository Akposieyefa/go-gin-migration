package handlers

import (
	"net/http"

	"github.com/akposiyefa/go-gin-migration/core/models"
	"github.com/akposiyefa/go-gin-migration/internal"
	"github.com/akposiyefa/go-gin-migration/utils"
	"github.com/gin-gonic/gin"
)

func ChangeUserPassword(c *gin.Context) {

	payload := models.UserPasswordChangePayload{}
	err := c.BindJSON(&payload)

	if err != nil {
		utils.WriteError(c, http.StatusUnprocessableEntity, "Sorry validation failed", false)
		return
	}

	if payload.Password != payload.PasswordConfirmation {
		utils.WriteError(c, http.StatusOK, "Sorry password do nor match confirm password", false)
		return
	}

	password, err := internal.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(c, http.StatusBadRequest, "Sorry failed to hash password", false)
		return
	}

	user := models.User{
		Password: string(password),
	}
	result := internal.DB.Create(&user)
	if result.Error != nil {
		utils.WriteError(c, http.StatusBadRequest, "Sorry unable to change user password", false)
		return
	}
	utils.WriteSuccess(c, http.StatusOK, "User changed successfully", map[string]string{}, true)
}
