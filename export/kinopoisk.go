package export

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/oklookat/kp2imdb/text"
)

func LoadKinopoisk(path string) ([]KpExport, error) {
	var data []KpExport
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	if err = json.NewDecoder(f).Decode(&data); err != nil {
		return nil, err
	}
	for i := range data {
		altTitle, err := text.CleanTitle(data[i].AltName)
		if err != nil {
			return nil, err
		}
		if len(altTitle.Title) != 0 {
			data[i].ParsedAltName = altTitle
		}

		data[i].ParsedName = text.CleanedTitle{
			Title: strings.TrimSpace(data[i].Name),
			Year:  altTitle.Year,
		}
	}
	return data, err
}

type KpExport struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	AltName string `json:"alt_name"`

	ParsedName    text.CleanedTitle  `json:"-"`
	ParsedAltName *text.CleanedTitle `json:"-"`
}
