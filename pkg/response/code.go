package response

type ResCode int64

const (
	CodeSuccess ResCode = 2000

	CodeInvalidParam = 4999 + iota
	CodeServerBusy
	CodeNeedLogin
	CodeGenericError
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess: "Success",

	CodeInvalidParam: "Request parameter error",
	CodeServerBusy:   "System busy, please try again later",
	CodeNeedLogin:    "Not logged in",
	CodeGenericError: "Error",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
