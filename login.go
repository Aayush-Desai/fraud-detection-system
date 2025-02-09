package main

type Login struct {
	id   string
	user *User
	password string
}

func GetNewLogin(id string, user *User, password string) *Login {
	return &Login{
		id: id,
		user: user,
		password: password,
	}
}

func (t *Login) GetEventType() EventType {
	return EventTypeLogin
}

