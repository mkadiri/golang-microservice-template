package usr

import "time"

type User struct {
	Id 			int 		`json:"id"`
	FirstName 	string 		`json:"first_name"`
	LastName 	string 		`json:"last_name"`
	DateOfBirth time.Time 	`json:"date_of_birth"`
}