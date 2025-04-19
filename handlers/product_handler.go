package handlers

import (
	"net/http"
	"productsapp/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetProducts(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	var products []models.Product
	db.Find(&products)
	return c.JSON(http.StatusOK, products)
}

func GetProduct(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	var product models.Product
	if err := db.First(&product, c.Param("id")).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Product not found"})
	}
	return c.JSON(http.StatusOK, product)
}

func CreateProduct(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	var product models.Product
	if err := c.Bind(&product); err != nil {
		return err
	}
	db.Create(&product)
	return c.JSON(http.StatusCreated, product)
}

func UpdateProduct(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	var product models.Product
	id := c.Param("id")
	if err := db.First(&product, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Product not found"})
	}
	if err := c.Bind(&product); err != nil {
		return err
	}
	db.Save(&product)
	return c.JSON(http.StatusOK, product)
}

func DeleteProduct(c echo.Context) error {
	db := c.Get("db").(*gorm.DB)
	var product models.Product
	if err := db.First(&product, c.Param("id")).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Product not found"})
	}
	db.Delete(&product)
	return c.NoContent(http.StatusNoContent)
}
