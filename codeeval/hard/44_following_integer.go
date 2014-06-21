/*
You are writing out a list of numbers.
Your list contains all numbers with exactly Di digits in its
decimal representation which are equal to i, for each i between 1 and 9, inclusive.
You are writing them out in ascending order.
For example, you might be writing every number with two '1's and one '5'.
Your list would begin 115, 151, 511, 1015, 1051.
Given N, the last number you wrote, compute what the next number in the list will be.
The number of 1s, 2s, ..., 9s is fixed but the number of 0s is arbitrary.
*/

package main

import "fmt"

func lexiPermute(nums []byte, perms int) int {
	for p := 0; p < perms-1; p++ {
		var pivot int
		end := len(nums) - 1
		// Find pivot, cell in which the cell before it is higher in value.
		for i := end; i > 0; i-- {
			// In 143 the 1 is less than 4. The pivot is set to the 1 position.
			if nums[i-1] < nums[i] {
				pivot = i - 1
				break
			}
		}
		// Find first number before pivot which is higher than it.
		for i := end; i > pivot; i-- {
			if nums[i] > nums[pivot] {
				nums[i], nums[pivot] = nums[pivot], nums[i]
				pivot++
				break
			}
		}
		for i := end; pivot < i; {
			nums[i], nums[pivot] = nums[pivot], nums[i]
			pivot++
			i--
		}
	}
	permutation := 0
	for _, v := range nums {
		permutation = (permutation * 10) + int(v)
	}
	return permutation
}
func main() {
	fmt.Println(lexiPermute([]byte{1, 1, 5}, 4))
}
