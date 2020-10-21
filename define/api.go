package define

type Api struct {
	BaseUri string `mapstructure:"baseuri,omitempty"`
	ConnectTimeout int `mapstructure:"connecttimeout,omitempty"`
	RequestTimeout int `mapstructure:"requesttimeout,omitempty"`
	RequestRetryCount int `mapstructure:"requestretrycount,omitempty"`
	DataKey string `mapstructure:"datakey,omitempty"`
	CodeKey string `mapstructure:"codekey,omitempty"`
	MessageKey string `mapstructure:"messagekey,omitempty"`
	SuccessCode *int `mapstructure:"successcode,omitempty"`
	Uri string `mapstructure:"uri,omitempty"`
	Method string `mapstructure:"method,omitempty"`
	Parameter string `mapstructure:"parameter,omitempty"`
	Required string `mapstructure:"required,omitempty"`
}

type ApiResponse struct {
	HttpCode int
	HttpCost int64

	RealResponse interface{}

	Code string
	Message string

	Data string

	ErrorType string
}
