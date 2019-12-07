package additionaltest

import "fmt"

// LoopTest 迴圈測試
func LoopTest() {
	arr := [4]int{4, 5, 6, 7}

	fmt.Println("arr = ", arr)
	for i := 0; i < len(arr); i++ {
		arr[i] = i + 5
	}
	for i, val := range arr {
		fmt.Println("run times", i+1, "val = ", val)
	}
}
