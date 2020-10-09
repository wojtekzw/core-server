package io

import "github.com/nats-io/nats.go"

// NatsWriter - nats writer struct
type NatsWriter struct {
	Conn    *nats.Conn
	Subject string
}

// Write - writer inteface
func (w *NatsWriter) Write(b []byte) (int, error) {
	err := w.Conn.Publish(w.Subject, b)
	return len(b), err
}

// NewNatsWriter - creat new writer
func NewNatsWriter(server string, subject string) (*NatsWriter, error) {
	conn, err := nats.Connect(server)
	return &NatsWriter{Conn: conn, Subject: subject}, err
}
