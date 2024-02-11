package export

import (
	"encoding/json"
	"os"

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
		data[i].FromAltName = true
		cTitle, err := text.CleanTitle(data[i].AltName)
		if err != nil {
			return nil, err
		}
		if len(cTitle.Title) == 0 {
			cTitle.Title = data[i].Name
			data[i].FromAltName = false
		}
		data[i].Parsed = *cTitle
	}
	return data, err
}

type KpExport struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	AltName   string `json:"alt_name"`
	DateAdded string `json:"date_added"`

	Parsed      text.CleanedTitle `json:"-"`
	FromAltName bool              `json:"-"`
}
