package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(r *http.Request, result any) {
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(result)
}

func WriteToResponseBody(w http.ResponseWriter, response any) {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}
