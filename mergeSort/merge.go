package mergeSort

import "slices"

func MergeSortedArray(arr1, arr2 []int) []int {
	if len(arr1) == 0 {
		return arr2
	}
	if len(arr2) == 0 {
		return arr1
	}
	var mergeArray = []int{}
	mergeArray = arr1[:]
	for i := 0; i < len(arr2); i++ {
		mergeArray = append(mergeArray, arr2[i])
	}
	slices.Sort(mergeArray)
	return mergeArray
}
