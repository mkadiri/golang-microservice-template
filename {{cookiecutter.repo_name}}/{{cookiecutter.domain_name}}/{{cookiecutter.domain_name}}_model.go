package {{cookiecutter.domain_name}}

import "time"

type {{cookiecutter.domain_name_title}} struct {
	Id 			int 		`json:"id"`
	FirstName 	string 		`json:"first_name"`
	LastName 	string 		`json:"last_name"`
	DateOfBirth time.Time 	`json:"date_of_birth"`
}