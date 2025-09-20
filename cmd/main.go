package main

import (
	"log"
	"rianRestApi/config"
	"rianRestApi/internal/models"
	"rianRestApi/route"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDb()
	db.AutoMigrate(&models.BarangModel{})
	db.AutoMigrate(&models.CategoryBarang{})
	db.AutoMigrate(&models.PurcashingModel{})
	db.AutoMigrate(&models.UserModel{})

	r := gin.Default()
	r.Static("/components", "./templates/components")

	route.SetRoutingApps(r, db)
	if err := r.Run(":6969"); err != nil {
		log.Fatal("Erro coult not start server properly")
	}
}

// func main() {
// 	var x int = 10

// 	fmt.Println(x, "pointer ex")
// 	for i := 0; i < (x); i++ {
// 		fmt.Println(i)
// 		for j := 0; j < x+i; j++ {
// 			fmt.Println(j)
// 		}
// 	}

// 	var p *int = &x
// 	fmt.Println("Pointer p menyimpan alamat:", p)
// 	fmt.Println("Nilai dari alamat yang ditunjuk p:", *p)
// 	db := config.ConnectDb()
// 	println(db)

// }
