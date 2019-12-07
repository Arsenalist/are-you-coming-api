package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func AnEventWithWithOneRsvp() *Event {
	return &Event{"royalrumble", "Royal Rumble", "permalink", []Rsvp{
		{Name: "Zarar",
			UserId:    "zarar",
			EventHash: "royalrumble",
			Rsvp:      "yes"}}, "userid"}
}

func AnEventWithThreeRsvps() *Event {
	return &Event{"royalrumble", "Royal Rumble", "permalink", []Rsvp{
		{Name: "Zarar",
			UserId:    "zarar",
			EventHash: "royalrumble",
			Rsvp:      "yes"},
		{Name: "Jim",
			UserId:    "zarar",
			EventHash: "royalrumble",
			Rsvp:      "yes"},
		{Name: "Ted",
			UserId:    "ted",
			EventHash: "royalrumble",
			Rsvp:      "no"}}, "userid"}
}

func TestEvent_AddRsvp(t *testing.T) {
	e := AnEventWithWithOneRsvp()
	e.AddNewRsvp("Zarar", "zarar", "no")
	assert.Equal(t, 1, len(e.Rsvps), "There should still be only 1 RSVP as User ID is the same")
	e.AddNewRsvp("Ben", "zarar2", "no")
	assert.Equal(t, 2, len(e.Rsvps), "There should be 2 RSVPs")
	assert.Equal(t, "yes", e.Rsvps[1].Rsvp, "The first RSVP should be a yes (unchanged)")
	assert.Equal(t, "no", e.Rsvps[0].Rsvp, "The second RSVP should be a no")
}

func TestEvent_UpdateExistingRsvp(t *testing.T) {
	e := AnEventWithWithOneRsvp()
	e.UpdateExistingRsvp("Zarar", "zarar", "no")
	assert.Equal(t, "no", e.Rsvps[0].Rsvp, "RSVP should be switched to no")
	e.UpdateExistingRsvp("Zarar", "zarar", "yes")
	assert.Equal(t, "yes", e.Rsvps[0].Rsvp, "RSVP should be switched to yes")
	e.UpdateExistingRsvp("Zarar", "zarar", "no")
	assert.Equal(t, "no", e.Rsvps[0].Rsvp, "RSVP should be switched to no")

	e.AddNewRsvp("Jim", "zarar", "no")
	assert.Equal(t, 2, len(e.Rsvps), "Added another RSVP to make it 2")

	e.UpdateExistingRsvp("Jim", "zarar", "yes")
	assert.Equal(t, "yes", e.Rsvps[0].Rsvp, "Should change Jim/zarar RSVP")
	assert.Equal(t, "no", e.Rsvps[1].Rsvp, "Should not change Zarar/zarar RSVP")
}

func TestEvent_GetRsvp(t *testing.T) {
	e := AnEventWithThreeRsvps()
	rsvp, err := e.GetRsvp("Jim", "zarar")
	assert.Equal(t, "zarar", rsvp.UserId, "zarar should be returned as the RSVP")
	assert.Equal(t, "Jim", rsvp.Name, "zarar should be returned as the RSVP")
	assert.Equal(t, nil, err)

	rsvp, err = e.GetRsvp("Zarar", "zarar")
	assert.Equal(t, "zarar", rsvp.UserId, "zarar should be returned as the RSVP")
	assert.Equal(t, nil, err)
	rsvp, err = e.GetRsvp("James", "zarar")
	assert.Equal(t, "", rsvp.UserId, "rsvp.UserId should be empty as nothing was found")
	assert.NotNil(t, err, "There should be an error as no RSVP was found")
}

func TestEvent_DeleteRsvp(t *testing.T) {
	e := AnEventWithThreeRsvps()
	e.DeleteRsvp("Jim", "zarar")
	assert.Equal(t, 2, len(e.Rsvps), "Should be one less")
	assert.Equal(t, "Zarar", e.Rsvps[0].Name, "First one should be Zarar (order retained")

	e.DeleteRsvp("Ted", "ted")
	assert.Equal(t, 1, len(e.Rsvps), "Should be one less")
	assert.Equal(t, "Zarar", e.Rsvps[0].Name, "First one should be Zarar (order retained")

	e.DeleteRsvp("Ted", "tedy")
	assert.Equal(t, 1, len(e.Rsvps), "Should not have deleted anything")

	e.DeleteRsvp("Tedy", "ted")
	assert.Equal(t, 1, len(e.Rsvps), "Should not have deleted anything")

	e.DeleteRsvp("Zarar", "zarar")
	assert.Equal(t, 0, len(e.Rsvps), "Should have deleted the last RSVP")
}

func TestEvent_UpdateEventAttributes(t *testing.T) {
	real := AnEventWithWithOneRsvp()
	from := AnEventWithWithOneRsvp()
	from.Name = "A new name"
	from.Hash = "A new hash"
	from.Permalink = "A new permalink"
	from.Rsvps = []Rsvp{
		{Name: "Jim",
			UserId:    "jim",
			EventHash: "royalrumble",
			Rsvp:      "yes"}}
	real.UpdateEventAttributes(*from)
	assert.Equal(t, "A new name", real.Name, "Name should have been updated based on from")
	assert.Equal(t, "royalrumble", real.Hash, "Hash should remain the same")
	assert.Equal(t, "permalink", real.Permalink, "Permalink should remain the same")
	assert.Equal(t, "Zarar", real.Rsvps[0].Name, "RSVPs should remain the same")
}

func TestEvent_UpdateEventNameWithEmptyString(t *testing.T) {
	real := AnEventWithWithOneRsvp()
	from := AnEventWithWithOneRsvp()
	real.Name = "My Name"
	from.Name = ""
	real.UpdateEventAttributes(*from)
	assert.Equal(t, "My Name", real.Name, "Empty string should get ignored")
}

func TestEvent_SaveRsvp(t *testing.T) {
	e := AnEventWithWithOneRsvp()
	e.SaveRsvp("Zarar", "zarar", "yes")
	assert.Equal(t, 1, len(e.Rsvps), "Should have one RSVP")
	e.SaveRsvp("Zarar", "zarar2", "yes")
	assert.Equal(t, 2, len(e.Rsvps), "Should have two RSVPs")
}

func TestNewEventSuccess(t *testing.T) {
	event, err := NewEvent("raptors vs kings", "myuserid")
	assert.Equal(t, "raptors vs kings", event.Name, "Name is set correctly")
	assert.Equal(t, event.UserId, "myuserid", "User ID is set correctly")
	assert.NotEmpty(t, event.Hash, "Hash should exist")
	assert.NotEmpty(t, event.Permalink, "Permalink should exist")
	assert.Nil(t, err)
	assert.Equal(t, "/", string(event.Permalink[0]), "First character of permalink should be /")
}

func TestNewEventWithEmptyName(t *testing.T) {
	_, err := NewEvent("", "userid")
	assert.NotNil(t, err)
	_, err = NewEvent(" ", "userid")
	assert.NotNil(t, err)
}

func TestNewEventWithEmptyUserId(t *testing.T) {
	_, err := NewEvent("valid name", "")
	assert.NotNil(t, err)
	_, err = NewEvent("valid name", " ")
	assert.NotNil(t, err)
}

func TestIsValidRsvpString(t *testing.T) {
	assert.True(t, isValidRsvpString("yes"), "Yes is allowed as a RsvpString")
	assert.True(t, isValidRsvpString("no"), "No  is allowed as a RsvpString")
	assert.False(t, isValidRsvpString(""), "Empty string is not allowed as an RsvpString")
	assert.False(t, isValidRsvpString(" "), "Blank string is not allowed as an RsvpString")
}
