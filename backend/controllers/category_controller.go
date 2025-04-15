package controllers

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "gorm.io/gorm"
)

const categoryIDPath = "/categories/:id"

// Category model
type Category struct {
    ID   uint   `json:"id"`
    Name string `json:"name"`
}

type CategoryController struct {
    DB *gorm.DB
}

func RegisterCategoryRoutes(e *echo.Echo, db *gorm.DB) {
    categoryController := &CategoryController{DB: db}

    e.GET("/categories", categoryController.GetCategories)
    e.POST("/categories", categoryController.CreateCategory)
    e.GET(categoryIDPath, categoryController.GetCategoryByID)
    e.PUT(categoryIDPath, categoryController.UpdateCategory)
    e.DELETE(categoryIDPath, categoryController.DeleteCategory)
}

func (cc *CategoryController) GetCategories(ctx echo.Context) error {
    var categories []Category
    if err := cc.DB.Find(&categories).Error; err != nil {
        return ctx.JSON(http.StatusInternalServerError, err.Error())
    }
    return ctx.JSON(http.StatusOK, categories)
}

func (cc *CategoryController) GetCategoryByID(ctx echo.Context) error {
    var category Category
    id := ctx.Param("id")
    if err := cc.DB.First(&category, id).Error; err != nil {
        return ctx.JSON(http.StatusNotFound, echo.Map{"message": "Category not found"})
    }
    return ctx.JSON(http.StatusOK, category)
}

func (cc *CategoryController) CreateCategory(ctx echo.Context) error {
    var category Category
    if err := ctx.Bind(&category); err != nil {
        return err
    }

    if err := cc.DB.Create(&category).Error; err != nil {
        return ctx.JSON(http.StatusInternalServerError, err.Error())
    }
    return ctx.JSON(http.StatusCreated, category)
}

func (cc *CategoryController) UpdateCategory(ctx echo.Context) error {
    var category Category
    id := ctx.Param("id")
    if err := cc.DB.First(&category, id).Error; err != nil {
        return ctx.JSON(http.StatusNotFound, echo.Map{"message": "Category not found"})
    }

    if err := ctx.Bind(&category); err != nil {
        return err
    }

    if err := cc.DB.Save(&category).Error; err != nil {
        return ctx.JSON(http.StatusInternalServerError, err.Error())
    }

    return ctx.JSON(http.StatusOK, category)
}

func (cc *CategoryController) DeleteCategory(ctx echo.Context) error {
    id := ctx.Param("id")
    if err := cc.DB.Delete(&Category{}, id).Error; err != nil {
        return ctx.JSON(http.StatusInternalServerError, err.Error())
    }
    return ctx.NoContent(http.StatusNoContent)
}
