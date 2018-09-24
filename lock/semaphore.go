package lock

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type empty struct{}
type Semaphore chan empty

const size = 50

func (s Semaphore) Lock() {
	s <- empty{}
}
func (s Semaphore) Unlock() {
	<-s
}

func New() Semaphore {
	return make(Semaphore, size)
}

/*
use semaphore to limit the number of times some task performs concurrently
*/
func PerformSemaphoreExample() {
	wg := &sync.WaitGroup{}
	sem := New()
	for i := 0; i < 1000; i++ {
		num := i
		sem.Lock()
		wg.Add(1)
		go func() {
			DoStuff(num)
			sem.Unlock()
			wg.Done()
		}()

	}
	wg.Wait()
}

func DoStuff(num int) {
	r := rand.Intn(2) + 1
	time.Sleep(time.Duration(r) * time.Second)
	fmt.Printf("Number is %d time is %v \n", num, time.Now())
}
