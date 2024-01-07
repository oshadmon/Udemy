package main

import (
	"fmt"
	"sync"
	"time"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	leftCS, rightCS *ChopS
	name            int
}

type Host struct {
	sync.Mutex
	concurrentEat int
}

var wg sync.WaitGroup

func (h *Host) getPermissionEat(p *Philo) bool {
	h.Lock()
	defer h.Unlock()
	if h.concurrentEat >= 2 {
		return false
	}
	h.concurrentEat++
	fmt.Printf("Philosopher %d you can eat\n", p.name)
	return true
}

func (h *Host) doneEating(p *Philo) {
	h.Lock()
	defer h.Unlock()
	h.concurrentEat--
}

func (p Philo) eat(h *Host) {

	//for eat:=0;eat<3;eat++{
	for eat := 0; eat < 3; eat++ {
		fmt.Printf("Philosopher %d would like to eat the %d time\n", p.name, eat)
		for {
			if h.getPermissionEat(&p) {
				break
			}
			fmt.Printf("Philosopher %d WAiting to eat the %d time\n", p.name, eat)
			time.Sleep(20 * time.Millisecond)
		}
		p.leftCS.Lock()
		p.rightCS.Lock()
		fmt.Printf("Philo %d finished to eat the %d time\n\n", p.name, eat)
		h.doneEating(&p)

		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}
	wg.Done()
}

func main() {
	CSticks := make([]*ChopS, 5)

	var h = Host{}
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{leftCS: CSticks[i], rightCS: CSticks[(i+1)%5], name: i}
	}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philos[i].eat(&h)
	}
	wg.Wait()

}
