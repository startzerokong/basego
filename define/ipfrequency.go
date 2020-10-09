package define

import "time"

type Ip struct {
	Expire time.Duration `mapstructure:"expire,omitempty"`
	Limit int64 `mapstructure:"limit,omitempty"`
}
