package define

type User struct {
	Expire int `mapstructure:"expire,omitempty"`
	Limit int64 `mapstructure:"limit,omitempty"`
}
