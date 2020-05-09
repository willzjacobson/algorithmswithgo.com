package module01

// GCD : using the Euclidean algorithm
func GCD(a, b int) int {
	// Step 1: if b == 0, return a
	for b != 0 {
		// Step 2: a = b, b = a%b
		a, b = b, a%b
	}

	return a
}

// GCDRecursive : using the Euclidean algorithm (recursively)
func GCDRecursive(a, b int) int {
	if b == 0 {
		return a
	}
	return GCDRecursive(b, a%b)
}

/*
a = 252
b = 105

a = 105
b = 252%105 = 42

a = 42
b = 105%42 = 21

a = 21
b = 42%21 = 0

=> 21
*/
