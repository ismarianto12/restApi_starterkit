package controller

import (
	"log"
	"net/http"
	"rianRestApi/config"
	"rianRestApi/internal/models"
	"rianRestApi/internal/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BarangController struct {
	repo *repository.BarangRepository
}

func NewBarangController() *BarangController {
	db := config.ConnectDb()
	repo := repository.NewBarangRepository(db)
	return &BarangController{repo: repo}
}

func (br *BarangController) GetAll(c *gin.Context) {
	data, err := br.repo.GetAlldata()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})

}

func (br *BarangController) Save(c *gin.Context) {
	var b models.BarangModel

	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}

	data, err := br.repo.Save(&b)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (bc *BarangController) FilteringData(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	Searching := c.DefaultQuery("Searching", "")

	barang, totalrow, err := bc.repo.Filtering(page, pageSize, Searching)
	if err != nil {
		log.Fatal("Erro coult not start server properly")

	}

	c.JSON(http.StatusOK, gin.H{
		"data":    barang,
		"totarow": totalrow,
		"message": "response",
	})

}
