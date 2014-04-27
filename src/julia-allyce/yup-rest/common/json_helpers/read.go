package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// ReadJson will parses the JSON-encoded data in the http request and store the result in v
func ReadJson(w http.ResponseWriter, r *http.Request, v interface{}) bool {
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		WriteError(w, "ReadJson couldn't read request body %v", err)
		return false
	}

	if err = json.Unmarshal(body, v); err != nil {
		WriteError(w, "ReadJson couldn't parse request; body: %q err: %q", body, err)
		return false
	}

	return true
}
