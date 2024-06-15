package db

import (
	"github.com/Harsh-apk/notesPostgres/types"
	"gorm.io/gorm"
)

// Interface

type UserDB interface {
	CreateNewUser(*types.User) error
	ReadUserById(*uint) (*types.User, error)
	ReadAllUsers() (*[]types.User, error)
	AuthenticateUser(id uint) error
}

// Store that Implements the Interface

type PostgresUserStoreDB struct {
	DB *gorm.DB
}

// Factory function

func NewPostgressUserDb(db *gorm.DB) *PostgresUserStoreDB {
	return &PostgresUserStoreDB{
		DB: db,
	}
}

// Functions implementation

func (p *PostgresUserStoreDB) CreateNewUser(user *types.User) error {
	err := p.DB.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgresUserStoreDB) ReadUserById(id *uint) (*types.User, error) {
	var user types.User
	err := p.DB.First(&user, id).Error
	if err != nil {
		return &user, err
	}

	return &user, nil
}
func (p *PostgresUserStoreDB) ReadAllUsers() (*[]types.User, error) {
	var users []types.User
	err := p.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (p *PostgresUserStoreDB) AuthenticateUser(id uint) error {
	err := p.DB.Model(&types.User{}).Where("id = ?", id).Update("email_auth", true).Error
	if err != nil {
		return err
	}
	return nil
}
