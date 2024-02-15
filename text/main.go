package text

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type CleanedTitle struct {
	Title string
	Year  int
}

func CleanTitle(title string) (*CleanedTitle, error) {
	title = strings.TrimSpace(title)

	/*
		Possible titles:
		(2006) 92 мин.
		(2006)
		(2006-2010) 92 мин.
		(2006-...) 92 мин.
		+ same but with title on front.
	*/
	if len(title) == 0 {
		return nil, errWrongTitle(title)
	}

	titleS := strings.Split(title, " ")

	yearInfoCutIdx := len(titleS) - 1

	// Have minutes?
	if strings.HasSuffix(title, ".") {
		yearInfoCutIdx = len(titleS) - 3
	}

	titleName := strings.Join(titleS[:yearInfoCutIdx], " ")
	yearInfo := titleS[yearInfoCutIdx:]

	res, err := parseYear(yearInfo[0])
	if err != nil {
		return nil, fmt.Errorf("%w, %w", errWrongTitle(title), err)
	}

	return &CleanedTitle{
		Title: titleName,
		Year:  res,
	}, err
}

// (2006) 92 мин.
func parseYear(year string) (int, error) {
	// 2006)
	yearS := strings.TrimPrefix(year, "(")
	// 2006
	yearS = strings.TrimSuffix(yearS, ")")
	// [2006, 2010?]
	yearSP := strings.Split(yearS, "-")
	if len(yearSP) == 0 {
		return 0, errors.New("parseYear: empty yearSP")
	}
	return strconv.Atoi(yearSP[0])
}

func errWrongTitle(title string) error {
	return fmt.Errorf("wrong title format: %s", title)
}

// func FindSimilar(target string, where []string) int {

// 	target = Normalize(target)

// 	bestCf := 0.0
// 	bestIdx := 0

// 	for i, v := range where {
// 		v = Normalize(v)
// 		curCf := 0.0
// 		curCf += strutil.Similarity(target, v, metrics.NewHamming())
// 		curCf += strutil.Similarity(target, v, metrics.NewJaccard())
// 		curCf += strutil.Similarity(target, v, metrics.NewJaroWinkler())
// 		if curCf > bestCf {
// 			bestCf = curCf
// 			bestIdx = i
// 		}
// 	}

// 	return bestIdx
// }

// Trim space -> to ASCII -> to upper.
// func Normalize(str string) string {
// 	return strings.TrimSpace(strings.ToUpper(slug.Make(str)))
// }
