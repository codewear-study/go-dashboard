package main

import (
	"math/rand"
	"time"
)

func MeasurePercent(p *int, d time.Duration) {
	for {
		*p = (90**p + 10*rand.Intn(101)) / 100
		time.Sleep(d)
	}
}

func MeasureInt(i *int, d time.Duration) {
	for {
		*i = (95**i + 5*rand.Intn(70)) / 100
		time.Sleep(d)
	}
}
