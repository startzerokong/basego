package define

const (
	CodeTokenIsExpire = "9999994"
	CodeTokenIsWrong = "9999995"
	CodeDefaultIsNull = "9999996"
	CodeParamError = "9999997"
	CodeSignWrong = "9999998"
	CodeIpLimit = "9999999"
	CodeUnknown = "10000000"
)

var CodeText = map[string]string{
	CodeTokenIsExpire:  "Token已过期",
	CodeTokenIsWrong:   "Token错误",
	CodeDefaultIsNull:  "默认参数为空",
	CodeParamError:     "参数错误",
	CodeSignWrong:      "签名错误",
	CodeIpLimit:        "ip超限",
	CodeUnknown:        "未知错误",
}

func Message(code string) string {
	str, ok := CodeText[code]
	if ok {
		return str
	}
	return CodeText[CodeUnknown]
}
