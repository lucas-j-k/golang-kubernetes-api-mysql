package user

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

// LoginRequestBody defines the incoming login JSON body.
// Implements chi render.Binder
type LoginRequestBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func (schema *LoginRequestBody) Bind(r *http.Request) error {
	return validator.New().Struct(schema)
}

// SignupJSONBody defines the incoming signup JSON body.
// Implements chi render.Binder
type SignupJSONBody struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName"  validate:"required"`
	Password  string `json:"password"  validate:"required,min=8"`
	Email     string `json:"email"  validate:"required,email"`
}

func (schema *SignupJSONBody) Bind(r *http.Request) error {
	return validator.New().Struct(schema)
}

// InsertUserPayload represents the data required to create a user.
type InsertUserPayload struct {
	FirstName    string
	LastName     string
	PasswordHash string
	Email        string
}

// UserRow defines the marshalled struct of a User row from the database
type UserRow struct {
	Id           int64  `db:"id"`
	FirstName    string `db:"first_name"`
	LastName     string `db:"last_name"`
	PasswordHash string `db:"password_hash"`
	Email        string `db:"email"`
}

// UserProfileRow defines the marshalled struct
// of a profile row from the database
type UserProfileRow struct {
	Id        int64  `db:"id" json:"id"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
	Email     string `db:"email" json:"email"`
}

// FindUserByEmailPayload defines the input param struct for the FindUserByEmail func
type FindUserByEmailPayload struct {
	Email string
}

// FindUserByEmailPayload defines the input param struct for the FindUserByEmail func
type GetUserProfilePayload struct {
	Id int64
}
