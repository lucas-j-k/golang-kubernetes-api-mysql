package note

import "github.com/jmoiron/sqlx"

// NoteService provides a persistence layer to interact with the database
type NoteService struct {
	Db *sqlx.DB
}

// ListNotesForUser retrieves all notes from the DB for a given user ID
func (service NoteService) ListNotesForUser(payload ListNotesByUserPayload) ([]NoteRow, error) {
	marshalledRows := []NoteRow{}
	err := service.Db.Select(
		&marshalledRows,
		"SELECT id, user_id, label, body FROM note WHERE user_id = ?",
		payload.UserId,
	)

	return marshalledRows, err
}

// DeleteNote deletes a note row from the DB, based on combination of note id and user id
func (service NoteService) DeleteNote(payload DeleteNotePayload) (int64, error) {
	_, err := service.Db.Exec(
		"DELETE FROM note WHERE id = ? AND user_id = ?",
		payload.Id,
		payload.UserId,
	)
	if err != nil {
		return 0, err
	}

	return payload.Id, nil
}

// UpdateNoteById updates fields in the DB for a specific note row
func (service NoteService) UpdateNoteById(payload UpdateNotePayload) (int64, error) {
	_, err := service.Db.Exec(
		"UPDATE note SET label = ?, body = ?, row_last_updated = NOW() WHERE id = ? AND user_id = ?",
		payload.Label,
		payload.Body,
		payload.Id,
		payload.UserId,
	)
	if err != nil {
		return 0, err
	}

	return payload.Id, nil
}

// DeleteNote deletes a note row from the DB, based on combination of note id and user id
func (service NoteService) InsertNote(payload InsertNotePayload) (int64, error) {
	result, err := service.Db.Exec(
		"INSERT INTO note (user_id, label, body, row_inserted) VALUES (?, ?, ?, NOW())",
		payload.UserID,
		payload.Label,
		payload.Body,
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
