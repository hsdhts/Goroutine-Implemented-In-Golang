package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("Batman")
	pool.Put("Spiderman")
	pool.Put("Wonderwomen")

	for i := 0; i < 100; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
		}()
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}
