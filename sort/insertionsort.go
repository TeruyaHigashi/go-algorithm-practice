package sort

func InsertionSort(arr []int) []int {
	for curr := 1; curr < len(arr); curr++ {
		temp := arr[curr]
		iter := curr - 1
		for ; iter >= 0 && arr[iter] > temp; iter-- {
			arr[iter+1] = arr[iter]
		}
		arr[iter+1] = temp
	}
	return arr
}
