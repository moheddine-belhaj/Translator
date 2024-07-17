package model

import (
	"github.com/google/uuid"
)

type User struct {
	UID		uuid.UUID `db:"uid" json:"uid"`
	Email	string	  `db:"email" json:"email"`
	Password	string	  `db:"Password" json:"-"`
	Name	string	  `db:"Name" json:"Name"`
	ImageURL	string	  `db:"image_url" json:"imageUrl"`
	Website	string	  `db:"website" json:"website"`

}