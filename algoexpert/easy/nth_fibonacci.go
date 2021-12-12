package easy

func GetNthFibonacci(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	return GetNthFibonacci(n-1) + GetNthFibonacci(n-2)
}
