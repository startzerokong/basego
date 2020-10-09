package define

const (
	CodeSignWrong = "9999998"
	CodeIpLimit = "9999999"
	CodeUnknown = "10000000"
)

var CodeText = map[string]string{
	CodeSignWrong:      "签名错误",
	CodeIpLimit:        "ip超限",

}

func Message(code string) string {
	str, ok := CodeText[code]
	if ok {
		return str
	}
	return CodeText[CodeUnknown]
}
