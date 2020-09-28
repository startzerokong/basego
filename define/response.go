package define

type Response struct {
	RequestId string
	Code int
	Message string
	Data interface{}
}
