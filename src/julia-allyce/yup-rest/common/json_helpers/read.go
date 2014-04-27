package json_helpers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Read will parses the JSON-encoded data in the http request and store the result in v
func Read(r *http.Request, v interface{}) bool {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Printf("ReadJson couldn't read request body %v", err)
		return false
	}

	if err = json.Unmarshal(body, v); err != nil {
		log.Printf("ReadJson couldn't parse request; body: %q err: %q", body, err)
		return false
	}

	return true
}
