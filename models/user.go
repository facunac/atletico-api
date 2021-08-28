package models

import (
	"time"
)

type User struct {
	id          uint
	Email       string
	Rut         string
	Typeuser    string
	Password    string
	Created_on  time.Time
	Last_login  time.Time
	First_name  string
	Last_name   string
	Phone       string
	Birthdate   time.Time
	Gender      string
	Nationality string
	Graduate    uint
}
