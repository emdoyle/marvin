package main

import (
	"fmt"
	"net/http"
)

//EventHandler handles responding to subscribed events from Slack
func EventHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	// TODO: actually decode the JSON event payload
	fmt.Fprintf(w, "Event handled!")
}
