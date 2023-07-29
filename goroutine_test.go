package golang_goroutine

import (
	"fmt"
	"sync"
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

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}
func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock", user1.Name)
	user1.Change(amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}
