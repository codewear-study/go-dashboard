package main

import "math/rand"

type Storage struct {
	CPUUtilization    int
	MemoryUtilization int
	CPUTemperature    int
	MemoryTemperature int
	RoomTemperature   int
	MessengerKaKao    int
	MessengerSlack    int
	MessengerDiscord  int
}

func NewStorage() *Storage {
	s := &Storage{}
	go MeasurePercent(&s.CPUUtilization, 1000)
	go MeasurePercent(&s.MemoryUtilization, 1000)
	go MeasureInt(&s.CPUTemperature, 1000)
	go MeasureInt(&s.MemoryTemperature, 1000)
	go MeasureInt(&s.RoomTemperature, 1000)
	s.MessengerKaKao = rand.Intn(100)
	s.MessengerSlack = rand.Intn(50)
	s.MessengerDiscord = rand.Intn(5)
	return s
}

func (s Storage) Retrieve() Storage {
	return s
}
