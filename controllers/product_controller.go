package controllers

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
)

type Product struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var products []Product
var nextID uint = 1

func RegisterProductRoutes(e *echo.Echo) {
	e.GET("/products", GetProducts)
	e.POST("/products", CreateProduct)
	e.GET("/products/:id", GetProductByID)
	e.PUT("/products/:id", UpdateProduct)
	e.DELETE("/products/:id", DeleteProduct)
}

func GetProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

func CreateProduct(c echo.Context) error {
	var p Product
	if err := c.Bind(&p); err != nil {
		return err
	}
	p.ID = nextID
	nextID++
	products = append(products, p)
	return c.JSON(http.StatusCreated, p)
}

func GetProductByID(c echo.Context) error {
    // Convert the URL parameter ID to uint
    id := c.Param("id")

    // Iterate through products and compare the ID
    for _, p := range products {
        if fmt.Sprintf("%d", p.ID) == id { // Convert product ID to string and compare
            return c.JSON(http.StatusOK, p)
        }
    }
    return c.JSON(http.StatusNotFound, echo.Map{"message": "Product not found"})
}


func UpdateProduct(c echo.Context) error {
    id := c.Param("id") // Get the ID from the URL parameter
    for i, p := range products {
        if fmt.Sprintf("%d", p.ID) == id { // Convert product ID to string and compare
            if err := c.Bind(&products[i]); err != nil {
                return err
            }
            return c.JSON(http.StatusOK, products[i])
        }
    }
    return c.JSON(http.StatusNotFound, echo.Map{"message": "Product not found"})
}


func DeleteProduct(c echo.Context) error {
    id := c.Param("id") // Get the ID from the URL parameter
    for i, p := range products {
        if fmt.Sprintf("%d", p.ID) == id { // Convert product ID to string and compare
            products = append(products[:i], products[i+1:]...)
            return c.NoContent(http.StatusNoContent)
        }
    }
    return c.JSON(http.StatusNotFound, echo.Map{"message": "Product not found"})
}

