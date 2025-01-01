package db

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id           string    `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Credit       int32     `json:"credit"`
	Usage        int32     `json:"usage"`
	Role         string    `json:"role"`
	RegisteredAt time.Time `json:"registered_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// inserts a new user into the database
func InsertUser(username string, email string, password string, credit int32, usage int32, role string) (*User, error) {
	query := "INSERT INTO users (id, username, email, password, credit, usage, role, registered_at, updated_at) VALUES ($1, $2, $3, $4, $5, $7)"
	id := uuid.New().String()
	registeredAt := time.Now()
	updatedAt := time.Now()

	_, err := DB.Exec(query, id, username, email, password, credit, usage, role, registeredAt, updatedAt)
	if err != nil {
		return nil, err
	}
	user := User{
		Id:           id,
		Username:     username,
		Email:        email,
		Password:     password,
		Credit:       credit,
		Usage:        usage,
		Role:         role,
		RegisteredAt: registeredAt,
		UpdatedAt:    updatedAt,
	}
	return &user, nil
}

// updates a user in database
func UpdateUser(user *User) error {
	updatedAt := time.Now()
	query := "UPDATE users SET username=$1, email=$2, password=$3, credit=$4, usage=$5, role=$6, updated_at=$7 WHERE id=$8"

	_, err := DB.Exec(query, user.Username, user.Password, user.Credit, user.Usage, user.Role, updatedAt, user.Id)
	return err
}

// deletes a user from database
func DeleteUser(id string) error {
	query := "DELETE FROM users WHERE id=$1"
	_, err := DB.Exec(query, id)
	return err
}

func FindUserByAuth(email string, password string) (*User, error) {
	query := "SELECT * FROM users WHERE email=$1 AND password=$2"

	rows, err := DB.Query(query, email, password)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user User
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Credit, &user.Usage, &user.RegisteredAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		return &user, nil
	}
	return nil, nil
}
