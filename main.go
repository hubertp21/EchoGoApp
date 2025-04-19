package main

import (
	"productsapp/database"
	"productsapp/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db := database.InitDB()
	e.Use(database.DBMiddleware(db))

	e.GET("/products", handlers.GetProducts)
	e.GET("/products/:id", handlers.GetProduct)
	e.POST("/products", handlers.CreateProduct)
	e.PUT("/products/:id", handlers.UpdateProduct)
	e.DELETE("/products/:id", handlers.DeleteProduct)

	e.Logger.Fatal(e.Start(":8080"))
}
