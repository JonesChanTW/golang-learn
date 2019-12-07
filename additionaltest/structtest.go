package additionaltest

// Human is a struct for test
type Human struct {
	Name string
	Age  int
}

/// GetName 類似 成員函式, 這樣寫表示這是給 Human 的func
func (man Human) GetName() string {
	return man.Name
	// return man.name
}

// StructTest for struce test
func StructTest() {

}
