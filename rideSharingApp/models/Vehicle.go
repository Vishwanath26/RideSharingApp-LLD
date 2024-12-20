package models

import "github.com/rs/xid"

type Vehicle struct {
	id             string
	registrationID string
	desc           string
	availableSeats int
}

func (v *Vehicle) Id() string {
	return v.id
}

func (v *Vehicle) SetId(id string) {
	v.id = id
}

func (v *Vehicle) RegistrationID() string {
	return v.registrationID
}

func (v *Vehicle) SetRegistrationID(registrationID string) {
	v.registrationID = registrationID
}

func (v *Vehicle) Desc() string {
	return v.desc
}

func (v *Vehicle) SetDesc(desc string) {
	v.desc = desc
}

func (v *Vehicle) AvailableSeats() int {
	return v.availableSeats
}

func (v *Vehicle) SetAvailableSeats(availableSeats int) {
	v.availableSeats = availableSeats
}

func NewVehicle(registrationID string, desc string, availableSeats int) *Vehicle {
	return &Vehicle{id: xid.New().String(), registrationID: registrationID, desc: desc, availableSeats: availableSeats}
}
