package tournament_winner

import (
	"strconv"
	"testing"
)

var stages = []struct {
	competitions [][]string
	results      []int
	expect       string
}{
	{
		competitions: [][]string{
			{"HTML", "C#"},
			{"C#", "Python"},
			{"Python", "HTML"},
		},
		results: []int{0, 0, 1},
		expect:  "Python",
	},
	{
		competitions: [][]string{
			{"HTML", "Java"},
			{"Java", "Python"},
			{"Python", "HTML"},
		},
		results: []int{0, 1, 1},
		expect:  "Java",
	},
	{
		competitions: [][]string{
			{"HTML", "Java"},
			{"Java", "Python"},
			{"Python", "HTML"},
			{"C#", "Python"},
			{"Java", "C#"},
			{"C#", "HTML"},
		},
		results: []int{0, 1, 1, 1, 0, 1},
		expect:  "C#",
	},
	{
		competitions: [][]string{
			{"HTML", "Java"},
			{"Java", "Python"},
			{"Python", "HTML"},
			{"C#", "Python"},
			{"Java", "C#"},
			{"C#", "HTML"},
			{"SQL", "C#"},
			{"HTML", "SQL"},
			{"SQL", "Python"},
			{"SQL", "Java"},
		},
		results: []int{0, 1, 1, 1, 0, 1, 0, 1, 1, 0},
		expect:  "C#",
	},
	{
		competitions: [][]string{
			{"Bulls", "Eagles"},
		},
		results: []int{1},
		expect:  "Bulls",
	},
	{
		competitions: [][]string{
			{"Bulls", "Eagles"},
			{"Bulls", "Bears"},
			{"Bears", "Eagles"},
		},
		results: []int{0, 0, 0},
		expect:  "Eagles",
	},
	{
		competitions: [][]string{
			{"Bulls", "Eagles"},
			{"Bulls", "Bears"},
			{"Bulls", "Monkeys"},
			{"Eagles", "Bears"},
			{"Eagles", "Monkeys"},
			{"Bears", "Monkeys"},
		},
		results: []int{1, 1, 1, 1, 1, 1},
		expect:  "Bulls",
	},
	{
		competitions: [][]string{
			{"AlgoMasters", "FrontPage Freebirds"},
			{"Runtime Terror", "Static Startup"},
			{"WeC#", "Hypertext Assassins"},
			{"AlgoMasters", "WeC#"},
			{"Static Startup", "Hypertext Assassins"},
			{"Runtime Terror", "FrontPage Freebirds"},
			{"AlgoMasters", "Runtime Terror"},
			{"Hypertext Assassins", "FrontPage Freebirds"},
			{"Static Startup", "WeC#"},
			{"AlgoMasters", "Static Startup"},
			{"FrontPage Freebirds", "WeC#"},
			{"Hypertext Assassins", "Runtime Terror"},
			{"AlgoMasters", "Hypertext Assassins"},
			{"WeC#", "Runtime Terror"},
			{"FrontPage Freebirds", "Static Startup"},
		},
		results: []int{1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0},
		expect:  "AlgoMasters",
	},
	{
		competitions: [][]string{
			{"HTML", "Java"},
			{"Java", "Python"},
			{"Python", "HTML"},
			{"C#", "Python"},
			{"Java", "C#"},
			{"C#", "HTML"},
			{"SQL", "C#"},
			{"HTML", "SQL"},
			{"SQL", "Python"},
			{"SQL", "Java"},
		},
		results: []int{0, 0, 0, 0, 0, 0, 1, 0, 1, 1},
		expect:  "SQL",
	}, {
		competitions: [][]string{
			{"A", "B"},
		},
		results: []int{0},
		expect:  "B",
	},
}

func TestTournamentWinner(t *testing.T) {
	for i := 0; i < len(stages); i++ {
		stage := &stages[i]
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := TournamentWinner(stage.competitions, stage.results)
			if r != stage.expect {
				t.FailNow()
			}
		})
	}
}
