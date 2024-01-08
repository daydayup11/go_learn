package once

import (
	"fmt"
	"sync"
	"testing"
)

type Single struct {
	name string
}

var single *Single
var once sync.Once

func TestOnce(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ob := getSingle()
			fmt.Println(i, "get single", ob.name)
		}(i)
	}
	wg.Wait()
}

func getSingle() *Single {
	once.Do(func() {
		fmt.Print("create single")
		single = &Single{
			name: "single",
		}
	})
	return single
}
