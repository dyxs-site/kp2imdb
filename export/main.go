package export

import "os"

func SaveIds(path string, ids []string) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, v := range ids {
		if _, err = f.WriteString(v + "\n"); err != nil {
			return err
		}
	}
	return err
}
