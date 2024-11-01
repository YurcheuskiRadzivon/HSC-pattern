package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/YurcheuskiRadzivon/HSC-pattern/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepository interface {
	GetUser(nickname string) (*model.User, error)
	InsertUser(User model.UserHash) error
	UpdateUser(id int, User model.UserUpd) error
	DeleteUser(id int) error
	GetUserPassword(nickname string) ([]byte, error)
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
func (ur *userRepository) GetUser(nickname string) (*model.User, error) {
	var User model.User
	query := `SELECT id,name,nickname,email FROM "user" WHERE nickname=$1 `
	err := ur.db.QueryRow(context.Background(), query, User.Nickname).Scan(&User)
	if err != nil {
		return nil, err
	}
	return &User, nil
}
func (ur *userRepository) InsertUser(User model.UserHash) error {
	query := `INSERT INTO "user"(name,nickname,email,password) VALUES($1,$2,$3,$4)`
	_, err := ur.db.Exec(context.Background(), query, User.Name, User.Nickname, User.Email, User.Password)
	if err != nil {
		return err
	}
	return nil
}
func (ur *userRepository) UpdateUser(id int, User model.UserUpd) error {
	query := `UPDATE "user" SET name=$1, nickname=$2, email=$3 WHERE id=$4`
	_, err := ur.db.Exec(context.Background(), query, User.Name, User.Nickname, User.Email, User.ID)
	if err != nil {
		return err
	}
	return nil
}
func (ur *userRepository) DeleteUser(id int) error {
	query := `DELETE FROM "user" WHERE id=$1`
	_, err := ur.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return nil
}
func (ur *userRepository) GetUserPassword(nickname string) ([]byte, error) {
	var hashedPassword []byte
	query := `SELECT password FROM "user" WHERE nickname = $1`
	err := ur.db.QueryRow(context.Background(), query, nickname).Scan(&hashedPassword)
	if err != nil {
		return nil, err
	}
	return hashedPassword, nil

}
