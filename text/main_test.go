package text

import (
	"fmt"
	"strings"
	"testing"
)

func TestCleanTitle(t *testing.T) {
	type Case struct {
		Actual,
		Expected string
	}
	cases := []Case{
		{
			Actual:   "(2006) 92 мин.",
			Expected: "2006",
		},
		{
			Actual:   "Rampage (2009) 85 мин.",
			Expected: "Rampage 2009",
		},
		{
			Actual:   "Black Mirror (2011-...) 43 мин.",
			Expected: "Black Mirror 2011",
		},
		{
			Actual:   "Better Call Saul (2015-2022) 46 мин.",
			Expected: "Better Call Saul 2015",
		},
		{
			Actual:   "G.I. Joe: The Rise of Cobra (2009) 118 мин.",
			Expected: "G.I. Joe: The Rise of Cobra 2009",
		},
	}
	for _, casy := range cases {
		got, err := CleanTitle(casy.Actual)
		if err != nil {
			t.Error(err)
			t.FailNow()
			break
		}
		got2 := fmt.Sprintf("%s %d", got.Title, got.Year)
		got2 = strings.TrimSpace(got2)
		if got2 != casy.Expected {
			t.Fatalf("Actual: %s | Expected: %s | Got: %s", casy.Actual, casy.Expected, got2)
			break
		}
	}
}

func TestFindSimilar(t *testing.T) {
	type Case struct {
		Target      string
		ExpectedIdx int
		Candidates  []string
	}
	cases := []Case{
		{
			Target:      "}{отт@бь)ч",
			ExpectedIdx: 1,
			Candidates: []string{
				"Мисс Поттер",
				"Хоттабыч",
				"Паутина шарлотты",
				"Ключ от тайной комнаты",
				"Скотт Уокер: Человек ХХХ столетия",
			},
		},
		{
			Target:      "Выкрутасы",
			ExpectedIdx: 0,
			Candidates: []string{
				"Vykrutasy",
				"Yalan Dolan",
				"Twists and Turns (1987)",
				"Potap & Nastya: Vikrutasi",
			},
		},
	}
	for _, casy := range cases {
		bestIdx := FindSimilar(casy.Target, casy.Candidates)
		if bestIdx != casy.ExpectedIdx {
			t.Fatalf("Expected idx: %d | Got: %d", casy.ExpectedIdx, bestIdx)
		}
	}
}
