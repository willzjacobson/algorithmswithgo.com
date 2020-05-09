package module01

// FibonacciRecursive returns the nth fibonacci number.
//
// A Fibonacci number N is defined as:
//
//   Fibonacci(N) = Fibonacci(N-1) + Fibonacci(N-2)
//
// Where the following base cases are used:
//
//   Fibonacci(0) => 0
//   Fibonacci(1) => 1
//
//
// Examples:
//
//   Fibonacci(0) => 0
//   Fibonacci(1) => 1
//   Fibonacci(2) => 1
//   Fibonacci(3) => 2
//   Fibonacci(4) => 3
//   Fibonacci(5) => 5
//   Fibonacci(6) => 8
//   Fibonacci(7) => 13
//   Fibonacci(14) => 377
//
func FibonacciRecursive(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

// Fibonacci returns the nth fibonacci number.
func Fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	low := 1  // fib(1)
	high := 1 // fib(2)

	for i := 2; i <= n; i++ {
		oldHigh := high
		high = low + oldHigh
		low = oldHigh
	}

	return low
}
