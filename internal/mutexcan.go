package mygorep1

/*
https://github.com/kchugalinskiy/education2024-1
*/
import (
	"fmt"
)

type MyMutex struct {
	//exitCh = make(chan interface{})
	//canal_ chan interface{}
	//thCount atomic.Int32
	canal_ chan int
	count  int
}

func (m MyMutex) Lock() {
	//*&m.thCount++
	//m.thCount.Add(1)
	m.canal_ <- 1
	fmt.Println("lock")
}
func (m MyMutex) Unlock() {
	//m.thCount.Add(-1)
	m.canal_ <- (-1)
	fmt.Println("unlock")
}
func (m MyMutex) Wait() {
	//sum := 0
	//*&sum++
	for true {
		//*&sum += (int(<-m.canal_))
		t := <-m.canal_
		m.count += t
		if m.count != 0 {

		} else {
			break
		}
		/*sum = sum + int(t)
		if sum == 0 {
			break
		}*/
	}
}
func NewMyMutex() *MyMutex {
	return &MyMutex{
		canal_: make(chan int),
		count:  0,
		//thCount: 0,
	}
}
