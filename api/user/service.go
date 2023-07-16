package user

import (
	"github.com/jmoiron/sqlx"
)

// UserService provides a persistence layer to interact with the database
type UserService struct {
	Db *sqlx.DB
}

// InsertUser inserts a new user into the database. If succesful, returns the inserted record ID
func (service UserService) InsertUser(payload InsertUserPayload) (int64, error) {
	result, err := service.Db.Exec(
		"INSERT INTO user (first_name, last_name, email, password_hash, row_inserted) VALUES (?, ?, ?, ?, NOW())",
		payload.FirstName,
		payload.LastName,
		payload.Email,
		payload.PasswordHash,
	)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

// FindUserByEmail searches for a specific user record by email
func (service UserService) FindUserByEmail(payload FindUserByEmailPayload) ([]UserRow, error) {
	marshalledRows := []UserRow{}

	rows, err := service.Db.Queryx(
		"SELECT id, first_name, last_name, email, password_hash FROM user WHERE email = ?",
		payload.Email,
	)
	if err != nil {
		return marshalledRows, err
	}

	for rows.Next() {
		var marshalled UserRow
		err := rows.StructScan(&marshalled)
		if err != nil {
			return marshalledRows, err
		}
		marshalledRows = append(marshalledRows, marshalled)
	}

	return marshalledRows, nil
}

// GetUserProfile returns the basic profile information for a single user
func (service UserService) FindUserProfile(payload GetUserProfilePayload) (UserProfileRow, error) {
	result := UserProfileRow{}
	err := service.Db.Get(
		&result,
		`SELECT
			id, first_name, last_name, email
			FROM user
			WHERE id = ?`,
		payload.Id,
	)
	return result, err
}
