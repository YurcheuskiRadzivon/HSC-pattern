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
	InserProperty(Property *model.Property) error
	//UpdateProperty(id int, Property model.Property) error
	//DeleteProperty(id int) error
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
	var Property *model.Property
	query := `
    SELECT
    p.Id AS PropertyId,
    a.Country,
    a.City,
    a.Street,
    a.NumOfHome,
    pr.Value,
    pr.Currency
FROM
    Property p
JOIN
    Adrs a ON p.AddressId = a.Id
JOIN
    Price pr ON p.PriceId = pr.Id
WHERE 
    p.Id=$1;`
	err := pr.db.QueryRow(context.Background(), query, id).Scan(&Property.Id, &Property.Address.Country, &Property.Address.City, &Property.Address.Street, &Property.Address.NumOfHome, &Property.Price.Value, &Property.Price.Currency)
	if err != nil {
		return nil, err
	}
	return Property, nil
}
func (pr *propertyRepository) InserProperty(Property *model.Property) error {

	query := `BEGIN;


WITH adrs_insert AS (
    INSERT INTO Adrs (Country, City, Street, NumOfHome)
    VALUES ($1, $2, $3, $4)
    RETURNING Id AS AddressId
),


price_insert AS (
    INSERT INTO Price (Value, Currency)
    VALUES (($5, $6)
    RETURNING Id AS PriceId
)


INSERT INTO Property (AddressId, PriceId)
SELECT AddressId, PriceId
FROM adrs_insert, price_insert;

COMMIT;`
	_, err := pr.db.Exec(context.Background(), query, Property.Address.Country, Property.Address.City, Property.Address.Street, Property.Address.NumOfHome, Property.Price.Value, Property.Price.Currency)
	if err != nil {
		return err
	}
	return nil
}