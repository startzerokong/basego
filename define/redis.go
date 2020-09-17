package define

type Redis struct {
	Host string
	Port string
	PassWord string
	Db int
	ConnectTimeout int
	ReadTimeout int
	WriteTimeout int
}
