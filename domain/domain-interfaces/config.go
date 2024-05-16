package domain_interfaces

import "time"

type Config interface {
	TimeCacheGateways() func() time.Duration
}
