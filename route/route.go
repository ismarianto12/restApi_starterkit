package route

import (
	"rianRestApi/internal/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetRoutingApps(r *gin.Engine, db *gorm.DB) {

	barangReal := controller.NewBarangController()
	cat := controller.NewCategoryController()
	prc := controller.NewPurchasingController()

	api := r.Group("/api")
	{
		api.GET("/barang/getALl", barangReal.GetAll)
		api.GET("/barang/filtering", barangReal.FilteringData)
		//
		api.POST("/barang/save", barangReal.Save)
		api.GET("/purchasing/all", prc.Index)
		api.GET("/purchasing/save", prc.SaveData)
		api.GET("/purchasing/upload", prc.Upload)

		api.GET("/cat/all", cat.All)

	}

	pantek := r.Group("/pantek")
	{
		pantek.GET("/barang/getALl", barangReal.GetAll)
		pantek.GET("/barang/filtering", barangReal.FilteringData)

		pantek.POST("/barang/save", barangReal.Save)
		pantek.GET("/cat/all", cat.All)

	}

}
