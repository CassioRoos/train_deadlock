package main

import (
	"github.com/CassioRoos/train_deadlock/deadlock"
	"github.com/CassioRoos/train_deadlock/models"
	"github.com/hajimehoshi/ebiten"
	"log"
	"sync"
)

var (
	trains        [4]*models.Train
	intersections [4]*models.Intersection
)

const trainLength = 70

func update(screen *ebiten.Image) error {
	if !ebiten.IsDrawingSkipped() {
		DrawTracks(screen)
		DrawIntersections(screen)
		DrawTrains(screen)
	}
	return nil
}

func main() {
	for i := 0; i < 4; i++ {
		trains[i] = &models.Train{Id: i, TrainLength: trainLength, Front: 0}
	}

	for i := 0; i < 4; i++ {
		intersections[i] = &models.Intersection{Id: i, Mutex: sync.Mutex{}, LockedBy: -1}
	}

	go deadlock.MoveTrain(trains[0], 300, []*models.Crossing{{Position: 125, Intersection: intersections[0]},
		{Position: 175, Intersection: intersections[1]}})

	go deadlock.MoveTrain(trains[1], 300, []*models.Crossing{{Position: 125, Intersection: intersections[1]},
		{Position: 175, Intersection: intersections[2]}})

	go deadlock.MoveTrain(trains[2], 300, []*models.Crossing{{Position: 125, Intersection: intersections[2]},
		{Position: 175, Intersection: intersections[3]}})

	go deadlock.MoveTrain(trains[3], 300, []*models.Crossing{{Position: 125, Intersection: intersections[3]},
		{Position: 175, Intersection: intersections[0]}})

	if err := ebiten.Run(update, 320, 320, 3, "Trains in a box"); err != nil {
		log.Fatal(err)
	}
}
