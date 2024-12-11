package models

import (
	"time"

	db "notes.com/app/database"
)

type Note struct {
	Id           int64
	Title        string
	Description  string
	CreationDate time.Time
	UserId       int64
}

func GetAllNotes() ([]Note, error) {
	query := `
	SELECT * FROM notes`

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var notes []Note

	for rows.Next() {
		var note Note

		err := rows.Scan(&note.Id, &note.Title, &note.Description, &note.CreationDate, &note.UserId)

		if err != nil {
			return nil, err
		}

		notes = append(notes, note)
	}

	return notes, nil
}

func GetSingleNote(noteId int64) (*Note, error) {
	query := `SELECT * FROM notes WHERE id = ?`

	row := db.DB.QueryRow(query, noteId)

	var note Note
	err := row.Scan(&note.Id, &note.Title, &note.Description, &note.CreationDate, &note.UserId)

	if err != nil {
		return nil, err
	}

	return &note, nil
}

func (note *Note) CreateNote() error {
	query := `INSERT INTO notes(title, description, creationDate, userId) VALUES (?,?,?,?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	res, err := stmt.Exec(note.Title, note.Description, note.CreationDate, note.UserId)

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	note.Id = id

	return err
}

func (note *Note) UpdateNote() error {
	query := `UPDATE notes SET title = ?, description = ?, creationDate = ?, userId = ? WHERE id = ?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(note.Title, note.Description, note.CreationDate, note.UserId, note.Id)

	if err != nil {
		return err
	}

	return err
}

func DeleteNote(noteId int64) error {
	query := `DELETE FROM notes WHERE id = ?`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(noteId)

	return err
}
