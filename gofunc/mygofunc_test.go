package gofunc

import (
	"testing"
	"time"
)

func Test_sendData(t *testing.T) {
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)

	time.Sleep(1e9)
}

func Test_P(t *testing.T) {
	s := make(semaphore, 1)
	go func() {
		t.Log("do do do do~~~")
		//s.Unlock()
		s.Lock()
	}()
	//s.Lock()
	s.Unlock()
}