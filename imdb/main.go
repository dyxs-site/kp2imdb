package imdb

import (
	"fmt"
	"net/http"
	"time"

	"github.com/StalkR/imdb"
	"github.com/oklookat/kp2imdb/text"
)

var client = &http.Client{
	Transport: &customTransport{http.DefaultTransport},
}

// todo? use
// https://pkg.go.dev/github.com/grailbio/base/tsv#example-Reader
// https://developer.imdb.com/non-commercial-datasets/
// if title not found, then use api

// Returns imdb id.
func SearchTitle(cTitle text.CleanedTitle, bypassSimilarity bool) (*imdb.Title, error) {
	searchTxt := fmt.Sprintf("%s (%d)", cTitle.Title, cTitle.Year)

	titles, err := imdb.SearchTitle(client, searchTxt)
	if err != nil {
		return nil, err
	}

	if len(titles) == 0 {
		return nil, nil
	}

	// Find best.
	var titlesFiltered []imdb.Title
	for _, t2 := range titles {
		alphaYear := t2.Year - cTitle.Year
		// Can be 1 year difference.
		if alphaYear < -1 || alphaYear > 1 {
			continue
		}
		if len(titles) > 1 {
			alphaTitleLen := len(t2.Name) - len(cTitle.Title)
			if alphaTitleLen < -5 || alphaTitleLen > 5 {
				continue
			}
		}
		titlesFiltered = append(titlesFiltered, t2)
	}

	if len(titlesFiltered) == 0 {
		return nil, err
	}

	if bypassSimilarity {
		return &titlesFiltered[0], err
	}

	tIdx := make([]string, len(titlesFiltered))
	for i := range tIdx {
		tIdx[i] = titles[i].Name
	}
	bestIdx := text.FindSimilar(cTitle.Title, tIdx)

	return &titlesFiltered[bestIdx], err
}

// IMDb deployed awswaf and denies requests using the default Go user-agent (Go-http-client/1.1).
// For now it still allows requests from a browser user-agent. Remain respectful, no spam, etc.
const userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36"

type customTransport struct {
	http.RoundTripper
}

func (e *customTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	defer time.Sleep(time.Second)         // don't go too fast or risk being blocked by awswaf
	r.Header.Set("Accept-Language", "en") // avoid IP-based language detection
	r.Header.Set("User-Agent", userAgent)
	return e.RoundTripper.RoundTrip(r)
}
