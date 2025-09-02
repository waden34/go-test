package main

import (
	"fmt"
	"maps"
)

func main() {
	// var mapVariable map[keyType]valueType
	// mapVariable = make(map[keyType]valueType]

	// using a Map Literal
	// mapVariable = map[keyType]valueType{
	// key1: value1,
	// key2: value2
	// }

	myMap := make(map[string]int)

	fmt.Println(myMap)
	myMap["key1"] = 9
	myMap["code"] = 18
	fmt.Println(myMap)
	fmt.Println(myMap["key1"]) // Invalid key will return default value for type

	delete(myMap, "key1")
	fmt.Println(myMap)
	myMap["key2"] = 9
	myMap["key4"] = 18
	myMap["key3"] = 98
	fmt.Println(myMap)

	//clear(myMap)
	//fmt.Println(myMap)

	_, ok := myMap["key2"]
	if ok {
		fmt.Println("A value exists with key2")
	} else {
		fmt.Println("A value does not exist with key2")
	}
	//fmt.Println(value)
	// fmt.Println("Is a value associated with key2:", ok)

	myMap2 := map[string]int{"a": 1, "b": 2}
	fmt.Println(myMap2)
	myMap3 := map[string]int{"a": 1, "b": 2}
	if maps.Equal(myMap3, myMap2) {
		fmt.Println("Maps are equal")
	}

	for key, value := range myMap3 {
		fmt.Println("Key:", key, "Value:", value)
	}

	var myMap4 map[string]string
	if myMap4 == nil {
		fmt.Println("myMap4 is nil")
	}

	val := myMap4["key"]
	fmt.Println("myMap4:", val) // nil string returns blank

	//myMap4["key"] = "Value"  // Will result in panic
	//fmt.Println("myMap4:", myMap4)
	myMap4 = make(map[string]string)
	myMap4["key"] = "value"
	myMap4["key2"] = "value2"
	fmt.Println(myMap4)

	fmt.Println("myMap length is ", len(myMap))

	myMap5 := make(map[string]map[string]string)
	myMap5["map1"] = myMap4

	fmt.Println(myMap5)
}
