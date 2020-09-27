package define

type MemCache struct {
	Host string `mapstructure:"host,omitempty"`
	Port string `mapstructure:"port,omitempty"`
}