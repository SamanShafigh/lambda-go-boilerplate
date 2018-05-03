package app

import (
	"database/sql"
)

// UserModel provides the app structure
type UserModel struct {
	db *sql.DB
}

// UserQuery defies the User query structure
type UserQuery struct {
	Username string
}

// User defies the User structure
type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// GetUserModel initialise a User model
func (model *Model) GetUserModel() UserModel {
	return UserModel{db: model.db}
}

// GetUser finds a user
func (model *UserModel) GetUser(query UserQuery) (*User, error) {
	var user User

	err := model.db.QueryRow(
		`SELECT id, username, password FROM user where username = ?`,
		query.Username).Scan(&user.Id, &user.Username, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
