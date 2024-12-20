package models

import "github.com/rs/xid"

type User struct {
	id        string
	name      string
	vehicleID string
}

func (u *User) Id() string {
	return u.id
}

func (u *User) SetId(id string) {
	u.id = id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) VehicleID() string {
	return u.vehicleID
}

func (u *User) SetVehicleID(vehicleID string) {
	u.vehicleID = vehicleID
}

func NewUser(name string) *User {
	return &User{
		id:   xid.New().String(),
		name: name,
	}
}
