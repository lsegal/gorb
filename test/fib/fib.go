package fib

type Fibonacci struct {
}

func (f *Fibonacci) Fib(n int) int {
	return f.fib(n)
}

func (f *Fibonacci) fib(n int) int {
	if n < 2 {
		return n
	}
	return f.fib(n-1) + f.fib(n-2)
}

func IsPrime(n int) bool {
	return true
}
