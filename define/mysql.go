package define

type Db struct {
	Hostname string `mapstructure:"hostname,omitempty"`
	User     string `mapstructure:"user,omitempty"`
	Password string `mapstructure:"password,omitempty"`
	Dbname   string `mapstructure:"dbname,omitempty"`
}