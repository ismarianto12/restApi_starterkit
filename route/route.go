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
	userctrl := controller.NewUserController()

	api := r.Group("/api")
	{
		api.GET("/v1", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"api": "v1", "description": "Response api",
			})
		})
		api.GET("/barang/getALl", barangReal.GetAll)
		api.GET("/barang/filtering", barangReal.FilteringData)
		//
		api.POST("/barang/save", barangReal.Save)
		api.GET("/purchasing/all", prc.Index)
		api.GET("/purchasing/save", prc.SaveData)
		api.GET("/purchasing/upload", prc.Upload)
		api.GET("/purchasing/:id", prc.Show)
		api.GET("/purchasing/delete/:id", prc.Delete)

		api.GET("/cat/all", cat.All)

	}

	pantek := r.Group("/pantek")
	{
		pantek.GET("/barang/getALl", barangReal.GetAll)
		pantek.GET("/barang/filtering", barangReal.FilteringData)

		pantek.POST("/barang/save", barangReal.Save)
		pantek.GET("/cat/all", cat.All)

	}
	user := r.Group("/user")
	{
		user.GET("/index", userctrl.GetAll)
		user.POST("/save", userctrl.Insert)
		// user.GET("/barang/filtering", barangReal.FilteringData)

		// user.POST("/barang/save", barangReal.Save)
		// user.GET("/cat/all", cat.All)

	}

}
