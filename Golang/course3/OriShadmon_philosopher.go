package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	id        int
	leftCS    *ChopS
	rightCS   *ChopS
}


func (p Philo) eat(ch chan bool) {
	for i := 0; i < 3; i++ {
		<-ch // acquire host permission
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("starting to eat", p.id)
		fmt.Println("finishing eating", p.id)

		p.rightCS.Unlock()
		p.leftCS.Unlock()
		ch <- true // release host permission
	}
	wg.Done()
}

func hostPermission(ch chan bool) {
	// allow no more than 2 philosophers to eat concurrently
	// by limiting the channel to a capacity of 2
	for i := 0; i < 2; i++ {
		ch <- true
	}
}

func main() {
    var philo_num int = 5 // number of philosophers
    var host_cap int = 2

    ch := make(chan bool, host_cap)
    CSticks := make([]*ChopS, philo_num)
	for i := 0; i < philo_num; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, philo_num)
	for i := 0; i < philo_num; i++ {
		philos[i] = &Philo{i + 1, CSticks[i], CSticks[(i+1)%philo_num]}
	}

    go hostPermission(ch)

	for i := 0; i < philo_num; i++ {
		wg.Add(1)
		go philos[i].eat(ch)
	}

	wg.Wait()
	close(ch)
}

/**
 go run Golang/course3/OriShadmon_philosopher.go #> Golang/course3/outputl.txt
starting to eat 1
finishing eating 1
starting to eat 2
finishing eating 2
starting to eat 3
finishing eating 3
starting to eat 5
finishing eating 5
starting to eat 1
finishing eating 1
starting to eat 2
finishing eating 2
starting to eat 4
finishing eating 4
starting to eat 5
finishing eating 5
starting to eat 1
finishing eating 1
starting to eat 3
finishing eating 3
starting to eat 4
finishing eating 4
starting to eat 5
finishing eating 5
starting to eat 3
finishing eating 3
starting to eat 4
finishing eating 4
starting to eat 2
finishing eating 2
*/