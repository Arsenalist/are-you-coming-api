package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func AnEventWithWithOneRsvp() *Event {
	return &Event{"royalrumble", "Royal Rumble", "permalink", []Rsvp{
		{Name: "Zarar",
			UserId:    "zarar",
			EventHash: "royalrumble",
			Rsvp:      "yes"}}}
}

func TestEvent_GenerateHash(t *testing.T) {
	e := AnEventWithWithOneRsvp()
	got := e.GenerateHash()
	assert.Equal(t, e.Hash, got, "The generated hash and the returned has should be the same")
	assert.Equal(t, fmt.Sprintf("/%s", got), e.Permalink, "Permalink should be slash followed by hash")
}

func TestEvent_AddRsvp(t *testing.T) {
	e := AnEventWithWithOneRsvp()
	e.AddRsvp("Zarar", "zarar", "no")
	assert.Equal(t, 1, len(e.Rsvps), "There should still be only 1 RSVP as User ID is the same")
	e.AddRsvp("Ben", "zarar2", "no")
	assert.Equal(t, 2, len(e.Rsvps), "There should be 2 RSVPs")
	assert.Equal(t, "yes", e.Rsvps[0].Rsvp, "The first RSVP should be a yes (unchanged)")
	assert.Equal(t, "no", e.Rsvps[1].Rsvp, "The second RSVP should be a no")
}

func TestEvent_UpdateExistingRsvp(t *testing.T) {
	e := AnEventWithWithOneRsvp()
	e.UpdateExistingRsvp("zarar", "no")
	assert.Equal(t, "no", e.Rsvps[0].Rsvp, "RSVP should be switched to no")
	e.UpdateExistingRsvp("zarar", "yes")
	assert.Equal(t, "yes", e.Rsvps[0].Rsvp, "RSVP should be switched to yes")
}

func TestEvent_GetRsvp(t *testing.T) {
	e := &Event{"royalrumble", "Royal Rumble", "permalink", []Rsvp{
		{Name: "Zarar",
			UserId:    "zarar",
			EventHash: "royalrumble",
			Rsvp:      "yes"},
		{Name: "Jim",
			UserId:    "jim",
			EventHash: "royalrumble",
			Rsvp:      "no"}}}
	rsvp, err := e.GetRsvp("zarar")
	assert.Equal(t, "zarar", rsvp.UserId, "zarar should be returned as the RSVP")
	assert.Equal(t, nil, err)
	rsvp, err = e.GetRsvp("jim")
	assert.Equal(t, "jim", rsvp.UserId, "jim should be returned as the RSVP")
	assert.Equal(t, nil, err)
	rsvp, err = e.GetRsvp("notauserid")
	assert.Equal(t, "", rsvp.UserId, "rsvp.UserId should be empty as nothing was found")
	assert.NotNil(t, err, "There should be an error as no RSVP was found")

}
