package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/jinzhu/gorm"
	"github.com/go-playground/validator/v10"
	"github.com/spoik/animal-crossing/admin"
	"github.com/spoik/animal-crossing/handlers"
	"net/http"
)

func Create(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	mux := http.NewServeMux()
	Admin := admin.Setup(db)
	validate, translator := createValidator()

	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Set("validator", validate)
		c.Set("translator", translator)
		c.Next()
	})

	router.GET("/items", handlers.AllBooks)
	router.POST("/items", handlers.CreateBook)

	Admin.MountTo("/admin", mux)
	router.Any("/admin/*resources", gin.WrapH(mux))

	return router
}

func createValidator() (*validator.Validate, *ut.Translator) {
	validator := validator.New()
	en := en.New()
	uni := ut.New(en, en)
	trans, found := uni.GetTranslator("en")

	if !found {
		panic("Unable to find translator")
	}

	en_translations.RegisterDefaultTranslations(validator, trans)
	return validator, &trans
}

