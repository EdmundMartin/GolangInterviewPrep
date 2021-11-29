package easy


func IsPalindrome(target string) bool {
	backIdx := len(target) - 1

	for fwdIdx, _ := range target {
		if target[fwdIdx] != target[backIdx] {
			return false
		}
		if fwdIdx >= backIdx {
			return true
		}
		backIdx -= 1
	}
	return false
}
