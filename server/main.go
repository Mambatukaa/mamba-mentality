package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"mamba-mentality.com/controllers"
)

func main() {
	fmt.Print(("hello world"));

	router := gin.Default();

	router.GET("/albums", controllers.GetAlbums)
	router.GET("/album", controllers.GetAlbum)
	router.POST("/album", controllers.PostAlbum)
	router.DELETE("/album", controllers.DeleteAlbum)


	router.Run("localhost:8000");
};


// func getCurrencyData(currency string) {
// 	rate, err := api.GetRate(currency);

// 	if(err == nil) {
// 		fmt.Printf("The rate for %v is %.2f\n", rate.Currency, rate.Price);
// 	};
// }