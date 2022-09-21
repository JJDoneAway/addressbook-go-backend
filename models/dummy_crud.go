package models

import (
	"time"

	"github.com/sony/sonyflake"
)

var (

	//sonyflake for ID generation
	sf *sonyflake.Sonyflake

	cache map[uint64]*Address
)

func init() {
	cache = make(map[uint64]*Address)

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

func (u *Address) GetAllAddresses() []*Address {
	ret := make([]*Address, 0, len(cache))
	for _, v := range cache {
		ret = append(ret, v)
	}
	return ret
}

func (u *Address) InsertAddress() error {
	if u.ID != 0 {
		return ErrIdMustBeZero
	}

	u.ID, _ = sf.NextID()

	//as the id generator ensures unique IDs this is a bit paranoid
	//but better save than sorry
	if cache[u.ID] != nil {
		return ErrDuplicatedKey
	}

	cache[u.ID] = u
	return nil
}

func (u *Address) GetAddressByID() (*Address, error) {
	ret := cache[u.ID]
	if ret == nil {
		return nil, ErrUnknownID
	}

	return ret, nil

}

func (u *Address) UpdateAddress() error {
	if cache[u.ID] == nil {
		return ErrUnknownID
	}
	cache[u.ID] = u
	return nil
}

func (u *Address) DeleteAddressByID() error {
	if cache[u.ID] == nil {
		return ErrUnknownID
	}

	delete(cache, u.ID)
	return nil
}

func (u *Address) DeleteAllAddress() error {
	cache = make(map[uint64]*Address)
	return nil
}
