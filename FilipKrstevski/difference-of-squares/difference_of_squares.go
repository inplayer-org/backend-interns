package diffsquares

func Difference(num int) int {
	return SquareOfSum(num) - SumOfSquares(num)
}

func SquareOfSum(num int) int {
	var sum int
	for i := 1; i < num+1; i++ {
		sum += i
	}
	return sum * sum
}
func SumOfSquares(num int) int {
	var sum int
	for i := 1; i < num+1; i++ {
		sum += i * i
	}
	return sum
}
