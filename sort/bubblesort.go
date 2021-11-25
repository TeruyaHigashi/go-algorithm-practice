package sort

func BubbleSort(arr []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := len(arr) - 1; i > 0; i-- {
			if arr[i] < arr[i-1] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
				swapped = true
			}
		}
	}
	return arr
}
