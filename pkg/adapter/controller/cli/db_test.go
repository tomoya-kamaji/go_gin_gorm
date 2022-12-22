package controller

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestDo(t *testing.T) {
	t.Run("並行処理テスト", func(t *testing.T) {
		main()
	})
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func omikuji(c1, c2, c3 chan int) {
	for i := 0; i < 10; i++ {
		select {
		case <-c1:
			fmt.Println("大吉")
		case <-c2:
			fmt.Println("中吉")
		case <-c3:
			fmt.Println("小吉")
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ch1 := make(chan rune)
	ch2 := make(chan int)
	done := make(chan struct{})

	go func() {
		for {
			select {
			case c := <-ch1:
				fmt.Printf("[R1] %c\n", c)
			case i := <-ch2:
				fmt.Println("[R2]", i)
			case <-done:
				fmt.Println("Done!")
				return
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for c := 'A'; c <= 'C'; c++ {
			ch1 <- c
			time.Sleep(time.Second)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i < 5; i++ {
			ch2 <- i
			time.Sleep(500 * time.Millisecond)
		}
	}()

	wg.Wait()

	close(done)

	time.Sleep(time.Second)
}
