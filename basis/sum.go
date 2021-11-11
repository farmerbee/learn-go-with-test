package main

// accumulate the numbers of each slice
func Sum(slices ...[]int) []int {
	res := make([]int, len(slices))
	for index, slice := range slices {
		var total int
		for _, num := range slice {
			total += num
		}
		res[index] = total
	}

	return res
}

// accumulate tails of each slice
func SumTails(slices ...[]int) []int {
	res := make([]int, len(slices))

	for index, slice := range slices {
		if len(slice) == 0 {
			res[index] = 0
		} else {
			res[index] = Sum(slice[1:])[0]
		}
	}

	return res
}
