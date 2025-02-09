package main

type User struct {
	id   string
	name string
	location string
	device *Device
	accounts map[string]*Account
	events map[EventType]Event
}

func RegesterUser(id string, name string, location string, device *Device) *User {
	return &User{
		id:   id,
		name: name,
		location: location,
		device: device,
	}
}

func (u *User) getLocation() string {
	return u.location
}

func (u *User) getDevice() *Device {
	return u.device
}

func (u *User) GetAccounts() map[string]*Account {
	return u.accounts
}

func (u *User) AddAccount(account *Account) {
	u.accounts[account.id] = account
}

func (u *User) GetEvents() map[EventType]Event {
	return u.events
}

func (u *User) AddEvent(event Event) {
	u.events[event.GetEventType()] = event
}

func (u *User) GetTransactionLimit(account *Account) float64 {
	return account.GetTransactionLimit()
}