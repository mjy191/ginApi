package Enum

var Status = map[int]string{
	1: "未付款",
	2: "已付款",
	3: "已发货",
	4: "退款",
}

const CodeSuccess = 1

const CodeParamError = 2

const CodeSignError = 3

const CodeSystemError = 500

const CodeTokenError = 401
