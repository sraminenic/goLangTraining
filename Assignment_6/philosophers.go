package main

import (
	"fmt"
	"sync"
	"time"
)

type Fork struct{ sync.Mutex }

type Philosopher struct {
	id                  int
	leftFork, rightFork *Fork
}

// philosoper at time dining will lock two forks and release after eating
func (p Philosopher) dine(wg *sync.WaitGroup) {
	defer wg.Done()
	p.leftFork.Lock()
	p.rightFork.Lock()

	fmt.Println(p.id, " is eating")
	//used pause values to minimise.
	time.Sleep(2 * time.Second)
	p.rightFork.Unlock()
	p.leftFork.Unlock()

}

func main() {
	// Number of philosophers and forks
	count := 5
	// Forks Creation Logic
	forks := make([]*Fork, count)
	for i := 0; i < count; i++ {
		forks[i] = new(Fork)
	}
	//waits for a collection of goroutines to finish
	wg := &sync.WaitGroup{}
	// philospoher creation with 2 fork assignment. send to dinning.
	philosophers := make([]*Philosopher, count)
	for i := 0; i < count; i++ {
		wg.Add(1)
		philosophers[i] = &Philosopher{
			id: i, leftFork: forks[i], rightFork: forks[(i+1)%count]}
		go philosophers[i].dine(wg)
	}
	wg.Wait()
}
