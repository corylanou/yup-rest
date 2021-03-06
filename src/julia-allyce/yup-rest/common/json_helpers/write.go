package json_helpers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func Write(w http.ResponseWriter, code int, v interface{}) {
	// avoid json vulnerabilities, always wrap v in an object literal
	// http://haacked.com/archive/2008/11/20/anatomy-of-a-subtle-json-vulnerability.aspx/
	doc := map[string]interface{}{"d": v}

	data, err := json.Marshal(doc)

	if err != nil {
		log.Printf("Error marshalling json: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
