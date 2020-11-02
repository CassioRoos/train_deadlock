package models

import "sync"

type Train struct {
	Id          int
	Front       int
	TrainLength int
}

type Intersection struct {
	Id       int
	Mutex    sync.Mutex
	LockedBy int
}

type Crossing struct {
	Position     int
	Intersection *Intersection
}
