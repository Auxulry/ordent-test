// Package handler describe http handler
package handler

import (
	"net/http"

	"github.com/MochamadAkbar/ordent-test/api"
	commonErr "github.com/MochamadAkbar/ordent-test/common/errors"
	"github.com/MochamadAkbar/ordent-test/common/serializer"
	"github.com/MochamadAkbar/ordent-test/entity"
	"github.com/MochamadAkbar/ordent-test/usecase"
	"github.com/go-chi/chi/v5"
)

type UserHandlerImpl struct {
	Usecase usecase.UserUsecase
}

func _(usecase usecase.UserUsecase, router *chi.Mux) {
	handler := &UserHandlerImpl{Usecase: usecase}

	router.Post("v1/authentication/register", handler.Register)
	router.Post("v1/authentication/login", handler.Login)
}

func (handler *UserHandlerImpl) Register(w http.ResponseWriter, r *http.Request) {
	var req api.UserRequestRegister

	err := serializer.SerializeRequest[*api.UserRequestRegister](r, &req)
	if err != nil {
		panic(err)
	}

	user := &entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	resp, err := handler.Usecase.Register(r.Context(), user)
	if err != nil {
		panic(err)
	}

	res := api.Response[api.UserResponse]{
		Code:   http.StatusCreated,
		Status: http.StatusText(http.StatusCreated),
		Data:   resp,
	}

	err = serializer.SerializeWriter(w, res.Code, res)
	if err != nil {
		panic(err)
	}
}

func (handler *UserHandlerImpl) Login(w http.ResponseWriter, r *http.Request) {
	var req api.UserRequestLogin

	err := serializer.SerializeRequest[*api.UserRequestLogin](r, &req)
	if err != nil {
		panic(err)
	}
	user := &entity.User{
		Email:    req.Email,
		Password: req.Password,
	}
	resp, err := handler.Usecase.Login(r.Context(), user)
	if err != nil {
		res := api.ErrResponse{
			Code:    commonErr.GetHTTPErrorCode(err),
			Status:  http.StatusText(commonErr.GetHTTPErrorCode(err)),
			Message: err.Error(),
		}
		_ = serializer.SerializeWriter(w, res.Code, res)
		return
	}

	res := api.Response[api.UserResponse]{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   resp,
	}
	err = serializer.SerializeWriter(w, res.Code, res)
	if err != nil {
		panic(err)
	}
}
