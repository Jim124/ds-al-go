package reverse

import (
	"slices"
	"strings"
)

func ReverseStr(str string) string {
	arr := strings.Split(str, "")
	len := len(arr) - 1
	reverseArray := []string{}
	for i := len; i >= 0; i-- {
		reverseArray = append(reverseArray, arr[i])
	}
	return strings.Join(reverseArray, "")
}

func ReverseStr2(str string) string {
	reverseArray := strings.Split(str, "")
	slices.Reverse(reverseArray)
	return strings.Join(reverseArray, "")
}
