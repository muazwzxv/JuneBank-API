package domain

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type User struct {
	ID        uint64       `json:"id"`
	FirstName string       `json:"first_name"`
	LastName  string       `json:"last_name"`
	Email     string       `json:"email"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"upated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}

type CreateUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func (c *User) ToJsonStr() ([]byte, error) {
	if res, err := json.Marshal(c); err != nil {
		return nil, err
	} else {
		return res, nil
	}
}

func (c *CreateUser) Bind(r *http.Request) error {
	if c.FirstName == "" {
		return errors.New("missing required first name")
	}

	if c.LastName == "" {
		return errors.New("missing required last name")
	}

	if c.Email == "" {
		return errors.New("missing required email")
	}

	return nil
}
