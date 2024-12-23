package main

func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	subArr := arr[len(arr)/2]
	l := 0
	r := len(arr) - 1
	for l <= r {
		for arr[l] < subArr {
			l++
		}

		for arr[r] > subArr {
			r--
		}

		if l <= r {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}

	quickSort(arr[l:]) // Cортировка правой половины

	quickSort(arr[:r+1]) // Cортировка левой половины
}
