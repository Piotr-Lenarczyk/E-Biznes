package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const productIDPath = "/products/:id"

type Product struct {
	ID         uint       `json:"id"`
	Name       string     `json:"name"`
	Price      float64    `json:"price"`
	Categories []Category `gorm:"many2many:product_categories;" json:"categories,omitempty"`
}

type ProductController struct {
	DB *gorm.DB
}

func RegisterProductRoutes(e *echo.Echo, db *gorm.DB) {
	pc := &ProductController{DB: db}

	e.GET("/products", pc.GetProducts)
	e.POST("/products", pc.CreateProduct)
	e.GET(productIDPath, pc.GetProductByID)
	e.PUT(productIDPath, pc.UpdateProduct)
	e.DELETE(productIDPath, pc.DeleteProduct)
}

func (pc *ProductController) GetProducts(ctx echo.Context) error {
	var products []Product
	if err := pc.DB.Preload("Categories").Find(&products).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, products)
}

func (pc *ProductController) GetProductByID(ctx echo.Context) error {
	var product Product
	id := ctx.Param("id")

	if err := pc.DB.Preload("Categories").First(&product, id).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"message": "Product not found"})
	}
	return ctx.JSON(http.StatusOK, product)
}

func (pc *ProductController) CreateProduct(ctx echo.Context) error {
	var product Product
	if err := ctx.Bind(&product); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	// Validate categories
	if err := pc.validateCategories(&product); err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"message": err.Error()})
	}

	if err := pc.DB.Create(&product).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	// Reload with categories
	if err := pc.DB.Preload("Categories").First(&product, product.ID).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, product)
}

func (pc *ProductController) UpdateProduct(ctx echo.Context) error {
	var product Product
	id := ctx.Param("id")

	if err := pc.DB.First(&product, id).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"message": "Product not found"})
	}

	if err := ctx.Bind(&product); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	// Validate categories
	if err := pc.validateCategories(&product); err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"message": err.Error()})
	}

	if err := pc.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&product).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	// Reload with categories
	if err := pc.DB.Preload("Categories").First(&product, product.ID).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, product)
}

func (pc *ProductController) DeleteProduct(ctx echo.Context) error {
	id := ctx.Param("id")

	if err := pc.DB.Delete(&Product{}, id).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.NoContent(http.StatusNoContent)
}

// validateCategories ensures all passed categories exist in DB
func (pc *ProductController) validateCategories(product *Product) error {
	if len(product.Categories) == 0 {
		return nil
	}

	var categoryIDs []uint
	for _, c := range product.Categories {
		categoryIDs = append(categoryIDs, c.ID)
	}

	var existingCategories []Category
	if err := pc.DB.Where("id IN ?", categoryIDs).Find(&existingCategories).Error; err != nil {
		return err
	}

	if len(existingCategories) != len(categoryIDs) {
		return echo.NewHTTPError(http.StatusNotFound, "One or more categories not found")
	}

	// Replace with fully loaded categories
	product.Categories = existingCategories

	return nil
}
