package controller

import (
	"net/http"
	"rianRestApi/config"
	"rianRestApi/internal/models"
	"rianRestApi/internal/repository"
	"rianRestApi/utils"

	"github.com/gin-gonic/gin"
)

type PurchasingController struct {
	rp *repository.PurchasingRepository
}

func NewPurchasingController() *PurchasingController {
	db := config.ConnectDb()
	repo := repository.NewPurchasingRepository(db)
	return &PurchasingController{rp: repo}
}

// funct get Alldata()
func (repo *PurchasingController) Index(c *gin.Context) {
	data, err := repo.rp.IndexAll()
	if err != nil {
		c.JSON(200, gin.H{
			"data": data,
		})
	}
	c.JSON(200, utils.SuccessResponse(data, 200, "success"))
}

//func save

func (repo *PurchasingController) SaveData(c *gin.Context) {
	var b models.PurcashingModel
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}
	data, err := repo.rp.Insert(b)

	if err != nil {
		c.JSON(200, gin.H{
			"error": data,
		})
	}
	c.JSON(200, gin.H{
		"payload": b,
		"data":    "success insert data",
	})
}

func (repo *PurchasingController) Upload(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{
			"payload": file,
			"data":    "empty file",
		})
	}
	defer file.Close()

	filename := header.Filename
	path := "upload/" + filename
	if err := c.SaveUploadedFile(header, path); err != nil {
		c.JSON(500, gin.H{
			"data":    "failed to save file",
			"payload": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": "success upload file",
		"payload": gin.H{
			"filename": filename,
			"path":     path,
			"size":     header.Size,
		},
	})

}
