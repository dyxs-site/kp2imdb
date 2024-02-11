package text

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/adrg/strutil"
	"github.com/adrg/strutil/metrics"
	"github.com/gosimple/slug"
)

type CleanedTitle struct {
	Title string
	Year  int
}

func CleanTitle(title string) (*CleanedTitle, error) {
	title = strings.TrimSpace(title)

	pattern := `^(?:([^()]+) )?\((\d{4})(?:-\d{4})?.*?\).*$`
	regex := regexp.MustCompile(pattern)
	match := regex.FindStringSubmatch(title)

	if len(match) != 3 {
		return nil, fmt.Errorf("wrong title format: %s", title)
	}

	year, err := strconv.Atoi(match[2])
	if err != nil {
		return nil, err
	}
	res := &CleanedTitle{
		Title: match[1],
		Year:  year,
	}
	return res, err
}

func FindSimilar(target string, where []string) int {

	target = Normalize(target)

	bestCf := 0.0
	bestIdx := 0

	for i, v := range where {
		v = Normalize(v)
		curCf := 0.0
		curCf += strutil.Similarity(target, v, metrics.NewHamming())
		curCf += strutil.Similarity(target, v, metrics.NewJaccard())
		curCf += strutil.Similarity(target, v, metrics.NewJaroWinkler())
		if curCf > bestCf {
			bestCf = curCf
			bestIdx = i
		}
	}

	return bestIdx
}

// Trim space -> to ASCII -> to upper.
func Normalize(str string) string {
	return strings.TrimSpace(strings.ToUpper(slug.Make(str)))
}
