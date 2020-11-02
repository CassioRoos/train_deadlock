package arbitrator

import (
	"github.com/CassioRoos/train_deadlock/models"
	"sync"
	"time"
)

var (
	controller = sync.Mutex{}
	cond = sync.NewCond(&controller)
)

func allFree(itl []*models.Intersection) bool{
	for _, it := range itl{
		if it.LockedBy >= 0{
			return false
		}
	}
	return true
}

func lockIntersectionDistance(id, reserveStart, reserveEnd int, crossings []*models.Crossing ){
	var intersectionsToLock[]*models.Intersection
	for _, crossing := range crossings{
		if reserveEnd >= crossing.Position && reserveStart <= crossing.Position && crossing.Intersection.LockedBy != id{
			intersectionsToLock = append(intersectionsToLock, crossing.Intersection)
		}
	}


	controller.Lock()
	for !allFree(intersectionsToLock){
		cond.Wait()
	}
	for _, it := range intersectionsToLock{
		it.LockedBy = id
		// If the slices are not sorted a dead lock can still happen
		time.Sleep(10 * time.Millisecond)
	}
	controller.Unlock()
}

func MoveTrain(train *models.Train, distance int, crossings []*models.Crossing) {
	for train.Front < distance {
		train.Front += 1
		for _, crossing := range crossings {
			if train.Front == crossing.Position {
				lockIntersectionDistance(train.Id, crossing.Position, crossing.Position + train.TrainLength, crossings)
			}
			back := train.Front - train.TrainLength
			if back == crossing.Position {
				controller.Lock()
				crossing.Intersection.LockedBy = -1
				cond.Broadcast()
				controller.Unlock()
			}
		}
		time.Sleep(30 * time.Millisecond)
	}
}