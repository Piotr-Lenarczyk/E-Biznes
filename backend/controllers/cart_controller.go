package controllers

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Cart struct {
	ID       uint      `json:"id"`
	Products []Product `gorm:"many2many:cart_products;" json:"products"`
}

type CartController struct {
	DB *gorm.DB
}

func RegisterCartRoutes(e *echo.Echo, db *gorm.DB) {
	cartController := &CartController{DB: db}

	e.GET("/carts", cartController.GetCarts)
	e.POST("/carts", cartController.CreateCart)
	e.GET("/carts/:id", cartController.GetCartByID)
	e.PUT("/carts/:id", cartController.UpdateCart)
	e.DELETE("/carts/:id", cartController.DeleteCart)
}

// ProductIDs returns a slice of Product IDs from the Cart struct.
func (c *Cart) ProductIDs() []uint {
	var productIDs []uint
	for _, product := range c.Products {
		productIDs = append(productIDs, product.ID)
	}
	return productIDs
}

func (c *CartController) GetCarts(ctx echo.Context) error {
	var carts []Cart
	if err := c.DB.Preload("Products").Find(&carts).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, carts)
}

func (c *CartController) CreateCart(ctx echo.Context) error {
	var cart Cart
	if err := ctx.Bind(&cart); err != nil {
		return err
	}

	// Get all products based on the provided product IDs
	var products []Product
	if err := c.DB.Where("id IN ?", cart.ProductIDs()).Find(&products).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	// Associate the products with the cart
	cart.Products = products
	if err := c.DB.Create(&cart).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, cart)
}

func (c *CartController) GetCartByID(ctx echo.Context) error {
	var cart Cart
	id := ctx.Param("id")
	if err := c.DB.Preload("Products").First(&cart, id).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"message": "Cart not found"})
	}
	return ctx.JSON(http.StatusOK, cart)
}

func (c *CartController) UpdateCart(ctx echo.Context) error {
	var cart Cart
	id := ctx.Param("id")
	if err := c.DB.Preload("Products").First(&cart, id).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"message": "Cart not found"})
	}

	// Bind the new data to the cart
	if err := ctx.Bind(&cart); err != nil {
		return err
	}

	// Get all products based on the provided product IDs
	var products []Product
	if err := c.DB.Where("id IN ?", cart.ProductIDs()).Find(&products).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	// Associate the products with the cart
	cart.Products = products

	if err := c.DB.Save(&cart).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, cart)
}

func (c *CartController) DeleteCart(ctx echo.Context) error {
	id := ctx.Param("id")
	if err := c.DB.Delete(&Cart{}, id).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.NoContent(http.StatusNoContent)
}
