package errcode

func InitZhTw() map[int]string {
	return map[int]string{
		URI_NOT_FOUND:          "請求的方法不存在",
		SYSTEM_ERROR:           "系統異常，請聯繫客服",
		DB_CONNECT_ERROR:       "資料庫連接錯誤",
		REQUEST_PARAMETER_LOST: "請求參數不可為空",
		DATA_NOT_FOUND:         "無數據",
		USER_NOT_FOUND:         "無此用戶",
		PASSWORD_ERROR:         "密碼錯誤",
		VERIFY_CODE_ERROR:      "驗證碼錯誤",
		TOKEN_ERROR:            "無效token",
		AUTHENTICATION_FAILURE: "安全認證失敗",
		SERVICE_UNAVAILABLE:    "服務不存在",
	}
}
