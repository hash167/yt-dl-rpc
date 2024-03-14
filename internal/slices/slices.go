package slices

// Array.prototype.reduce() inspired function.
// Reduces executes a "reducer" function on each element of the slice
// passing in the result of the previous execution.
// The final result is a single value.
//
//	a := []int{3, 5, 7, 9}
//	sum := slices.Map(a, func(c int, n int) int {
//		return c + n
//	})
//	//sum: 24
func Reduce[T, U any](s []T, f func(c U, n T) U, initValue U) U {
	acc := initValue
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}

// Partition takes a generic slice and divides it into n sub arrays.
//
//	s := []int{1, 2, 4, 5, 16}
//	p := slices.Partition(s, 2)
//	// p: {[2, 4], [8, 10], [32]}
func Partition[T any](arr []T, chunkSize int) (temp [][]T) {
	temp = [][]T{}
	for i := 0; i < len(arr); i += chunkSize {
		if i == len(arr)-1 {
			temp = append(temp, arr[i:])
			return
		}
		if i >= len(arr) {
			return
		}
		temp = append(temp, arr[i:i+chunkSize])
	}
	return
}

// Array.prototype.map() inspired function.
// Map takes a slices of type T and a certain value and
// transform each value with a mapper function.
// The result will be a slice of the same type.
//
//	a := []int{1, 2, 4, 5, 16}
//	mapped := slices.Map(a, func(e int) int {
//		return e * 2
//	})
//	//mapped: {2, 4, 8, 10, 32}
func Map[T any, M any](a []T, f func(x T) M) []M {
	n := make([]M, len(a))
	for i, e := range a {
		n[i] = f(e)
	}
	return n
}

// Array.prototype.includes() inspired function.
// Includes takes a slices of type T and a certain value and
// determines whether the slice include the value among its entries.
//
//	a := []string{"ciao", "come", "va", "grande"}
//	c1 := slices.Includes(a, "va") // true
//	c2 := slices.Includes(a, "vaa") // false
func Includes[T comparable](elems []T, v T) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}

// Array.prototype.filter() inspired function.
// Filter takes a slices of type T and a certain value and
// filters it down to the elements that pass the test function.
//
//	a := []string{"ciao", "come", "va", "grande"}
//	filtered := slices.Filter(a, func(e string) bool {
//		return strings.HasPrefix(e, "c")
//	})
//	//filtered: {"ciao", "come"}
func Filter[T any](slice []T, f func(e T) bool) []T {
	var n []T
	for _, e := range slice {
		if f(e) {
			n = append(n, e)
		}
	}
	return n
}

// Array.prototype.concat inspired function
// Concat takes a slice of slices and concat them without
// using append() function.
func Concat[T any](s [][]T) []T {
	var (
		i        int
		totalLen int
	)

	for _, s := range s {
		totalLen += len(s)
	}

	res := make([]T, totalLen)
	for _, s := range s {
		i += copy(res[i:], s)
	}

	return res
}
