package repository

import (
	"junebank_v1/entity"

	"gorm.io/gorm"
)

type userDB struct {
	gorm *gorm.DB
}

type UserRepository interface {
	Create(user *entity.User) error
	GetByID(id uint) (*entity.User, error)
}

func InitializeUserRepository(gorm *gorm.DB) UserRepository {
	return &userDB{gorm}
}

func (db *userDB) Create(user *entity.User) error {
	if err := db.gorm.Debug().Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (db *userDB) GetByID(id uint) (*entity.User, error) {
	user := new(entity.User)
	if err := db.gorm.Debug().Where("id = ?", id).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
