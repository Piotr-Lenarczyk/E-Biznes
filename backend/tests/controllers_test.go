package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"go-echo-gorm-app/controllers"
)

func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&controllers.Product{}, &controllers.Category{}, &controllers.Cart{}, &controllers.Payment{})
	return db
}

func TestProductController(t *testing.T) {
	e := echo.New()
	db := setupTestDB()
	pc := &controllers.ProductController{DB: db}

	db.Create(&controllers.Product{Name: "Test Product", Price: 10.0})

	// Test GetProducts
	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if assert.NoError(t, pc.GetProducts(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Test Product")
	}

	// Test CreateProduct
	productJSON := `{"name":"New Product","price":20.0}`
	req = httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(productJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	ctx = e.NewContext(req, rec)

	if assert.NoError(t, pc.CreateProduct(ctx)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Contains(t, rec.Body.String(), "New Product")
	}
}

func TestCartController(t *testing.T) {
	e := echo.New()
	db := setupTestDB()
	cc := &controllers.CartController{DB: db}

	// Dodanie przykładowych danych
	product := controllers.Product{Name: "Test Product", Price: 10.0}
	db.Create(&product)

	// Test CreateCart
	cartJSON := `{"products":[{"id":1}]}`
	req := httptest.NewRequest(http.MethodPost, "/carts", strings.NewReader(cartJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if assert.NoError(t, cc.CreateCart(ctx)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Contains(t, rec.Body.String(), "Test Product")
	}

	// Test GetCarts
	req = httptest.NewRequest(http.MethodGet, "/carts", nil)
	rec = httptest.NewRecorder()
	ctx = e.NewContext(req, rec)

	if assert.NoError(t, cc.GetCarts(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Test Product")
	}
}

func TestCategoryController(t *testing.T) {
	e := echo.New()
	db := setupTestDB()
	cc := &controllers.CategoryController{DB: db}

	// Test CreateCategory
	categoryJSON := `{"name":"Test Category"}`
	req := httptest.NewRequest(http.MethodPost, "/categories", strings.NewReader(categoryJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if assert.NoError(t, cc.CreateCategory(ctx)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Contains(t, rec.Body.String(), "Test Category")
	}

	// Test GetCategories
	req = httptest.NewRequest(http.MethodGet, "/categories", nil)
	rec = httptest.NewRecorder()
	ctx = e.NewContext(req, rec)

	if assert.NoError(t, cc.GetCategories(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Test Category")
	}
}

func TestPaymentController(t *testing.T) {
	e := echo.New()
	db := setupTestDB()
	pc := &controllers.PaymentController{DB: db}

	// Test CreatePayment
	paymentJSON := `{"amount":100,"cardNumber":"1234567812345678","expirationDate":"12/25"}`
	req := httptest.NewRequest(http.MethodPost, "/payments", strings.NewReader(paymentJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if assert.NoError(t, pc.CreatePayment(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Payment Successful")
	}
}