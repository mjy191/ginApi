package enum

var Status = map[int]string{
	1: "未付款",
	2: "已付款",
	3: "已发货",
	4: "退款",
}

// 错误码
const CodeSuccess = 1
const CodeParamError = 2
const CodeSignError = 3
const CodeTokenError = 401
const CodeSystemError = 500

// 错误信息
var ErrMsg = map[int]string{
	CodeSuccess:     "success",
	CodeParamError:  "fail",
	CodeSignError:   "sign error",
	CodeTokenError:  "token error",
	CodeSystemError: "system error",
}
