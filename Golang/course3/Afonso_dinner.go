package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	id    int
	chopS []*ChopS
}

func (p Philo) eat(permission chan bool, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		// Obtain permission from host
		<- permission

		// Pick the chopsticks in random order
		for j := range rand.Perm(2) {
			p.chopS[j].Lock()
		}

		fmt.Printf("starting to eat %v\n", p.id)
		fmt.Printf("finished eating %v\n", p.id)

		// Put down chopstick
		for j := 0; j < 2; j++ {
			p.chopS[j].Unlock()
		}
	}
	wg.Done()
}

func finishedEating(allowed []int) bool {
	for _, a := range allowed {
		if a > 0 {
			return false
		}
	}

	return true
}

// Host that allows philosphers to eat
func allowEat(channels []chan bool, wg *sync.WaitGroup) {
	allowed := make([]int, 5)
	for i := 0; i < 5; i++ {
		allowed[i] = 3 // Each philosopher is allowed to eat 3 times
	}

	// Allow up to two philosophers to eat concurrently
	for !finishedEating(allowed) {
		who := rand.Perm(5)
		for i := 0; i < 2; i++ {
			if allowed[who[i]] > 0 {
				channels[who[i]] <- true
				allowed[who[i]]--
			}
		}
	}

	wg.Done()
}

func main() {
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		cs := make([]*ChopS, 2)
		cs[0] = CSticks[i]
		cs[1] = CSticks[(i+1)%5]
		philos[i] = &Philo{i, cs}
	}

	permissions := make([]chan bool, 5)
	for i := 0; i < 5; i++ {
		permissions[i] = make(chan bool)
	}

	var wg sync.WaitGroup
	wg.Add(6)
	for i := 0; i < 5; i++ {
		go philos[i].eat(permissions[i], &wg)
	}
	go allowEat(permissions, &wg)
	wg.Wait()
}
