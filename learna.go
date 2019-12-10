package main

import (
	"fmt"

	"learn/additionaltest"

	"learn/additionaltest/asynchronous"

	_ "github.com/go-sql-driver/mysql"
)

// type Human additionaltest.Human

var n = 111 ///這應該是相對Package內部的全域變數了,整個Package都看的到他
// x := 100	/// := 只提供區域變數,最少得在func內部,但也可以是 func 內部的 {} 裡面使用,反正他一旦離開最接近的{} 區塊就無效了
func main() {
	var n = 1 ///這只有 main() 內部有效

	if !true { ///variaable and scope
		{
			var n int = 10 ///這是這個{} 區間內的變數

			n++
			fmt.Println("n in {} = ", n)
		}
		fmt.Println("n out {}= ", n)
	}

	if !true { /// array
		arr := [5]int{5, 6, 7, 8, 9}

		fmt.Println("arr = ", arr)
		arr[0] = 10
		fmt.Println("arr = ", arr)
	}

	if !true { ///for loop
		additionaltest.LoopTest()
	}

	if !true { ///Map
		additionaltest.TestMap()
	}

	if !true { // switch test
		additionaltest.SwitchTest()
	}

	if !true { ///struct
		someone := additionaltest.Human{Name: "smith", Age: 29}
		so2 := additionaltest.Human{}

		so2.Age = 30
		so2.Name = "John"
		fmt.Println("a man ", someone.GetName())
		fmt.Println("another people", so2)
	}

	if true { ///asynchronous
		if !true { ///普通的gorouting test
			asynchronous.GoroutingTest()
		}

		if !true { ///test for race condition
			asynchronous.RaeConditionTest()
		}

		if !true {
			asynchronous.WaitGroupTest()
		}

		if !true {
			asynchronous.ChannelTest()
		}

		if !true { ///channel buffer / unbuffer test
			asynchronous.ChannelBufferOrUnbufferTest()
		}

		if !true {
			asynchronous.ContextWithCancelTest()
		}

		if true {
			asynchronous.ContextWithCancelTest()
		}
	}

	if !true { ///panic aand recover
		// panic(1)
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("Recive Error with panic %v", err)
			}
		}()
		throwPanic()
		panic(2) ///事實上這行根本不會執行

	}

	if !true { ///defer test 驗證defer 特性
		if !true {
			additionaltest.DeferTest()
		}
		if !true {
			additionaltest.DeferTestInFunc()
		}
	}

	if !true { /// file operator
		additionaltest.FileOpenOrCreateAndWrite("Hello file")
	}

	if !true { ///database

	}

	// var scanStr string = ""
	// res, err := fmt.Scanln(scanStr)
	// fmt.Println("scanStr = ", scanStr, " and res = ", res, " and err = ", err)
}

func showType(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case int8:
		return "int8"
	case string:
		return "string"
		// default:
	}
	return "unknow"
}
