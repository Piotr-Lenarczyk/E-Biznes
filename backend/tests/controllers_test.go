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

	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	assert.NoError(t, pc.GetProducts(ctx))
	assert.Equal(t, http.StatusOK, rec.Code)

	req = httptest.NewRequest(http.MethodGet, "/products/999", nil)
	rec = httptest.NewRecorder()
	ctx = e.NewContext(req, rec)

	assert.NoError(t, pc.GetProductByID(ctx))
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestCartController(t *testing.T) {
    e := echo.New()
    db := setupTestDB()
    cc := &controllers.CartController{DB: db}

    cartJSON := `{"products":[]}`
    req := httptest.NewRequest(http.MethodPost, "/carts", strings.NewReader(cartJSON))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()
    ctx := e.NewContext(req, rec)

    assert.NoError(t, cc.CreateCart(ctx))
    assert.Equal(t, http.StatusCreated, rec.Code)

    req = httptest.NewRequest(http.MethodGet, "/carts/999", nil)
    rec = httptest.NewRecorder()
    ctx = e.NewContext(req, rec)

    ctx.SetParamNames("id")
    ctx.SetParamValues("999")

    assert.NoError(t, cc.GetCartByID(ctx))
    assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestCategoryController(t *testing.T) {
    e := echo.New()
    db := setupTestDB()
    cc := &controllers.CategoryController{DB: db}

    categoryJSON := `{"name":"Test Category"}`
    req := httptest.NewRequest(http.MethodPost, "/categories", strings.NewReader(categoryJSON))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()
    ctx := e.NewContext(req, rec)

    assert.NoError(t, cc.CreateCategory(ctx))
    assert.Equal(t, http.StatusCreated, rec.Code)

    categoryJSON = `{"name":"Updated Category"}`
    req = httptest.NewRequest(http.MethodPut, "/categories/999", strings.NewReader(categoryJSON))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec = httptest.NewRecorder()
    ctx = e.NewContext(req, rec)

    ctx.SetParamNames("id")
    ctx.SetParamValues("999")

    assert.NoError(t, cc.UpdateCategory(ctx))
    assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestPaymentController(t *testing.T) {
	e := echo.New()
	db := setupTestDB()
	pc := &controllers.PaymentController{DB: db}

	paymentJSON := `{"amount":100,"cardNumber":"1234567812345678","expirationDate":"12/25"}`
	req := httptest.NewRequest(http.MethodPost, "/payments", strings.NewReader(paymentJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	assert.NoError(t, pc.CreatePayment(ctx))
	assert.Equal(t, http.StatusOK, rec.Code)

	paymentJSON = `{"amount":0,"cardNumber":"","expirationDate":""}`
	req = httptest.NewRequest(http.MethodPost, "/payments", strings.NewReader(paymentJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	ctx = e.NewContext(req, rec)

	assert.NoError(t, pc.CreatePayment(ctx))
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}