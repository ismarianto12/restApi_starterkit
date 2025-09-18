package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type HomeController struct {
	DB *gorm.DB
}

func NewHomeController(db *gorm.DB) *HomeController {
	return &HomeController{DB: db}
}

func (HomeController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello from BarangController"})

}
