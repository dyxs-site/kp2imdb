package imdb

import (
	"testing"

	"github.com/oklookat/kp2imdb/text"
)

func TestSearchTitle(t *testing.T) {
	tit, err := SearchTitle(text.CleanedTitle{
		Title: "Fast and the Furious: Tokyo Drift, The",
		Year:  2006,
	}, true)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%v", tit)
}
