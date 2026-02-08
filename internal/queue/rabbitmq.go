package queue

import (
	"encoding/json"
	"log/slog"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/tasks"
)

// RabbitMQ ...
type RabbitMQ struct {
	machineryServer *machinery.Server
}

// New ...
func New(machineryServer *machinery.Server) *RabbitMQ {
	return &RabbitMQ{machineryServer: machineryServer}
}

// Register ...
func (r *RabbitMQ) Register(taskName string, callBackFunc func(payload string) error) error {
	return nil
}

// Send ...
func (r *RabbitMQ) Send(taskName string, payload interface{}) error {
	b, _ := json.Marshal(payload)
	task := tasks.Signature{
		Name: taskName,
		Args: []tasks.Arg{{Type: "string", Value: string(b)}},
	}
	_, err := r.machineryServer.SendTask(&task)
	if err != nil {
		slog.Warn("failed to send task:", err)
	}
	return err
}

func (r *RabbitMQ) Retry(err error) error {
	slog.Warn("retry error:", err)
	return err
}
