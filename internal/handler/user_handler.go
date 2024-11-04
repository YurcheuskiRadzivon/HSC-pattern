package handler

import (
	"context"
	"net/http"

	"github.com/YurcheuskiRadzivon/HSC-pattern/internal/controller"
)

type UserHandler interface {
	GetUser(w http.ResponseWriter, r *http.Request)
	InsertUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
	GetUserPassword(w http.ResponseWriter, r *http.Request)
	LoginUser(w http.ResponseWriter, r *http.Request)
}
type userHandler struct {
	ctx        context.Context
	controller controller.UserController
}

func (us *userHandler) GetUser(w http.ResponseWriter, r *http.Request) {

}
func (us *userHandler) InsertUser(w http.ResponseWriter, r *http.Request) {

}
func (us *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {

}
func (us *userHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {

}
func (us *userHandler) GetUserPassword(w http.ResponseWriter, r *http.Request) {

}
func (us *userHandler) LoginUser(w http.ResponseWriter, r *http.Request) {

}
