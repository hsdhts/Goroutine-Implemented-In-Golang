package golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Husada Hutasoit"
		fmt.Println("Selesai Mengirim Data ke channel")
	}()

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}
