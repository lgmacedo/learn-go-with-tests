package arrays_and_slices

func Sum(numbers [5]int) (sum int) {
	sum = 0
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumSlice(numbers []int) (sum int) {
	sum = 0
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(slices ...[]int) (result []int) {
	for _, slice := range slices {
		sliceSum := 0
		for _, item := range slice {
			sliceSum += item
		}
		result = append(result, sliceSum)
	}
	return
}

func SumAllOption2(slices ...[]int) (result []int) {
	for _, slice := range slices {
		result = append(result, SumSlice(slice))
	}
	return
}

func SumAllTails(slices ...[]int) (result []int) {
	for _, slice := range slices {
		if len(slice) == 0 { // se fizermos [1:] num slice de tamanho 0, vai dar problema de slice bounds out of range
			result = append(result, 0)
		} else {
			result = append(result, SumSlice(slice[1:]))
		}
	}
	return
}
