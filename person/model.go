package person

import "time"

type Person struct {
	Id          int       `json:"id"`
	FirstName   string    `json:"firstName"`
	MiddleName  string    `json:"middleName"`
	LastName    string    `json:"lastName"`
	Ssn         string    `json:"ssn"`
	DateOfBirth time.Time `json:"dateOfBirth"`
}
