package repository

import (
	"database/sql"
	"go-learn-api/src/modules/people/model"
	"time"
)

type peopleRepository struct {
	db *sql.DB
}

func NewPeopleRepository(db *sql.DB) *peopleRepository {
	return &peopleRepository{db}
}

func (r *peopleRepository) Save(people *model.People) error {
	query := `INSERT INTO "peoples" ("name", "age", "address", "created_at", "updated_at") VALUES ($1, $2, $3, $4, $5)`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(people.Name, people.Age, people.Address, people.CreatedAt, people.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (r *peopleRepository) Update(peopleId int, people *model.People) error {
	query := `UPDATE "peoples" SET "name"=$1, "age"=$2, "address"=$3, "updated_at"=$4 WHERE "id"=$5`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(people.Name, people.Age, people.Address, time.Now(), peopleId)

	if err != nil {
		return err
	}

	return nil
}

func (r *peopleRepository) Delete(peopleId int) error {
	query := `DELETE FROM "peoples" WHERE "id"=$1`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(peopleId)

	if err != nil {
		return err
	}

	return nil
}

func (r *peopleRepository) FindById(peopleId int) (*model.People, error) {
	var people model.People

	query := `SELECT * FROM "peoples" WHERE "id"=$1`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	err = statement.QueryRow(peopleId).Scan(&people.Id, &people.Name, &people.Age, &people.Address, &people.CreatedAt, &people.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &people, nil
}

func (r *peopleRepository) FindAll() (*model.Peoples, error) {
	var peoples model.Peoples

	query := `SELECT * FROM "peoples"`

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var people model.People
		err = rows.Scan(&people.Id, &people.Name, &people.Age, &people.Address, &people.CreatedAt, &people.UpdatedAt)

		if err != nil {
			return nil, err
		}

		peoples = append(peoples, people)
	}

	return &peoples, nil
}
