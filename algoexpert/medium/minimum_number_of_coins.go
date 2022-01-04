package medium

import "math"

func MinimumNumberOfCoins(n int, denoms []int) int {
	numCoins := make([]int, n + 1)
	for i := 1; i < len(numCoins); i++ {
		numCoins[i] = math.MaxInt32
	}
	for _, denom := range denoms {
		for amount := range numCoins {
			if denom <= amount {
				numCoins[amount] = minInt(numCoins[amount], 1 + numCoins[amount - denom])
			}
		}
	}
	if numCoins[n] == math.MaxInt32 {
		return -1
	}
	return numCoins[n]
}