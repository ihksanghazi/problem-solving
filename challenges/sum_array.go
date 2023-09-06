package challenges

func SumArray(array []int32) int32 {
	var sum int32

	for _, num := range array {
		sum += num
	}

	return sum
}
