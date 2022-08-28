package models

import (
	"errors"
	"time"

	"github.com/sony/sonyflake"
)

type User struct {
	ID        uint64 `json:"ID"`
	FirstName string `json:"FirstName" binding:"required"`
	LastName  string `json:"LastName" binding:"required"`
}

type Entity interface {
	GetAllUsers() []*User
	InsertUser() error
	GetUserByID() (*User, error)
	UpdateUser() error
	DeleteUserByID() error
	DeleteAllUsers() error
}

var (

	//sonyflake for ID generation
	sf *sonyflake.Sonyflake

	//error to be thrown in case of inconsistencies
	ErrIdMustBeZero  = errors.New("if you insert new users the ID must be zero")
	ErrUnknownID     = errors.New("id is unknown")
	ErrDuplicatedKey = errors.New("the key we generated already exists")

	cache map[uint64]*User
)

func init() {
	cache = make(map[uint64]*User)

	var st sonyflake.Settings
	st.StartTime = time.Now()

	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		panic("sonyflake not created")
	}
}

func (u *User) GetAllUsers() []*User {
	ret := make([]*User, 0, len(cache))
	for _, v := range cache {
		ret = append(ret, v)
	}
	return ret
}

func (u *User) InsertUser() error {
	if u.ID != 0 {
		return ErrIdMustBeZero
	}

	u.ID, _ = sf.NextID()

	//as the id generator ensures unique IDs this is a bit paranoid
	//but better save than sory
	if cache[u.ID] != nil {
		return ErrDuplicatedKey
	}

	cache[u.ID] = u
	return nil
}

func (u *User) GetUserByID() (*User, error) {
	ret := cache[u.ID]
	if ret == nil {
		return nil, ErrUnknownID
	}

	return ret, nil

}

func (u *User) UpdateUser() error {
	if cache[u.ID] == nil {
		return ErrUnknownID
	}
	cache[u.ID] = u
	return nil
}

func (u *User) DeleteUserByID() error {
	if cache[u.ID] == nil {
		return ErrUnknownID
	}

	delete(cache, u.ID)
	return nil
}

func (u *User) DeleteAllUsers() error {
	cache = make(map[uint64]*User)
	return nil
}
