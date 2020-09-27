package define

type Redis struct {
	Host string `mapstructure:"host,omitempty"`
	Port string `mapstructure:"port,omitempty"`
	PassWord string `mapstructure:"password,omitempty"`
	Db int `mapstructure:"db,omitempty"`
	ConnectTimeout int `mapstructure:"connecttimeout,omitempty"`
	ReadTimeout int `mapstructure:"readtimeout,omitempty"`
	WriteTimeout int `mapstructure:"writetimeout,omitempty"`
}
