package config

import (
	"time"

	"github.com/spf13/viper"
)

var mq *MQCfg

// AMQP ...
type AMQP struct {
	Exchange      string
	ExchangeType  string
	BindingKey    string
	PrefetchCount int
}

// Worker ...
type Worker struct {
	Name  string
	Count int
}

// MQCfg ...
type MQCfg struct {
	Broker          string
	DefaultQueue    string
	ResultBackend   string
	ResultsExpireIn int
	TaskRetryDelay  time.Duration
	AMQP            AMQP
	Worker          Worker
}

// loadMQCfg ...
func loadMQCfg() {
	mq = &MQCfg{
		Broker:          viper.GetString("mq.broker"),
		DefaultQueue:    viper.GetString("mq.default_queue"),
		ResultBackend:   viper.GetString("mq.result_backend"),
		ResultsExpireIn: viper.GetInt("mq.results_expire_in"),
		TaskRetryDelay:  viper.GetDuration("mq.task_retry_delay") * time.Second,
		AMQP: AMQP{
			Exchange:      viper.GetString("mq.amqp.exchange"),
			ExchangeType:  viper.GetString("mq.amqp.exchange_type"),
			BindingKey:    viper.GetString("mq.amqp.binding_key"),
			PrefetchCount: viper.GetInt("mq.amqp.prefetch_count"),
		},
		Worker: Worker{
			Name:  viper.GetString("mq.worker.name"),
			Count: viper.GetInt("mq.worker.count"),
		},
	}
}

// MQ ...
func MQ() *MQCfg {
	return mq
}
