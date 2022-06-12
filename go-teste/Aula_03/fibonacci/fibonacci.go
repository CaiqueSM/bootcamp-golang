package fibonacci

func Fibonacci(n int)[]int64{
	terms := make([]int64, n+1)
	terms[0] = 0
	if n == 0 {
		return terms
	}
	terms[1] = 1
	if n == 1{
		return terms
	}

	for i := 2; i <= n; i++{
		terms[i] = terms[i - 1] + terms[i - 2]
	}
	return terms
}