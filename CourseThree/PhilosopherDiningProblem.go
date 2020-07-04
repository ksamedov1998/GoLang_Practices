package CourseThree

import (
	"fmt"
	"math/rand"
	"sync"
)

/*
Implement the dining philosopher’s problem with the following constraints/modifications.

There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
The host allows no more than 2 philosophers to eat concurrently.
Each philosopher is numbered, 1 through 5.
When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>”
on a line by itself, where <number> is the number of the philosopher.
When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on
a line by itself, where <number> is the number of the philosopher.
*/
type Philosopher struct {
	eatingNumber   int
	leftChopstick  Chopstick
	rightChopstick Chopstick
}

type Chopstick struct {
	sync.Mutex
}

func initialize(chopsticks []*Chopstick) []Philosopher {
	var philosophers = make([]Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = Philosopher{0, *chopsticks[i], *chopsticks[(i+1)%5]}
	}

	return philosophers
}
func createChopsticks() []*Chopstick {
	chopsticks := make([]*Chopstick, 5)
	for i := 0; i < 5; i++ {
		chopsticks[i] = new(Chopstick)
	}
	return chopsticks
}
func (p Philosopher) finishEating(number int) {
	fmt.Printf("finishing eating %d\n", number)
}

func (p Philosopher) eat(philosopherNumber int) {
	p.eatingNumber++
	fmt.Printf("%d philosopher start eating %d ", philosopherNumber, p.eatingNumber)
	p.leftChopstick.Lock()
	p.rightChopstick.Lock()
	for i := 0; i < rand.Intn(100); i++ {
		fmt.Printf("%d eating\n", philosopherNumber)
	}
	p.leftChopstick.Unlock()
	p.rightChopstick.Unlock()
}

func StartEating() {
	philosophers := initialize(createChopsticks())
	for {
		for i := 0; i < 5; i++ {
			go philosophers[i].eat(i)
		}
	}
}
