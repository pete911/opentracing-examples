package httputil

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteJSONResponse(w http.ResponseWriter, v interface{}) {

	data, err := json.Marshal(v)
	if err != nil {
		log.Printf("json marshal: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if _, err := w.Write(data); err != nil {
		log.Printf("write response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
