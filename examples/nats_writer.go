// Example how to log at the same time to stdout and Nats server
// Part of the example from nats.io documentation
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/nats-io/nats.go"
	csio "github.com/wojtekzw/core-server/io"
)

func main() {

	nw, err := csio.NewNatsWriter("demo.nats.io", "SANDBOX.logs")

	if err != nil {
		log.Fatal(err)
	}
	defer nw.Conn.Close()

	wr := io.MultiWriter(nw, os.Stdout)
	log.SetOutput(wr)

	nc := nw.Conn

	go func() {
		for i := 1; i <= 5; i++ {
			log.Printf("hello %d", i)
		}
	}()

	if _, err := nc.Subscribe("SANDBOX.logs", func(m *nats.Msg) {
		fmt.Printf("NATS LOGS: %s", m.Data)
	}); err != nil {
		log.Fatal(err)
	}

	time.Sleep(5 * time.Second)
}
