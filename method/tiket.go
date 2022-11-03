package method

import (
	"go_mini-project/database"
	"go_mini-project/model"
	"log"
	"strconv"
	"time"
)

type TiketMethodImpl struct{}

func (n *TiketMethodImpl) GetAll() []model.Tiket {
	rows, err := database.DB.Query("SELECT * FROM ticket ORDER BY created_at")

	if err != nil {
		log.Fatalf("ERROR WHEN FETCHING DATA: %s", err)
	}

	tiket := model.Tiket{}

	ticket := []model.Tiket{}

	for rows.Next() {
		var id int
		var title, description string
		var createdAt time.Time

		err = rows.Scan(&id, &title, &description, &createdAt)

		if err != nil {
			log.Fatalf("ERROR WHEN FETCHING DATA: %s", err)
		}

		tiket.ID = id
		tiket.Title = title
		tiket.Description = description
		tiket.CreatedAt = createdAt

		ticket = append(ticket, tiket)
	}

	return ticket
}

func (c *TiketMethodImpl) GetByID(id string) model.Tiket {
	rows, err := database.DB.Query("SELECT * FROM ticket WHERE id=?", id)

	if err != nil {
		log.Fatalf("ERROR WHEN FETCHING DATA: %s", err)
	}

	foundTiket := model.Tiket{}

	for rows.Next() {
		var id int
		var title, description string
		var createdAt time.Time

		err = rows.Scan(&id, &title, &description, &createdAt)
		if err != nil {
			log.Fatalf("ERROR WHEN FETCHING DATA: %s", err)
		}

		foundTiket.ID = id
		foundTiket.Title = title
		foundTiket.Description = description
		foundTiket.CreatedAt = createdAt
	}

	return foundTiket
}

func (n *TiketMethodImpl) Create(input model.TiketInput) model.Tiket {
	statement, err := database.DB.Prepare("INSERT INTO ticket(title,description) VALUES(?,?)")

	if err != nil {
		log.Fatalf("ERROR WHEN ADDING DATA: %s", err)
	}

	result, err := statement.Exec(input.Title, input.Description)

	if err != nil {
		log.Fatalf("ERROR WHEN ADDING DATA: %s", err)
	}

	insertedId, err := result.LastInsertId()

	if err != nil {
		log.Fatalf("ERROR WHEN ADDING DATA: %s", err)
	}

	rows, err := database.DB.Query("SELECT * FROM ticket WHERE id=?", int(insertedId))

	if err != nil {
		log.Fatalf("ERROR WHEN FETCHING DATA: %s", err)
	}

	createdTiket := model.Tiket{}

	for rows.Next() {
		var id int
		var title, description string
		var createdAt time.Time

		err = rows.Scan(&id, &title, &description, &createdAt)
		if err != nil {
			log.Fatalf("ERROR WHEN FETCHING DATA: %s", err)
		}

		createdTiket.ID = id
		createdTiket.Title = title
		createdTiket.Description = description
		createdTiket.CreatedAt = createdAt
	}

	return createdTiket
}

func (n *TiketMethodImpl) Update(id string, input model.TiketInput) model.Tiket {
	statement, err := database.DB.Prepare("UPDATE ticket SET title=?, description=? WHERE id=?")

	if err != nil {
		log.Fatalf("ERROR WHEN UPDATING DATA: %s", err)
	}

	_, err = statement.Exec(input.Title, input.Description, id)

	if err != nil {
		log.Fatalf("ERROR WHEN UPDATING DATA: %s", err)
	}

	noteId, _ := strconv.Atoi(id)

	rows, err := database.DB.Query("SELECT * FROM ticket WHERE id=?", int(noteId))

	if err != nil {
		log.Fatalf("ERROR WHEN FETCHING DATA: %s", err)
	}

	updatedNote := model.Tiket{}

	for rows.Next() {
		var id int
		var title, description string
		var createdAt time.Time

		err = rows.Scan(&id, &title, &description, &createdAt)
		if err != nil {
			log.Fatalf("ERROR WHEN FETCHING DATA: %s", err)
		}

		updatedNote.ID = id
		updatedNote.Title = title
		updatedNote.Description = description
		updatedNote.CreatedAt = createdAt
	}

	return updatedNote
}

func (n *TiketMethodImpl) Delete(id string) bool {
	statement, err := database.DB.Prepare("DELETE FROM ticket WHERE id=?")

	if err != nil {
		log.Fatalf("ERROR WHEN DELETING DATA: %s", err)
	}

	result, err := statement.Exec(id)

	if err != nil {
		log.Fatalf("ERROR WHEN DELETING DATA: %s", err)
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return false
	}

	return true

}
