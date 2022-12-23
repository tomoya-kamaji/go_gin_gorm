package controller

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// https://zenn.dev/hsaki/books/golang-concurrency/viewer/basicusage
func TestMutex(t *testing.T) {
	t.Run("並行処理テスト", func(t *testing.T) {
		mutex()
	})
}

func mutex() {
	var mu sync.Mutex
	src := []int{1, 2, 3, 4, 5}
	dst := []int{}
	fmt.Printf("src: %v\n", src)
	fmt.Printf("dst: %v\n", dst)

	for _, s := range src {
		go func(s int) {
			result := s * 2
			mu.Lock()
			dst = append(dst, result)
			mu.Unlock()
		}(s)
	}

	time.Sleep(time.Second)
	fmt.Println(dst)
}
