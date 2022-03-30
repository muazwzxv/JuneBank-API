package database

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestConnection(t *testing.T) {
	connection := newGorm(&Configuration{
		User:     "muaz",
		Password: "password",
		Host:     "localhost",
		Port:     5432,
		Name:     "Junebank_v1",
	})

	if assert.NotNil(t, connection) && assert.IsType(t, &gorm.DB{}, connection.Orm) {
		t.Log("Gorm Postgres connection passed")
	} else {
		t.Error("Test Failed")
	}
}
