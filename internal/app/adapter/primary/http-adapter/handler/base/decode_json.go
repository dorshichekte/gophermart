package basehandler

import (
	"encoding/json"
	"net/http"
	"strings"

	customerror "gophermarket/internal/error"
)

func (h *BaseHandler) DecodeJSON(req *http.Request, data interface{}) error {
	ct := req.Header.Get("Content-Type")
	if ct != "" {
		mediaType := strings.ToLower(strings.TrimSpace(strings.Split(ct, ";")[0]))

		isMediaTypeApplicationJSON := mediaType == "application/json"
		if !isMediaTypeApplicationJSON {
			msg := errWrongHeaderContentType + ": expected application/json"
			return customerror.New(msg)
		}
	}

	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(&data)
	if err != nil {
		return h.handleJSONDecodeError(err)
	}

	return nil
}
