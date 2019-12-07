package additionaltest

import "fmt"

// SwitchTest test switch
func SwitchTest() {
	var x int = 0

	if false { // Test 1 跟一般C++一樣的switch 用法
		for i := 0; i < 30; i++ {
			if i < 10 {
				x = i % 2
			} else if i < 20 {
				x = i % 3
			} else {
				x = i
			}

			switch x {
			case 0:
				fmt.Println("Test 1 in case 1")
				fmt.Println("Test 1 case 1 line 2")
				// break
				continue
			case 1:
				fmt.Println("Test 1 case 2")
				// break
				continue
			case 2:
				fmt.Println("Test 1 case 3")
				// break
				continue
			}
			fmt.Println("Test 1 default case, I want to break loop...", i)
			fmt.Println("======================")
			break
		}
	}

	if false { ///Test 2 go 才有的用法
		x := 0
		for i := 0; i < 30; i++ {
			fmt.Println("========Start loop Print i = ", i, " ============")
			if i < 10 {
				x = i % 2
			} else if i < 20 {
				x = i % 3
			} else {
				x = i
			}

			switch {
			// case i > 1:
			// 	fmt.Println("When i > 1 then other case are not run, now i = ", i)
			case x == 0:
				fmt.Println("In test 2 case 1 x = ", x)
			case x == 1:
				fmt.Println("In test 2 case 2 x = ", x)
			case i >= 15:
				fmt.Println("In test 2 case 3 i > 15 ?")
			}
		}
	}
}
