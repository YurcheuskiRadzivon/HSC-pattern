package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/YurcheuskiRadzivon/HSC-pattern/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepository interface {
	GetUser(id int) (*model.User, error)
	InsertUser(User model.User) error
	UpdateUser(id int, User model.User) error
	DeleteUser(id int) error
}

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(dbUser, dbPassword, dbHost, dbPort, dbName string) (UserRepository, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to the database: %v", err)
		return nil, err
	}
	return &userRepository{db: db}, nil
}
func (ur *userRepository) GetUser(id int) (*model.User, error) {
	return nil, nil
}
func (ur *userRepository) InsertUser(User model.User) error {
	return nil
}
func (ur *userRepository) UpdateUser(id int, User model.User) error {
	return nil
}
func (ur *userRepository) DeleteUser(id int) error {
	return nil
}
