package code

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeError
	CodeInvalidParams

	ErrorAuthCheckTokenFail ResCode = 20001 + iota
	ErrorAuthCheckTokenTimeout
	ErrorAuthToken
	ErrorAuth

	ErrorAccountExist ResCode = 30001 + iota
	ErrorAccountNotExist
	ErrorPwdNotCompare
)

// code和msg做一个映射
var codeMsgMap = map[ResCode]string{
	CodeSuccess:       "success",
	CodeError:         "fail",
	CodeInvalidParams: "请求参数错误",
}

// ResMsg 传入code返回对应的msg
func (c ResCode) ResMsg() string {
	m, ok := codeMsgMap[c]
	if !ok {
		return codeMsgMap[CodeError] // 如果传入的code出错，直接返回500
	}

	return m
}
