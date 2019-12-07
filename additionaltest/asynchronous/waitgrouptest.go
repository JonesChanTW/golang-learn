package asynchronous

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// BASIC_RUN_TIME 最少跑幾次
const BASIC_RUN_TIME int64 = 15000000000

// WaitGroupTest test for wait gorouting exit
// 順便以此程式證明 gorouting 可以確實使用到機器的多核心,並非模擬的多執行續,僅占用一核心XD
func WaitGroupTest() {
	var wg sync.WaitGroup

	task := func(wg *sync.WaitGroup, sn int) {
		runtimes := rand.Int63n(BASIC_RUN_TIME) + BASIC_RUN_TIME
		var i int64 = 0
		var x int64 = 0
		fmt.Printf("task %d Staart and run %d times\n", sn, runtimes)

		//time.Sleep(time.Nanosecond)
		for i = 0; i < runtimes; i++ {
			x++
		}
		defer wg.Done()
		fmt.Printf("task %d has finish \n", sn)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go task(&wg, i)
	}

	fmt.Println(time.Now(), " All gorouting was runnung... ")
	defer wg.Wait()
	fmt.Println(time.Now(), " All task has finish")
}
