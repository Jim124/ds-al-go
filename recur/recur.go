package recur

func FirstRecurringItem(arr []int) int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				return arr[i]
			}
		}
	}
	return 0
}

func FirstRecurringItem2(arr []int) int {
	obj := make(map[int]int, 0)
	for i := 0; i < len(arr); i++ {
		if obj[arr[i]] == arr[i] {
			return arr[i]
		} else {
			obj[arr[i]] = arr[i]
		}
	}
	return 0
}
