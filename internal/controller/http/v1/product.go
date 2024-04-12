package v1

import (
	ulid2 "github.com/oklog/ulid/v2"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"go-clean-template/internal/entity"
	"go-clean-template/internal/usecase"
	"go-clean-template/pkg/logger"
)

type productRoutes struct {
	t usecase.Product
	l logger.Interface
}

func newProductRoutes(handler *gin.RouterGroup, t usecase.Product, l logger.Interface) {
	r := &productRoutes{t, l}

	h := handler.Group("/product")
	{
		h.POST("/add-product", r.addProduct)
		h.GET("/get-product", r.getProduct)
	}
}

// @Summary     Product
// @Description Add a product to inventory
// @ID          add-product
// @Tags  	    product
// @Accept      json
// @Produce     json
// @Param       request body addProductRequest true "Enter product details"
// @Success     200 {object} entity.Product
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /product/add-product [post]
func (r *productRoutes) addProduct(c *gin.Context) {
	var request addProductRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - addProduct")
		errorResponse(c, http.StatusBadRequest, "invalid request body")

		return
	}
	t := time.Now()
	entropy := ulid2.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	ulid := ulid2.MustNew(ulid2.Timestamp(t), entropy)

	ulidString := ulid.String()

	product, err := r.t.AddProduct(
		c.Request.Context(),
		entity.Product{
			ID:          ulidString,
			Name:        request.Name,
			Description: request.Description,
			CategoryID:  request.CategoryID,
			UnitPrice:   request.UnitPrice,
		},
	)
	if err != nil {
		r.l.Error(err, "http - v1 - addProduct")
		errorResponse(c, http.StatusInternalServerError, "add-product service problems")

		return
	}

	c.JSON(http.StatusOK, product)
}

// @Summary     Product
// @Description Get one product details from inventory
// @ID          get-product
// @Tags  	    product
// @Accept      json
// @Produce     json
// @Param       id query string true "Product ID"
// @Success     200 {object} entity.Product
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /product/get-product [get]
func (r *productRoutes) getProduct(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		errorResponse(c, http.StatusBadRequest, "id parameter is required")
	}

	product, err := r.t.GetProduct(
		c.Request.Context(), id,
	)
	if err != nil {
		r.l.Error(err, "http - v1 - getProduct")
		errorResponse(c, http.StatusInternalServerError, "get-product service problems")

		return
	}

	c.JSON(http.StatusOK, product)
}
