package fib

type Color struct {
	R   int
	G   int
	B   int
	HSV string
}

func New(r, g, b int) *Color {
	return &Color{R: r, G: g, B: b}
}

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

func (f *Fibonacci) Red() Color {
	return Color{R: 255, G: 0, B: 0}
}

func IsPrime(n int) bool {
	return true
}
