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
