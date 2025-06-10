package util

// Permutations generates all permutations of arr.
// Again stolen from StackOverflow and commented generously here for understanding. I find these kinds of algorithms hard to understand right now.
// https://stackoverflow.com/questions/30226438/generate-all-permutations-in-go
// This is apparently Heap's Algorithm: https://en.wikipedia.org/wiki/Heap%27s_algorithm
func Permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
