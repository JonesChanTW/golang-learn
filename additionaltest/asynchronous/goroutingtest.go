package asynchronous

import "fmt"

// GoroutingTest basic gorouting test
func GoroutingTest() {
	var x int = 0
	fun1 := func() {
		for i := 0; i < 10000; i++ {
			x++
		}
	}

	go fun1()

	for x < 10000 {
		fmt.Println("now x = ", x)
	}
	fmt.Println("Gorouting test 1 finish x = ", x)
}
