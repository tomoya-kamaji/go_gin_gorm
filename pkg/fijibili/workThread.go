package parallelprocessing

import (
	"fmt"
	"time"
)

func workerThread() {
	ch := make(chan request, 3)

	alice := clientThread{"Alice", ch}
	bobby := clientThread{"Bobby", ch}
	chris := clientThread{"Chris", ch}

	// 送信するgoroutin
	go alice.Run()
	go bobby.Run()
	go chris.Run()

	// 受信するgoroutin：働く方
	for i := 0; i < 10; i++ {
		go worker(ch)
	}

}

type request struct {
	Name   string
	Number int
}

func (r *request) Execute() {
	fmt.Println(" executes ", r)
	time.Sleep(time.Duration(300) * time.Millisecond)
}

func (r *request) String() string {
	return fmt.Sprintf("[ request from %s No. %d ]", r.Name, r.Number)
}

type clientThread struct {
	Name string
	Chan chan request
}

// 送信側：ループでchに入れる
func (c *clientThread) Run() {
	i := 0
	for {
		c.Chan <- request{Name: c.Name, Number: i}
		i++
	}
}

// 受信し続けている
func worker(ch <-chan request) {
	for {
		r := <-ch
		r.Execute()
	}
}
