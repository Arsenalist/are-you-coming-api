package main

import (
	"fmt"
	"github.com/dchest/uniuri"
)

type Event struct {
	Hash      string `json:"hash"`
	Name      string `json:"name"`
	Permalink string `json:"permalink"`
	Rsvps     []Rsvp `json:"rsvps"`
}

type Rsvp struct {
	Name      string `json:"name"`
	UserId    string `json:"userId"`
	EventHash string `json:"eventHash"`
	Rsvp      string `json:"rsvp"`
}

func (e *Event) GenerateIdentity() string {
	e.Hash = uniuri.New()
	e.Permalink = "/" + e.Hash
	return e.Hash
}

func (e *Event) SaveRsvp(name string, userId string, rsvpString string) {
	_, err := e.GetRsvp(userId)
	if err == nil {
		e.UpdateExistingRsvp(name, userId, rsvpString)
	} else {
		e.AddNewRsvp(name, userId, rsvpString)
	}
}

func (e *Event) AddNewRsvp(name string, userId string, rsvpString string) {
	_, err := e.GetRsvp(userId)
	if err == nil {
		return
	}
	e.Rsvps = append(e.Rsvps,
		Rsvp{
			name,
			userId,
			e.Hash,
			rsvpString})
}

func (e *Event) UpdateExistingRsvp(name string, userId string, rsvpString string) {
	rsvp, err := e.GetRsvp(userId)
	if err == nil {
		rsvp.Rsvp = rsvpString
		rsvp.Name = name
	}
}

func (e *Event) GetRsvp(userId string) (rsvp *Rsvp, err error) {
	for i, element := range e.Rsvps {
		if element.UserId == userId {
			return &e.Rsvps[i], nil
		}
	}
	return &Rsvp{}, fmt.Errorf("could not find RSVP based on UserID")
}

// Copy the updateable fields on the event
func (e *Event) UpdateEventAttributes(from Event) {
	e.Name = from.Name
}
