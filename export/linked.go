package export

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

// [KP ID]IMDB ID.
type LinkedData map[string]string

func LoadLinks(path string) (LinkedData, error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	res := LinkedData{}
	err = json.NewDecoder(f).Decode(&res)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return res, nil
		}
	}
	return res, err
}

func SaveLinks(path string, links LinkedData) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(links)
}
