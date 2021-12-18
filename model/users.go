package model

import (
	"fmt"
	"time"
)

type Users struct {
	ID                  int
	DateUserCreated     *time.Time
	DateUserLastUpdated *time.Time
	FirstName           string
	LastName            string
	Username            string
	Password            string
}

func (u Users) String() string {
	return fmt.Sprintf("id: %d, dateUserCreated: %v, dateUserLastUpdated: %v, firstName: %v, lastName: %v, username: %v, password: %v",
		u.ID, u.DateUserCreated, u.DateUserLastUpdated, u.FirstName, u.LastName, u.Username, u.Password)
}
