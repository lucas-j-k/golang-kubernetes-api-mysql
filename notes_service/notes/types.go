package note

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

// InsertNotePayload defines the input params for an insert call
type InsertNotePayload struct {
	UserID int64
	Label  string
	Body   string
}

// NoteRow defines a marshalled Note row from the DB, with JSON tags to marshall for response
type NoteRow struct {
	Id     int64  `db:"id" json:"id"`
	UserId int64  `db:"user_id" json:"user_id"`
	Label  string `db:"label" json:"label"`
	Body   string `db:"body" json:"body"`
}

// ListNotesByUser payload defines the input for listing all notes for a user id
type ListNotesByUserPayload struct {
	UserId int64
}

// DeleteNotePayload
type DeleteNotePayload struct {
	Id     int64
	UserId int64
}

// UpdateNotePayload
type UpdateNotePayload struct {
	Id     int64
	UserId int64
	Label  string
	Body   string
}

// MutateNoteRequestBody defines the incoming login JSON body. Implements chi render.Binder
// Handles both Insert and Update requests
type MutateNoteRequestBody struct {
	Label string `json:"label" validate:"required,min=1"`
	Body  string `json:"body" validate:"required,min=1"`
}

func (schema *MutateNoteRequestBody) Bind(r *http.Request) error {
	return validator.New().Struct(schema)
}
