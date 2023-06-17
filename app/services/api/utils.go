package main

import (
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, data interface{}) {
	// set our json writer.
	out, _ := json.MarshalIndent(data, "", "  ")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(out)
}
