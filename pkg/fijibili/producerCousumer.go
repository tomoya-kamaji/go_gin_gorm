package parallelprocessing

import (
	"fmt"
	"time"
)

func producerCousumer() {
	in := make(chan int)
	out := make(chan int)
	// Workerを4つ作る
	for i := 0; i < 4; i++ {
		go smaileWorker(in, out)
	}
	go produce(in)
	consume(out)
}

// workerPool
func smaileWorker(in, out chan int) {
	for {
		n := <-in
		time.Sleep(500 * time.Millisecond)
		out <- n
	}
}

// 生産物：ジョブを送る 更新されたユーザidを渡す
func produce(in chan<- int) {
	i := 0
	for {
		fmt.Printf("-> Send job: %d\n", i)
		in <- i
		i++
	}
}

// 消費者：上部を調理する
func consume(out <-chan int) {
	for n := range out {
		fmt.Printf("<- Recv job: %d\n", n)
	}
}
