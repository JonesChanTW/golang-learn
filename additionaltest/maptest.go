package additionaltest

import "fmt"

// TestMap for test package import
func TestMap() {
	var myMap map[string]int

	myMap = make(map[string]int)
	// 用上兩行的方式宣告 x 這個map,並將它實體化 或者是 x := make(map[string]int)  直接進行短變數宣告+實體化動作
	myMap["help"] = 1
	fmt.Println("x[help] = ", myMap["help"])

	myMap["test"] = 20

	key := "test"
	if val, ok := myMap[key]; ok { ///那個ok 裡面只會放 true/false 用來表示key的存在與否
		fmt.Println("myMap['", key, "'] = ", val)
	} else {
		fmt.Println("myMap['", key, "'] = is not exist ")
	}

	myMap["test2"] = 25
	myMap["test3"] = 40
	fmt.Println("myMap = ", myMap)

	delete(myMap, "test")
	fmt.Println("myMap = ", myMap)
}

// func testMap() {
// 	lotestMap()
// }
