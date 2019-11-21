package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEvent_GenerateHash(t *testing.T) {

	e := &Event{"royalrumble", "Royal Rumble", "permalink", []Rsvp{
		{Name:      "Zarar",
		UserId:    "zarar",
		EventHash: "royalrumble",
		Rsvp:      "yes"}}}

	got := e.GenerateHash()
	assert.Equal(t, e.Hash, got, "The generated hash and the returned has should be the same")
	assert.Equal(t, fmt.Sprintf("/%s", got), e.Permalink, "Permalink should be slash followed by hash")
}

func TestEvent_AddRsvp(t *testing.T) {
	e := &Event{"royalrumble", "Royal Rumble", "permalink", []Rsvp{
		{Name:      "Zarar",
			UserId:    "zarar",
			EventHash: "royalrumble",
			Rsvp:      "yes"}}}

	e.AddRsvp("Zarar", "zarar", "no")
	assert.Equal(t, 1, len(e.Rsvps), "There should still be only 1 RSVP as User ID is the same")

	e.AddRsvp("Ben", "zarar2", "no")
	assert.Equal(t, 2, len(e.Rsvps), "There should be 2 RSVPs")
	assert.Equal(t, "yes", e.Rsvps[0].Rsvp, "The first RSVP should be a yes (unchanged)")
	assert.Equal(t, "no", e.Rsvps[1].Rsvp, "The second RSVP should be a no")
}

func TestEvent_UpdateExistingRsvp(t *testing.T) {
	e := &Event{"royalrumble", "Royal Rumble", "permalink", []Rsvp{
		{Name:      "Zarar",
			UserId:    "zarar",
			EventHash: "royalrumble",
			Rsvp:      "yes"}}}

	e.UpdateExistingRsvp("zarar", "no")
	assert.Equal(t, "no", e.Rsvps[0].Rsvp, "RSVP should be switched to no")
}


