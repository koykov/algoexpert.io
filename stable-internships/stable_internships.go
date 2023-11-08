package stable_internships

func StableInternships(interns [][]int, teams [][]int) (r [][]int) {
	intRest := make([]int, len(interns))
	for i := 0; i < len(interns); i++ {
		intRest[i] = i
	}
	t := make([]int, len(interns))

	tprefs := make(map[int]map[int]int, len(teams))
	for idx, team := range teams {
		tr := make(map[int]int, len(team))
		for rank, intern := range team {
			tr[intern] = rank
		}
		tprefs[idx] = tr
	}

	tplace := make(map[int]int, len(teams))

	for len(intRest) > 0 {
		intern := intRest[0]
		intRest = intRest[1:]
		tpref := interns[intern][t[intern]]
		t[intern]++

		if _, ok := tplace[tpref]; !ok {
			tplace[tpref] = intern
			continue
		}

		placedIntern := tplace[tpref]
		if tprefs[tpref][placedIntern] < tprefs[tpref][intern] {
			intRest = append(intRest, intern)
		} else {
			intRest = append(intRest, placedIntern)
			tplace[tpref] = intern
		}

	}

	r = make([][]int, 0, len(teams))
	for i := 0; i < len(teams); i++ {
		r = append(r, []int{tplace[i], i})
	}

	return
}
