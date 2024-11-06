package controller

import (
	"context"

	"github.com/YurcheuskiRadzivon/HSC-pattern/internal/repository"
	"github.com/YurcheuskiRadzivon/HSC-pattern/model"
)

type PropertyController interface {
	GetProperty(ctx context.Context, id int) (*model.Property, error)
	InserProperty(ctx context.Context, Property *model.Property) error
}
type propertyController struct {
	repo repository.PropertyRepository
}

func NewPropertyController(repo repository.PropertyRepository) PropertyController {
	return &propertyController{repo: repo}

}
func (pc *propertyController) GetProperty(ctx context.Context, id int) (*model.Property, error) {
	Property, err := pc.repo.GetProperty(id)
	if err != nil {
		return nil, err
	}
	return Property, nil
}
func (pc *propertyController) InserProperty(ctx context.Context, Property *model.Property) error {
	err := pc.repo.InsertProperty(Property)
	if err != nil {
		return err
	}
	return nil
}
