package additionaltest

import (
	"fmt"
)

const runtimes int = 5e9

// DeferTest 用來說明defer的特性
func DeferTest() {
	var x int = 0

	fmt.Println("defercall was started...")
	///defer 會在離開scope時候才執行,所以這行不會在一開始就看到
	defer fmt.Println("This is first defer, then we will in to a scope{...}")

	///defer 並不受到 {} 的限制因此執行順序為這個func中最後一行defer 最先執行
	{
		///defer 呼叫的func 會根據呼叫點所在的位置帶入當時參數所應有的值,所以此時x=0
		defer fmt.Println("2nd defer, print before for loop and x = ", x)
		///跑一堆迴圈,讓輸出等一下
		for i := 0; i < runtimes; i++ {
			x++
		}
		fmt.Println("for loop end, x = ", x)
		///證明defer並不是等到事情都跑完才去找變數值的第二部分
		defer fmt.Println("defer print after for loop and x = ", x)
	}
	for i := 0; i < runtimes; i++ {
		x++
	}
	defer fmt.Println("This is last defer")
}

// DeferTestInFunc 測試defer 在func call中的運作狀況
func DeferTestInFunc() {
	var x int = 0

	///這行只是普通的印出
	fmt.Println("defercall was started...")
	///defer 會在離開scope時候才執行,所以這行不會在一開始就看到
	defer fmt.Println("This is first defer, then we will in to a scope{...}")

	///內部的defer 會比這個func的都早印出,說明defer是以func call結束為界線
	func() {
		///defer 呼叫的func 會根據呼叫點所在的位置帶入當時參數所應有的值,所以此時x=0
		defer fmt.Println("2nd defer, print before for loop and x = ", x)
		///跑一堆迴圈,讓輸出等一下
		for i := 0; i < runtimes; i++ {
			x++
		}
		fmt.Println("for loop end, x = ", x)
		defer fmt.Println("defer print after for loop and x = ", x)
	}()

	for i := 0; i < runtimes; i++ {
		x++
	}
	defer fmt.Println("This is last defer")
}
