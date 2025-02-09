package main

type Device struct {
	id   string
	name string
}

func GetNewDevice(id string, name string) *Device {
	return &Device{
		id:   id,
		name: name,
	}
}