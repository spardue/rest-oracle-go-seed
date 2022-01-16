package person

import (
	"database/sql"
	"log"
)

func getAll(db *sql.DB) []Person {
	sql := "SELECT * FROM person"

	var person Person
	var people []Person

	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&person.Id, &person.FirstName, &person.MiddleName, &person.LastName, &person.Ssn, &person.DateOfBirth)
		if err != nil {
			log.Fatal(err)
		}
		people = append(people, person)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return people
}

func getById(db *sql.DB, id int) Person {
	sql := `SELECT * FROM person WHERE id = :1`

	var person Person

	err := db.QueryRow(sql, id).Scan(&person.Id, &person.FirstName, &person.MiddleName, &person.LastName, &person.Ssn, &person.DateOfBirth)
	if err != nil {
		log.Fatal(err)
	}
	return person
}
