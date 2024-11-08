package hsc_helpers

import (
	"fmt"
	"github.com/YurcheuskiRadzivon/HSC-pattern/api/handler"
	"github.com/YurcheuskiRadzivon/HSC-pattern/internal/controller"
	"github.com/YurcheuskiRadzivon/HSC-pattern/internal/repository"
	"github.com/pkg/errors"
)

func InitializeComponentsUser(dsnStr string) (handler.UserHandler, error) {
	userRepo, err := repository.NewUserRepository(dsnStr)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error DB connection: %v", err))
	}
	userController := controller.NewUserController(userRepo)
	userHandler := handler.NewUserHandler(userController)
	return userHandler, nil
}
func InitializeComponentsProperty(dsnStr string) (handler.PropertyHandler, error) {
	propRepo, err := repository.NewPropertyRepository(dsnStr)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error DB connection: %v", err))
	}
	propController := controller.NewPropertyController(propRepo)
	propHandler := handler.NewPropertyHandler(propController)
	return propHandler, nil
}
