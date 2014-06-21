package main

import (
	"fmt"
)

var testCases []int64

func isPalindrome(num int64) bool {
	original := num
	var reversed int64
	for num > 0 {
		reversed = (reversed * 10) + (num % 10)
		num = num / 10
	}
	return original == reversed
}

func getInput() {
	var numTestCases int64
	fmt.Scanf("%d", &numTestCases)
	for ; numTestCases > 0; numTestCases-- {
		var testcase int64
		fmt.Scanf("%d", &testcase)
		testCases = append(testCases, testcase)
	}
}

func nextPalindrome(num int64) int64 {
	num++
	for {
		if isPalindrome(num) {
			break
		} else {
			num++
		}
	}
	return num
}

func findPalindromes() {
	for _, testCase := range testCases {
		fmt.Println(nextPalindrome(testCase))
	}
}

func main() {
	getInput()
	findPalindromes()
}
