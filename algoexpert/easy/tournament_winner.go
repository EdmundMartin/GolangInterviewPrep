package easy


func reverseResult(result int) int {
	if result == 1 {
		return 0
	}
	return 1
}

func getWinner(scores map[string]int) string {
	winner := ""
	topScore := -1
	for key, val := range scores {
		if val > topScore {
			winner = key
			topScore = val
		}
	}
	return winner
}

func TournamentWinner(competitions [][]string, results []int) string {
	scores := make(map[string]int)

	for idx := 0; idx < len(results); idx++ {
		winner := reverseResult(results[idx])
		winningTeam := competitions[idx][winner]
		_, ok := scores[winningTeam]
		if !ok {
			scores[winningTeam] = 3
		} else {
			scores[winningTeam] += 3
		}
	}
	return getWinner(scores)
}