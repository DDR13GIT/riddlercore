package conn

import (
	"github.com/RichardKnop/machinery/v1"
	mCfg "github.com/RichardKnop/machinery/v1/config"

	"ddr13/riddlercore/internal/config"
)

var machineryServer *machinery.Server
var machineryWorker *machinery.Worker

func ConnectMachinery(cfg *config.MQCfg) error {
	machineryConf := &mCfg.Config{
		Broker:          cfg.Broker,
		ResultBackend:   cfg.ResultBackend,
		DefaultQueue:    cfg.DefaultQueue,
		ResultsExpireIn: cfg.ResultsExpireIn,
		AMQP: &mCfg.AMQPConfig{
			Exchange:      cfg.AMQP.Exchange,
			ExchangeType:  cfg.AMQP.ExchangeType,
			BindingKey:    cfg.AMQP.BindingKey,
			PrefetchCount: cfg.AMQP.PrefetchCount,
		},
	}

	// Create server instance
	server, err := machinery.NewServer(machineryConf)
	if err != nil {
		return err
	}
	machineryServer = server
	machineryWorker = server.NewWorker(cfg.Worker.Name, cfg.Worker.Count)
	return nil
}

func GetMachinery() *machinery.Server {
	return machineryServer
}

func GetMachineryWorker() *machinery.Worker {
	return machineryWorker
}
