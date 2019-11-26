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

func NewEvent(name string) Event {
	event := Event{}
	event.Name = name
	event.Rsvps = []Rsvp{}
	event.Hash = uniuri.New()
	event.Permalink = "/" + event.Hash
	return event
}

func (e *Event) SaveRsvp(name string, userId string, rsvpString string) {
	_, err := e.GetRsvp(name, userId)
	if err == nil {
		e.UpdateExistingRsvp(name, userId, rsvpString)
	} else {
		e.AddNewRsvp(name, userId, rsvpString)
	}
}

func (e *Event) AddNewRsvp(name string, userId string, rsvpString string) {
	_, err := e.GetRsvp(name, userId)
	if err == nil {
		return
	}
	newRsvp := []Rsvp{
		{Name: name,
			UserId:    userId,
			EventHash: e.Hash,
			Rsvp:      rsvpString}}
	e.Rsvps = append(newRsvp, e.Rsvps...)
}

func (e *Event) UpdateExistingRsvp(name string, userId string, rsvpString string) {
	rsvp, err := e.GetRsvp(name, userId)
	if err == nil {
		rsvp.Rsvp = rsvpString
		rsvp.Name = name
	}
}

func (e *Event) GetRsvp(name string, userId string) (rsvp *Rsvp, err error) {
	for i, element := range e.Rsvps {
		if element.UserId == userId && element.Name == name {
			return &e.Rsvps[i], nil
		}
	}
	return &Rsvp{}, fmt.Errorf("could not find RSVP based on UserID and Name")
}

// Copy the updateable fields on the event
func (e *Event) UpdateEventAttributes(from Event) {
	e.Name = from.Name
}
