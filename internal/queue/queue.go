package queue

type Queue interface {
	Register(taskName string, callBackFunc func(payload string) error) error
	Send(taskName string, payload interface{}) error
	Retry(err error) error
}
