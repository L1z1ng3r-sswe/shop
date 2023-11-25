package handler

import (
	"net/http"
	"shop/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (h *Handler) signUp(c *gin.Context) {
	var newUser models.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		responseErrorMessage(c, http.StatusBadRequest, "Invalid credentials")
		return
	}

	validate := validator.New()

	if err := validate.Struct(newUser); err != nil {
		responseErrorMessage(c, http.StatusBadRequest, "Validation is not allowed")
		return
	}

	id, err := h.service.Authorization.CreateUser(newUser)
	if err != nil {
		responseErrorMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, id)
}
