package main

import (
	"fmt"
	"os"

	"github.com/oklookat/kp2imdb/cmd"
	"github.com/oklookat/kp2imdb/export"
	"github.com/oklookat/kp2imdb/imdb"
	"github.com/schollz/progressbar/v3"
)

const (
	_linksPath = "./links.json"
)

func main() {
	if len(os.Args) != 2 {
		println("Usage example: ./kp2imdb kinopoisk.json")
		println("More info at https://github.com/oklookat/kp2imdb")
		os.Exit(0)
		return
	}
	kpFile := os.Args[1]

	links, err := export.LoadLinks(_linksPath)
	chk(err)

	exported, err := export.LoadKinopoisk(kpFile)
	chk(err)

	stack := cmd.NewStack(50)
	bar := progressbar.New(len(exported))
	barUid := stack.AddAlwaysBottom("")
	setBar := func(i int) {
		bar.Add(1)
		stack.AlwaysBottom[barUid] = bar.String()
		stack.Render()
	}

	var imdbIds []string
	for i, ke := range exported {
		imdbId, ok := links[ke.ID]
		if ok {
			// already linked.
			imdbIds = append(imdbIds, imdbId)
			continue
		}

		searchMsg := func() string {
			return fmt.Sprintf("%s (%d)", ke.Parsed.Title, ke.Parsed.Year)
		}

		searchUid := stack.Add("🔍 " + searchMsg() + " 🔍")
		stack.Render()

		imdbTitle, err := imdb.SearchTitle(ke.Parsed, ke.FromAltName)
		chk(err)
		if imdbTitle == nil {
			stack.Stack[searchUid] = "❌ " + searchMsg() + " ❌"
			stack.Render()
			setBar(i)
			continue
		}

		setBar(i)

		stack.Stack[searchUid] = fmt.Sprintf("✅ %s | %v ✅", searchMsg(), imdbTitle)
		stack.Render()

		imdbIds = append(imdbIds, imdbTitle.ID)
		links[ke.ID] = imdbTitle.ID

		err = export.SaveLinks(_linksPath, links)
		chk(err)
	}

	err = export.SaveIds("ids.txt", imdbIds)
	chk(err)
}

func chk(err error) {
	if err != nil {
		println(err.Error())
		panic(err)
	}
}
