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

func CreateEvent(name string) *Event {
	event := NewEvent(name)
	SaveEvent(&event)
	return &event
}

func SaveEvent(event *Event) {
	client := Client()
	eventJson, _ := json.Marshal(event)
	_, err := client.HSet("events", event.Hash, eventJson).Result()
	if err == redis.Nil {
		fmt.Println("Could not save event")
	}
}

func SaveRsvp(event *Event, name string, userId string, rsvpString string) {
	event.SaveRsvp(name, userId, rsvpString)
	SaveEvent(event)
}

func DeleteRsvp(event *Event, name string, userId string) {
	event.DeleteRsvp(name, userId)
	SaveEvent(event)
}
