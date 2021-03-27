package model

import "github.com/google/uuid"


// User defines domain model and its json and db representation
type User struct {
	UID 			uuid.UUID 	`json:"uid" db:"uid"`  // `` - struct tags
	Email   		string		`json:"email" db:"email"`
	Password		string   	`json:"password" db:"password"`
	Name 			string 		`json:"name" db:"name"`
	ImageURL		string 		`json:"image_url" db:"imageURL"`
	Website			string 		`json:"website" db:"website"`
}