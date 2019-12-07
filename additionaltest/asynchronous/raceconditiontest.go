package asynchronous

import (
	"fmt"
	"sync"
)

// RaeConditionTest test for race condition
func RaeConditionTest() {
	var x int = 0
	var runTimes int = 1000
	res := make(chan int)
	var m sync.Mutex ///互斥鎖

	fun1 := func(sn int) {
		for i := 0; i < 1000; i++ {
			///上鎖要被保護的變數
			m.Lock()
			///被保護的區段...
			x++
			m.Unlock()
			///做完存取後才解鎖
		}
		// time.Sleep(time.Microsecond * time.Duration(rand.Int31()%100))
		res <- sn
	}

	for i := 0; i < runTimes; i++ {
		go fun1(i)
	}
	runs := 0
	for i := 0; i < runTimes; i++ {
		<-res
		// fmt.Printf("get %d times result %v\n", i, ressn)
		runs = i
	}

	fmt.Printf("x = %d, and runs = %d", x, runs)
}
