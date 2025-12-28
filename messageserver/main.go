package messageserver

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/listeners"
)

func startServer() error {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		done <- true
	}()

	server := mqtt.New(&mqtt.Options{
		InlineClient: true,
	})
	l := listeners.NewTCP(listeners.Config{ID: "cbor", Address: ":2883"})
	err := server.AddListener(l)
	if err != nil {
		return err
	}

	// start the werver
	go func() {
		err := server.Serve()
		if err != nil {
			log.Fatal(err)
		}
	}()

	return nil
}
