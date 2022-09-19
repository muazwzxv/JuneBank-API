package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestConnection(t *testing.T) {
	connection := newGorm(&DatabaseConfiguration{
		User:     "muaz",
		Password: "password",
		Host:     "localhost",
		Port:     5432,
		Name:     "Junebank_v1",
	})

	if assert.NotNil(t, connection) &&
		assert.IsType(t, &gorm.DB{}, connection.Orm) &&
		assert.IsType(t, &DatabaseConfiguration{}, connection.Config) {

		t.Log("Gorm Postgres connection passed")
	} else {
		t.Error("Test Failed")
	}
}
