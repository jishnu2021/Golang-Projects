package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// ParseBody decodes JSON body into given struct
func ParseBody(r *http.Request, dst interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.Unmarshal(body, dst)
}
