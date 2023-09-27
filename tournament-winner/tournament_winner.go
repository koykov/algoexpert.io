package tournament_winner

const HOME_TEAM_WON = 1

func TournamentWinner(competitions [][]string, results []int) string {
	m := make(map[string]int, len(results))
	for i := 0; i < len(competitions); i++ {
		ht, at := competitions[i][0], competitions[i][1]
		if results[i] == HOME_TEAM_WON {
			m[ht]++
		} else {
			m[at]++
		}
	}
	mc, mx := "", -1
	for c, x := range m {
		if x > mx {
			mc, mx = c, x
		}
	}
	return mc
}
