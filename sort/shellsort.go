package sort

func ShellSort(arr []int) []int {
	interval := 0
	for interval < len(arr)/9 {
		interval = interval*3 + 1
	}
	for ; interval >= 1; interval /= 3 {
		for current := interval; current < len(arr); current++ {
			temp := arr[current]
			iter := current - interval
			for ; iter >= 0 && arr[iter] > temp; iter -= interval {
				arr[iter+interval] = arr[iter]
			}
			arr[iter+interval] = temp
		}
	}
	return arr
}
