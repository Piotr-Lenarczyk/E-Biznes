package controllers

import (
    "github.com/labstack/echo/v4"
    "gorm.io/gorm"
    "net/http"
)

type Product struct {
    ID    uint   `json:"id"`
    Name  string `json:"name"`
    Price int    `json:"price"`
}

var db *gorm.DB

func SetDB(database *gorm.DB) {
    db = database
}

func RegisterProductRoutes(e *echo.Echo) {
    e.GET("/products", GetProducts)
    e.POST("/products", CreateProduct)
    e.GET("/products/:id", GetProductByID)
    e.PUT("/products/:id", UpdateProduct)
    e.DELETE("/products/:id", DeleteProduct)
}

// Get all products
func GetProducts(c echo.Context) error {
    var products []Product
    if err := db.Find(&products).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, products)
}

// Create a new product
func CreateProduct(c echo.Context) error {
    var p Product
    if err := c.Bind(&p); err != nil {
        return err
    }
    if err := db.Create(&p).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusCreated, p)
}

// Get product by ID
func GetProductByID(c echo.Context) error {
    var p Product
    id := c.Param("id")
    if err := db.First(&p, id).Error; err != nil {
        return c.JSON(http.StatusNotFound, echo.Map{"message": "Product not found"})
    }
    return c.JSON(http.StatusOK, p)
}

// Update product by ID
func UpdateProduct(c echo.Context) error {
    var p Product
    id := c.Param("id")
    if err := db.First(&p, id).Error; err != nil {
        return c.JSON(http.StatusNotFound, echo.Map{"message": "Product not found"})
    }
    if err := c.Bind(&p); err != nil {
        return err
    }
    if err := db.Save(&p).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, p)
}

// Delete product by ID
func DeleteProduct(c echo.Context) error {
    var p Product
    id := c.Param("id")
    if err := db.Delete(&p, id).Error; err != nil {
        return c.JSON(http.StatusNotFound, echo.Map{"message": "Product not found"})
    }
    return c.NoContent(http.StatusNoContent)
}
