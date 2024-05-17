package dice_throws

func DiceThrows(numDice int, numSides int, target int) int {
	m := make([][]int, numDice+1)
	for i := 0; i < numDice+1; i++ {
		m[i] = make([]int, target+1)
	}
	m[0][0] = 1
	for i := 1; i < numDice+1; i++ {
		for j := 1; j < target+1; j++ {
			for k := 1; k < numSides+1; k++ {
				if j-k >= 0 {
					m[i][j] += m[i-1][j-k]
				}
			}
		}
	}
	return m[numDice][target]
}
