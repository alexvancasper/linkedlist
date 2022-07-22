
func bubleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	fmt.Printf("%+v\n", arr)
}

func sheykerSort(arr []int) {
	left := 0
	right := len(arr) - 1
	for left <= right {
		for i := right; i > left; i-- {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
			}
		}
		left++
		for i := left; i < right; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
		right--
	}
	fmt.Printf("%+v\n", arr)
}

func CombSort(arr []int) {
	var factor float64
	factor = 1.247
	var step float64
	step = float64(len(arr) - 1)
	for step >= 1 {
		for i := 0; float64(i)+step < float64(len(arr)); i++ {
			if arr[i] > arr[i+int(step)] {
				arr[i], arr[i+int(step)] = arr[i+int(step)], arr[i]
			}
		}
		step /= factor
	}
	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	fmt.Printf("%+v\n", arr)
}

func InsertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		x := arr[i]
		j := i
		for j > 0 && arr[j-1] > x {
			arr[j] = arr[j-1]
			j--
		}
		arr[j] = x
	}
	fmt.Printf("%+v\n", arr)
}

func SelectionSort(arr []int) {

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	fmt.Printf("%+v\n", arr)
}
