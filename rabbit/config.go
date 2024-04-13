package rabbit

import "time"

type Config struct {
	Host              string
	Exchange          string
	Queue             string
	Vhost             string
	Prefetch          int
	Heartbeat         time.Duration
	ReconnectDuration time.Duration
}
