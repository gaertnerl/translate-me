package communication

import (
	"encoding/json"
	"net/http"
)

const (
	HTTP  string = "http"
	HTTPS string = "https"
)

func GetJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)
	dec.Decode(target)
	return nil
}
