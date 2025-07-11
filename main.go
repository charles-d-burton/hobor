package main

import (
	"net/http"

	"github.com/charles-d-burton/homebor/hapi"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET "+hapi.API, func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implement the logic to get from th ehome assistant API
		getAPI := hapi.GetAPIMessage{Message: "API running."}
	})
}
