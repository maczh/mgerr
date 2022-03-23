package errcode

func InitZhCN() map[int]string {
	return map[int]string{
		URI_NOT_FOUND:          "请求的方法不存在",
		SYSTEM_ERROR:           "系统异常，请联系客服",
		DB_CONNECT_ERROR:       "数据库连接错误",
		REQUEST_PARAMETER_LOST: "请求参数不可为空",
		DATA_NOT_FOUND:         "无数据",
		USER_NOT_FOUND:         "无此用户",
		PASSWORD_ERROR:         "密码错误",
		VERIFY_CODE_ERROR:      "验证码错误",
		TOKEN_ERROR:            "无效token",
		AUTHENTICATION_FAILURE: "安全认证失败",
		SERVICE_UNAVAILABLE:    "服务不存在",
	}
}
