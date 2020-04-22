package handlers

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/spoik/animal-crossing/models"
)

type ErrorResponse struct {
	ErrorMessages []string `json:"error_messages"`
}

func validateItem(model models.Item, c *gin.Context) *ErrorResponse {
	validate := c.MustGet("validator").(*validator.Validate)
	translator := c.MustGet("translator").(*ut.Translator)
	validateErrors := validate.Struct(&model)

	if validateErrors == nil {
		return nil
	}

	validationErrors := validateErrors.(validator.ValidationErrors)

	var errorMessages []string

	for _, errMessage := range validationErrors.Translate(*translator) {
		errorMessages = append(errorMessages, errMessage)
	}

	return &ErrorResponse{ErrorMessages: errorMessages}
}
