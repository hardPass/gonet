package concurrent

import (
  "container/list"
	// "log"
	"testing"
)

var (
	nLen = 10000
)

type Client struct {
	validChan chan bool
	valid     int32
}

func BenchmarkInitSlice(b *testing.B) {
	clients := make([]*Client, nLen)

	for i := 0; i < nLen; i++ {
		c := &Client{}
		clients[i] = c
	}

	clients2 := make([]*Client, nLen)

	for i := 0; i < nLen; i++ {
		if i%3 != 0 {
			clients2[i] = clients[i]
		}
	}
	// log.Println(len(clients))
}

func BenchmarkInitSlice2(b *testing.B) {
	clients := make([]*Client, 0, nLen)

	for i := 0; i < nLen; i++ {
		c := &Client{}
		clients = append(clients, c)
	}
	// log.Println(len(clients))
}

func BenchmarkInitList(b *testing.B) {
	clients := list.New()

	for i := 0; i < nLen; i++ {
		c := &Client{}
		clients.PushBack(c)
	}
	// log.Println(clients.Len())

}
