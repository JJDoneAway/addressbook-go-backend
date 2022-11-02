package models

import (
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// see https://pkg.go.dev/github.com/go-playground/validator
// for validation
type Address struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey" json:"id"`
	FirstName string `json:"first-name" binding:"required,gt=1"`
	LastName  string `json:"last-name" binding:"required,gt=1"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone" binding:"required,e164"`
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type Entity interface {
	GetAllUsers() []*Address
	InsertUser() error
	GetUserByID() (*Address, error)
	UpdateUser() error
	DeleteUserByID() error
	DeleteAllUsers() error
}

const dsn = "host=localhost user=postgres password=1234 dbname=postgres port=5432 sslmode=disable"

var (

	//error to be thrown in case of inconsistencies
	ErrIdMustBeZero  = errors.New("if you insert new users the ID must be zero")
	ErrUnknownID     = errors.New("id is unknown")
	ErrDuplicatedKey = errors.New("the key we generated already exists")

	db *gorm.DB
)

func init() {

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		panic("Not able to init DB connection")
	}

	db.Debug().Migrator().DropTable(&Address{})
	db.Debug().AutoMigrate(&Address{})

}

func (u *Address) GetAllAddresses() []*Address {
	var ret []*Address
	db.Debug().Find(&ret)
	return ret
}

func (u *Address) InsertAddress() {
	db.Debug().Create(&u)
}

func (u *Address) GetAddressByID() *Address {
	var ret *Address
	db.First(ret, u.ID)
	return ret

}

func (u *Address) UpdateAddress() {

	db.Model(u.GetAddressByID()).Updates(u)
}

func (u *Address) DeleteAddressByID() error {

	return nil
}

func (u *Address) DeleteAllAddresses() error {

	return nil
}
