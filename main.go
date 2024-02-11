package main

import (
	"fmt"
	"os"

	"github.com/oklookat/kp2imdb/export"
	"github.com/oklookat/kp2imdb/imdb"
	"github.com/schollz/progressbar/v3"
)

var _links export.LinkedData

const (
	_linksPath = "./links.json"
)

func main() {
	if len(os.Args) != 2 {
		println("Usage example: ./kp2imdb kpexportFile.json")

	}

	links, err := export.LoadLinks(_linksPath)
	chk(err)
	_links = links

	var imdbIds []string

	exported, err := export.LoadKinopoisk("kp_1.json")
	chk(err)

	bar := progressbar.New(len(exported))
	setBar := func(i int) {
		bar.Add(1)
		fmt.Printf("%s", bar.String())
	}

	for i, ke := range exported {
		println("SEARCH: " + ke.Parsed.Title)
		imdbId, ok := _links[ke.ID]
		if ok {
			println("ALREADY LINKED.")
			imdbIds = append(imdbIds, imdbId)
			setBar(i)
			continue
		}

		imdbTitle, err := imdb.SearchTitle(ke.Parsed, ke.FromAltName)
		chk(err)
		if imdbTitle == nil {
			fmt.Printf("NOT FOUND: %s (%d)\n", ke.Parsed.Title, ke.Parsed.Year)
			setBar(i)
			continue
		}

		setBar(i)

		fmt.Printf("BEST: %v\n", imdbTitle)

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
