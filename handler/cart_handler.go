package handler

import (
	"github.com/MochamadAkbar/ordent-test/config"
	"github.com/MochamadAkbar/ordent-test/middleware"
	"net/http"

	"github.com/MochamadAkbar/ordent-test/common/serializer"
)

type CartHandlerImpl struct {
}

func NewCartHandler() http.Handler {
	handler := &CartHandlerImpl{}

	router := config.NewRouter()
	router.Use(middleware.Auth)
	router.Post("/v1/cart", handler.AddToCart)
	router.Put("/v1/cart/{id}", handler.UpdateCart)
	router.Post("/v1/cart/checkout", handler.Checkout)

	return router
}

func (handler *CartHandlerImpl) AddToCart(w http.ResponseWriter, r *http.Request) {
	_ = serializer.SerializeWriter(w, http.StatusCreated, "implement me Add To Cart")
}

func (handler *CartHandlerImpl) UpdateCart(w http.ResponseWriter, r *http.Request) {
	_ = serializer.SerializeWriter(w, http.StatusOK, "implement me Update Cart")
}

func (handler *CartHandlerImpl) Checkout(w http.ResponseWriter, r *http.Request) {
	_ = serializer.SerializeWriter(w, http.StatusOK, "implement me Checkout")
}
