package golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunGoroutine() {
	fmt.Println("Run Goroutine!")
}

func TestCreateGoroutine(t *testing.T) {
	go RunGoroutine()
	fmt.Println("Error message!")

	time.Sleep(1 * time.Nanosecond)

}

func DisplayName(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayName(i)
	}

	time.Sleep(5 * time.Second)
}
