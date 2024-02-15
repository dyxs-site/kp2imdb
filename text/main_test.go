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
		// Minimal.
		{
			Actual:   "(2006) 92 мин.",
			Expected: "2006",
		},
		{
			Actual:   "(2006)",
			Expected: "2006",
		},
		{
			Actual:   "(2006-2010) 92 мин.",
			Expected: "2006",
		},
		{
			Actual:   "(2006-...) 92 мин.",
			Expected: "2006",
		},
		// Symbols.
		{
			Actual:   "G.I. Joe: FF FF FF (2006) 92 мин.",
			Expected: "G.I. Joe: FF FF FF 2006",
		},
		{
			Actual:   "G.I. Joe: FF FF FF (2006)",
			Expected: "G.I. Joe: FF FF FF 2006",
		},
		{
			Actual:   "G.I. Joe: FF FF FF (2006-2010) 92 мин.",
			Expected: "G.I. Joe: FF FF FF 2006",
		},
		{
			Actual:   "G.I. Joe: FF FF FF (2006-...) 92 мин.",
			Expected: "G.I. Joe: FF FF FF 2006",
		},
		// Confusing.
		{
			Actual:   "2012 (2009) 92 мин.",
			Expected: "2012 2009",
		},
		{
			Actual:   "2012 (2009)",
			Expected: "2012 2009",
		},
		{
			Actual:   "2012 (2009-2012) 92 мин.",
			Expected: "2012 2009",
		},
		{
			Actual:   "2012 (2009-...) 92 мин.",
			Expected: "2012 2009",
		},
		// #1.
		{
			Actual:   "Birdman or (EEE E EEEE E) (2006) 92 мин.",
			Expected: "Birdman or (EEE E EEEE E) 2006",
		},
		{
			Actual:   "Birdman or (EEE E EEEE E) (2006)",
			Expected: "Birdman or (EEE E EEEE E) 2006",
		},
		{
			Actual:   "Tra(sgre)dire (2006-2010) 92 мин.",
			Expected: "Tra(sgre)dire 2006",
		},
		{
			Actual:   "Tra(sgre)dire (2006-...) 92 мин.",
			Expected: "Tra(sgre)dire 2006",
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

// func TestFindSimilar(t *testing.T) {
// 	type Case struct {
// 		Target      string
// 		ExpectedIdx int
// 		Candidates  []string
// 	}
// 	cases := []Case{
// 		{
// 			Target:      "}{отт@бь)ч",
// 			ExpectedIdx: 1,
// 			Candidates: []string{
// 				"Мисс Поттер",
// 				"Хоттабыч",
// 				"Паутина шарлотты",
// 				"Ключ от тайной комнаты",
// 				"Скотт Уокер: Человек ХХХ столетия",
// 			},
// 		},
// 		{
// 			Target:      "Выкрутасы",
// 			ExpectedIdx: 0,
// 			Candidates: []string{
// 				"Vykrutasy",
// 				"Yalan Dolan",
// 				"Twists and Turns (1987)",
// 				"Potap & Nastya: Vikrutasi",
// 			},
// 		},
// 	}
// 	for _, casy := range cases {
// 		bestIdx := FindSimilar(casy.Target, casy.Candidates)
// 		if bestIdx != casy.ExpectedIdx {
// 			t.Fatalf("Expected idx: %d | Got: %d", casy.ExpectedIdx, bestIdx)
// 		}
// 	}
// }
