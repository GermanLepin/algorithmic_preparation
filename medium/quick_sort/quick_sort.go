package quick_sort

func quickSort(ar []int) {
	if len(ar) <= 1 {
		return
	}

	split := partition(ar)

	quickSort(ar[:split])
	quickSort(ar[split:])
}

func partition(data []int) int {
	pivot := data[len(data)/2]
	left := 0
	right := len(data) - 1

	for {
		for ; data[left] < pivot; left++ {
		}

		for ; data[right] > pivot; right-- {
		}

		if left >= right {
			return right
		}

		swap(data, left, right)
	}
}

func swap(ar []int, left, right int) {
	tmp := ar[left]
	ar[left] = ar[right]
	ar[right] = tmp
}
