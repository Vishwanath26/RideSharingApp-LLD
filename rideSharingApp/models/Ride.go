package models

import "github.com/rs/xid"

type Ride struct {
	id             string
	origin         string
	destination    string
	startTime      int
	endTime        int
	availableSeats int
	status         string //pending, started, completed
	driverID       string
}

func (r *Ride) EndTime() int {
	return r.endTime
}

func (r *Ride) SetEndTime(endTime int) {
	r.endTime = endTime
}

func (r *Ride) Id() string {
	return r.id
}

func (r *Ride) SetId(id string) {
	r.id = id
}

func (r *Ride) Origin() string {
	return r.origin
}

func (r *Ride) SetOrigin(origin string) {
	r.origin = origin
}

func (r *Ride) Destination() string {
	return r.destination
}

func (r *Ride) SetDestination(destination string) {
	r.destination = destination
}

func (r *Ride) StartTime() int {
	return r.startTime
}

func (r *Ride) SetStartTime(startTime int) {
	r.startTime = startTime
}

func (r *Ride) AvailableSeats() int {
	return r.availableSeats
}

func (r *Ride) SetAvailableSeats(availableSeats int) {
	r.availableSeats = availableSeats
}

func (r *Ride) Status() string {
	return r.status
}

func (r *Ride) SetStatus(status string) {
	r.status = status
}

func (r *Ride) DriverID() string {
	return r.driverID
}

func (r *Ride) SetDriverID(driverID string) {
	r.driverID = driverID
}

func NewRide(origin string, destination string, startTime int, endTime int, availableSeats int, driverID string) (*Ride, error) {
	return &Ride{id: xid.New().String(),
		origin:         origin,
		destination:    destination,
		startTime:      startTime,
		endTime:        endTime,
		availableSeats: availableSeats,
		status:         "pending",
		driverID:       driverID,
	}, nil
}

type Location struct {
	latitude  float32
	longitude float32
}
