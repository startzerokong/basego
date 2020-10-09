package define

type Ip struct {
	Expire int `mapstructure:"expire,omitempty"`
	Limit int64 `mapstructure:"limit,omitempty"`
}
