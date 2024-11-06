package handler

import (
	"context"
	"net/http"

	"github.com/YurcheuskiRadzivon/HSC-pattern/internal/controller"
)

type PropertyHandler interface {
	GetProperty(w http.ResponseWriter, r *http.Request)
	InserProperty(w http.ResponseWriter, r *http.Request)
}

type propertyHandler struct {
	controller controller.PropertyController
	ctx        context.Context
}

func NewPropertyHandler(controller controller.PropertyController) PropertyHandler {
	return &propertyHandler{
		controller: controller,
		ctx:        context.Background(),
	}
}

func (ph *propertyHandler) GetProperty(w http.ResponseWriter, r *http.Request) {

}
func (ph *propertyHandler) InserProperty(w http.ResponseWriter, r *http.Request) {

}
