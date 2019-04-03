package sql

import (
	"github.com/jinzhu/gorm"
	"github.com/mvmaasakkers/certificates/database"
)

func GetError(err error) error {
	switch err {
	case gorm.ErrRecordNotFound:
		return database.ErrorObjectNotFound
	}

	chUnique := "UNIQUE constraint failed"

	if len(err.Error()) >= len(chUnique) && err.Error()[:len(chUnique)] == chUnique {
		return database.ErrorDuplicateObject
	}

	return err
}
