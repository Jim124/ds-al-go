package main

import (
	"ds-algorithms/reverse"
	"ds-algorithms/stack"
	"fmt"
)

func main() {
	// arr := make([]string, 100000)
	// for i := 0; i < len(arr); i++ {
	// 	arr[i] = "memo"
	// }
	// findMemo(arr)
	// getFirstElement(arr)
	// funChallenge(arr)
	var intVal int64
	fmt.Println(intVal)
	arr1 := []string{"a", "b", "c", "d", "e"}
	arr2 := []string{"c", "r", "t", "r"}
	isCommon := getCommonItems(arr1, arr2)
	fmt.Println(isCommon)
	array := stack.New()
	array.Push("a")
	array.Push("b")
	array.Push("c")
	array.Push("d")
	item, error := array.Pop()
	if error != nil {
		fmt.Print(error.Error())
	}
	fmt.Println(item)
	deleteItem, _ := array.Delete(2)

	fmt.Println(deleteItem)
	fmt.Println(array)
	fmt.Println(array.Len())
	fmt.Println(array.Get(0))
	str := reverse.ReverseStr("Hi my name is Jim")
	fmt.Println(str)
	str2 := reverse.ReverseStr2("Hi my name is Jim")
	fmt.Println(str2)

}

// big o O(n)
func findMemo(arr []string) {
	for _, value := range arr {
		if value == "memo" {
			fmt.Println("this is memo")
		}
	}

}

// big o O(1)
func getFirstElement(arr []string) {
	fmt.Println(arr[0])
}

// big O(3 + 4n)
func funChallenge(input []string) int64 {
	var a int64 = 10                  // O(1)
	a = 50 + 3                        // O(1)
	for i := 0; i < len(input); i++ { // O(n)
		// anotherFunc() O(n)
		a++ // o(n)
		// strange := true O(n)
	}
	return a // o(1)
}

// O(a*b)
func getCommonItems(arr1, arr2 []string) bool {
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if arr1[i] == arr2[j] {
				return true
			}
		}
	}
	return false
}

// O(a + b)
func containsCommonItem2(arr1, arr2 []string) {
	// loop through first array and create object where properties === items in the array
	// loop through second array and check if item in second array exists on created object.
}
