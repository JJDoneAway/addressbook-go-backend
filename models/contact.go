/*
 * Contacts
 *
 * Simple AddressBook API to manage contacts
 *
 * API version: 1.0
 * Contact: john.doe@mail.schwarz
 */

package models

import (
	"errors"
	"time"

	"github.com/sony/sonyflake"
)

// Represents a Contact entry.
type Contact struct {

	// uniqueID of the contact
	ID uint64 `json:"uniqueId,omitempty"`

	// first name of the contact
	FirstName string `json:"firstName"`

	// last name of the contact
	LastName string `json:"lastName"`

	// street of home address of the contact
	Street string `json:"street,omitempty"`

	// ZIP code of home address of the contact
	ZipCode int32 `json:"zipCode,omitempty"`

	// city of home address of the contact
	City string `json:"city,omitempty"`

	// email address of the contact
	Email string `json:"email,omitempty"`

	// phone number of the contact
	Phone string `json:"phone,omitempty"`
}

var (

	//sonyflake for ID generation
	sf *sonyflake.Sonyflake

	//error to be thrown in case of inconsistencies
	ErrIdMustBeZero  = errors.New("if you insert new users the ID must be zero")
	ErrUnknownID     = errors.New("id is unknown")
	ErrDuplicatedKey = errors.New("the key we generated already exists")

	cache map[uint64]*Contact
)

func init() {
	cache = make(map[uint64]*Contact)

	var st sonyflake.Settings
	st.StartTime = time.Now()

	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		panic("sonyflake not created")
	}
}

func NextID() uint64 {
	id, _ := sf.NextID()
	return id
}

func (c *Contact) GetAllContacts() []*Contact {
	ret := make([]*Contact, 0, len(cache))
	for _, v := range cache {
		ret = append(ret, v)
	}
	return ret
}

func (c *Contact) InsertContact() error {
	if c.ID != 0 {
		return ErrIdMustBeZero
	}

	c.ID, _ = sf.NextID()

	//as the id generator ensures unique IDs this is a bit paranoid
	//but better save than sorry
	if cache[c.ID] != nil {
		return ErrDuplicatedKey
	}

	cache[c.ID] = c
	return nil
}

func (c *Contact) GetContactByID() (*Contact, error) {
	ret := cache[c.ID]
	if ret == nil {
		return nil, ErrUnknownID
	}

	return ret, nil

}

func (c *Contact) UpdateContact() error {
	if cache[c.ID] == nil {
		return ErrUnknownID
	}
	cache[c.ID] = c
	return nil
}

func (c *Contact) DeleteContactByID() error {
	if cache[c.ID] == nil {
		return ErrUnknownID
	}

	delete(cache, c.ID)
	return nil
}

func (c *Contact) DeleteAllContacts() error {
	cache = make(map[uint64]*Contact)
	return nil
}
