package models

import (
	"errors"
)

type Address struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"first-name" binding:"required,gt=1"`
	LastName  string `json:"last-name" binding:"required,gt=1"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone" binding:"required,e164"`
}

var (
	//error to be thrown in case of inconsistencies
	ErrIdMustBeZero  = errors.New("if you insert new users the ID must be zero")
	ErrUnknownID     = errors.New("id is unknown")
	ErrDuplicatedKey = errors.New("the key we generated already exists")
)

type AddressCRUD interface {
	GetAllAddresses() []*Address
	InsertAddress() error
	GetAddressByID() (*Address, error)
	UpdateAddress() error
	DeleteAddressByID() error
	DeleteAllAddress() error
}
