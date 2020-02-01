package errcode

import "github.com/web_zktnotify/internal/app/model"

type Error int

func (e Error) Error() string {
	return errorMsg[int(e)]
}

const (
	// public status
	RET_SUCCESS           = 0
	RET_SYSTEM_ERROR      = -1
	RET_DECRYP_ERROR      = 1004
	RET_PARAM_ERROR       = 1005
	RET_PARAM_PARSE_ERROR = 1007
	RET_BUSINESS_ERROR    = 1008
	RET_RATE_ERROR        = 1014
	RET_PERMISSION_ERROR  = 1016
	RET_EMPTY_LIST_ERROR  = 6003
	RET_EMPTY_RESP_ERROR  = 6004
)

var (
	// public status
	ERR_SYSTEM_ERROR      = Error(RET_SYSTEM_ERROR)
	ERR_DECRYP_ERROR      = Error(RET_DECRYP_ERROR)
	ERR_PARAM_ERROR       = Error(RET_PARAM_ERROR)
	ERR_PARAM_PARSE_ERROR = Error(RET_PARAM_PARSE_ERROR)
	ERR_RATE_ERROR        = Error(RET_RATE_ERROR)
	ERR_EMPTY_LIST_ERROR  = Error(RET_EMPTY_LIST_ERROR)
	ERR_BUSINESS_ERROR    = Error(RET_BUSINESS_ERROR)
	ERR_EMPTY_RESP_ERROR  = Error(RET_EMPTY_RESP_ERROR)

	ERR_PERMISSION_ERROR = Error(RET_PERMISSION_ERROR)
)

var (
	errorMsg = map[int]string{
		// public status
		RET_SUCCESS:           "success",
		RET_SYSTEM_ERROR:      "系统繁忙",
		RET_DECRYP_ERROR:      "解密错误",
		RET_PARAM_ERROR:       "请求参数非法",
		RET_PARAM_PARSE_ERROR: "json或xml解析错误",
		RET_RATE_ERROR:        "请求过于频繁",
		RET_PERMISSION_ERROR:  "权限错误",
		RET_EMPTY_LIST_ERROR:  "空白的列表",
		RET_EMPTY_RESP_ERROR:  "响应体为空",
		RET_BUSINESS_ERROR:    "业务异常",
	}
)

func (e Error) Result() *model.Result {
	return &model.Result{
		Status: int(e),
		Msg:    e.Error(),
	}
}
