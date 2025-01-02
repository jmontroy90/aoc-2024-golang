package util

// Important: I fully admit to stealing this Cartesian Product generator from here:
// https://stackoverflow.com/questions/23412146/how-to-create-cartesian-product
// As penance, I've tried my best to explain it here, and renamed some variables for clarity.
// I might've been able to come up with something like this eventually, but it's funky.
func CartesianProduct(elems []int, choose int) func() []int {
	// Putting these outside the closure and slicing means we're only using one backing array.
	// Plus, we only generate one permutation at a time - no massive materialized list. Good job, StackOverflow person.
	perms := make([]int, choose)
	ixes := make([]int, choose)
	return func() []int {
		perms = perms[:len(ixes)]
		// Assign our next permutation to the values specified by our index slice.
		for i, xi := range ixes {
			perms[i] = elems[xi]
		}
		// This loop essentially holds all but one index constant every time. It's easiest to see in action.
		// For elems == [7 8 9], choose == 4:
		// [0 0 0 0] -> selects [7 7 7 7]
		// [0 0 0 1] -> selects [7 7 7 8]
		// [0 0 0 2] -> selects [7 7 7 9]
		// [0 0 1 0] -> selects [7 7 8 7]
		// [0 0 1 1] -> selects [7 7 8 8]
		// [0 0 1 2] -> selects [7 7 8 9]
		// [0 0 2 0] -> selects [7 7 9 7]
		// ...and so on.
		for i := len(ixes) - 1; i >= 0; i-- {
			ixes[i]++
			if ixes[i] < len(elems) {
				break
			}
			ixes[i] = 0
			if i <= 0 {
				ixes = ixes[0:0]
				break
			}
		}
		return perms
	}
}
