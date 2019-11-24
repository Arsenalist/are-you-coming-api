package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
)

func GetEvent(eventHash string) (*Event, error) {
	client := Client()
	result, err := client.HGet("events", eventHash).Result()
	if err == redis.Nil {
		fmt.Println("Event Hash not found " + eventHash)
		return nil, err
	}
	event := &Event{}
	_ = json.Unmarshal([]byte(result), event)
	return event, nil
}

func CreateEvent(event Event) {
	event.GenerateHash()
	SaveEvent(event)
}

func SaveEvent(event Event) {
	client := Client()
	eventJson, _ := json.Marshal(event)
	_, err := client.HSet("events", event.Hash, eventJson).Result()
	if err == redis.Nil {
		fmt.Println("Could not save event")
	}
}

func SaveRsvp(event Event, userId string, name string, rsvpString string) {
	_, err := event.GetRsvp(userId)
	if err == nil {
		fmt.Println("Found RSVP")
		event.UpdateExistingRsvp(userId, rsvpString)
	} else {
		fmt.Println("Did not find RSVP")
		event.AddRsvp(name, userId, rsvpString)
	}
	SaveEvent(event)
}
