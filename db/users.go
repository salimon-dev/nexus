package db

import (
	"database/sql"
	"salimon/nexus/types"
	"time"

	"github.com/google/uuid"
)

// inserts a new user into the database
func InsertUser(username string, email string, password string, credit int32, usage int32, role types.UserRole, status types.UserStatus) (*types.User, error) {
	query := "INSERT INTO users (id, username, email, password, credit, usage, role, status, registered_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"
	id := uuid.New().String()
	registeredAt := time.Now()
	updatedAt := time.Now()

	_, err := DB.Exec(query, id, username, email, password, credit, usage, role, status, registeredAt, updatedAt)
	if err != nil {
		return nil, err
	}
	user := types.User{
		Id:           id,
		Username:     username,
		Email:        email,
		Password:     password,
		Credit:       credit,
		Usage:        usage,
		Role:         role,
		Status:       status,
		RegisteredAt: registeredAt,
		UpdatedAt:    updatedAt,
	}
	return &user, nil
}

// updates a user in database
func UpdateUser(user *types.User) error {
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

func parseSingleUserQuery(rows *sql.Rows, err error) (*types.User, error) {
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var user types.User
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Credit, &user.Usage, &user.Role, &user.Status, &user.RegisteredAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		return &user, nil
	}
	return nil, nil
}

func FindUserByAuth(email string, password string) (*types.User, error) {
	query := "SELECT * FROM users WHERE email=$1 AND password=$2"

	rows, err := DB.Query(query, email, password)
	return parseSingleUserQuery(rows, err)
}

func FindUserByEmail(email string) (*types.User, error) {
	query := "SELECT * FROM users WHERE email=$1"

	rows, err := DB.Query(query, email)
	return parseSingleUserQuery(rows, err)
}

func FindUserByUsername(username string) (*types.User, error) {
	query := "SELECT * FROM users WHERE username=$1"

	rows, err := DB.Query(query, username)
	return parseSingleUserQuery(rows, err)
}

func FindUserById(id string) (*types.User, error) {
	query := "SELECT * FROM users WHERE id=$1"

	rows, err := DB.Query(query, id)
	return parseSingleUserQuery(rows, err)
}

func GetUserPublicObject(user *types.User) types.PublicUser {
	return types.PublicUser{
		Id:           user.Id,
		Username:     user.Username,
		Email:        user.Email,
		Credit:       user.Credit,
		Usage:        user.Usage,
		Role:         types.UserRoleToString(user.Role),
		Status:       types.UserStatusToString(user.Status),
		RegisteredAt: user.RegisteredAt,
		UpdatedAt:    user.UpdatedAt,
	}
}
