package main

type EventType string

const (

	EventTypeLogin EventType = "login"
	EventTypeTransaction EventType = "transaction"
)

type Event interface {
	GetEventType() EventType
}
