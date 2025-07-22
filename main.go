package main

import (
	"encoding/json"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/charles-d-burton/hobor/encoder"
	"github.com/fxamacker/cbor/v2"
)

const (
	coreAPI = `http://supervisor/core`
)

func httpClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns: 5,
		},
		Timeout: 30 * time.Second,
	}
	return client
}

func main() {
	slog.Info("starting")
	envars := os.Environ()
	for _, envar := range envars {
		slog.Info("ENV", "variable", envar)
	}

	f, err := os.ReadFile("/data/options.json")
	if err != nil {
		slog.Error("unable to read /data/options.json file")
		os.Exit(1)
	}
	slog.Info("options", "file", string(f))

	slog.Info("setting up http client")
	client := httpClient()

	mux := http.NewServeMux()
	mux.HandleFunc("GET "+encoder.API, func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implement the logic to get from th ehome assistant API
		req, err := http.NewRequest(http.MethodGet, coreAPI+encoder.API, nil)
		if err != nil {
			slog.Error("unable to createe http request", "error", err)
		}
		// req.Header = r.Header
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer "+os.Getenv("SUPERVISOR_TOKEN"))
		slog.Info("headers", "header", req.Header)
		res, err := client.Do(req)
		if err != nil {
			// TOOD: better error handling
			slog.Error("unable to process request", "error", err)
			return
		}
		defer res.Body.Close()
		data, err := io.ReadAll(res.Body)
		if err != nil {
			slog.Error("error reading response body", "error", err)
			return
		}

		if res.StatusCode != http.StatusOK {
			http.Error(w, string(data), res.StatusCode)
			return
		}
		slog.Info("message body", "body", string(data))
		var hapiResponse encoder.GetAPIMessage
		err = json.Unmarshal(data, &hapiResponse)
		if err != nil {
			slog.Error("unable to unmarshal response to API", "error", err)
		}
		encoded, err := cbor.Marshal(hapiResponse)
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
