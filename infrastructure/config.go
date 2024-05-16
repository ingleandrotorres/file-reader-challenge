package infrastructure

import "time"

type Config struct {
}

func NewConfig() *Config {
	return &Config{}
}

func (c Config) TimeCacheGateways() func() time.Duration {
	return func() time.Duration {
		return time.Minute * 5
	}
}
