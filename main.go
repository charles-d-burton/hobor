package main

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/charles-d-burton/hobor/hapi"
	"github.com/fxamacker/cbor/v2"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET "+hapi.API, func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implement the logic to get from th ehome assistant API
		getAPI := hapi.GetAPIMessage{Message: "API running."}
		// payload, err := io.ReadAll(r.Body)
		// if err != nil {
		// 	slog.Error("unable to decode messge body", "error", err)
		// }
		encoded, err := cbor.Marshal(getAPI)
		if err != nil {
			slog.Error("unable to marshal cbor getAPI", "error", err)
			return
		}
		w.Header().Set("Content-Type", "application/cbor")

		_, err = w.Write(encoded)
		if err != nil {
			slog.Error("unable to write encoded message", "error", err)
		}
	})
	log.Fatal(http.ListenAndServe(":9007", mux))
}
