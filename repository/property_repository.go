package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/YurcheuskiRadzivon/HSC-pattern/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PropertyRepository interface {
	GetProperty(id int) (*model.Property, error)
	InserProperty(Property model.Property) error
	UpdateProperty(id int, Property model.Property) error
	DeleteProperty(id int) error
}
type propertyRepository struct {
	db *pgxpool.Pool
}

func NewPropertyRepository(dbUser, dbPassword, dbHost, dbPort, dbName string) (PropertyRepository, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to the database: %v", err)
		return nil, err
	}
	return &propertyRepository{db: db}, nil

}
func (pr *propertyRepository) GetProperty(id int) (*model.Property, error) {

	return nil, nil

}
func (pr *propertyRepository) InserProperty(Property model.Property) error {
	return nil
}
func (pr *propertyRepository) UpdateProperty(id int, Property model.Property) error {
	return nil
}
func (pr *propertyRepository) DeleteProperty(id int) error {
	return nil
}
