package handler

import (
	"context"
	"net/http"

	"github.com/YurcheuskiRadzivon/HSC-pattern/controller"
)

type PropertyHandler interface {
	GetProperty(w http.ResponseWriter, r *http.Request)
	InserProperty(w http.ResponseWriter, r *http.Request)
}

type propertyHandler struct {
	controller controller.PropertyController
	ctx        context.Context
}

func (ph *propertyHandler) GetProperty(w http.ResponseWriter, r *http.Request) {

}
func (ph *propertyHandler) InserProperty(w http.ResponseWriter, r *http.Request) {

}
