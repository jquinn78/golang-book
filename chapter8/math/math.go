package math

//Finds the average of a series of numbers
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}

	return total / float64(len(xs))
}

//Finds the max number in a series of numbers
func Max(xs []float64) float64 {

	maxNum := xs[0]
	for _, x := range xs {
		if x > maxNum {
			maxNum = x
		}
	}

	return maxNum
}

//finds the min number in a series of numbers
func Min(xs []float64) float64 {

	minNum := xs[0]
	for _, x := range xs {
		if x < minNum {
			minNum = x
		}
	}

	return minNum
}
