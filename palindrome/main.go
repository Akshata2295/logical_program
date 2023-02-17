package main

func isPalindrome(x int) bool {
	num := x
	res := 0
	for x > 0 {
		remainder := x %10
		res = res + remainder
		x = x / 10
	}

	if num == res {
		return true
	} else {
		return false
	}
}

func main() {
	n := 123
	isPalindrome(n)
}