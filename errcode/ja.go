package errcode

func InitJa() map[int]string {
	return map[int]string{
		URI_NOT_FOUND:          "請求する方法は存在しない",
		SYSTEM_ERROR:           "システム異常の場合は顧客サービスに連絡してください",
		DB_CONNECT_ERROR:       "データベース接続エラー",
		REQUEST_PARAMETER_LOST: "要求パラメータは空ではいけない",
		DATA_NOT_FOUND:         "無数の根拠",
		USER_NOT_FOUND:         "このユーザーなし",
		PASSWORD_ERROR:         "パスワードの誤り",
		VERIFY_CODE_ERROR:      "captchaエラー",
		TOKEN_ERROR:            "無効トークン",
		AUTHENTICATION_FAILURE: "セキュリティ認証の失敗",
		SERVICE_UNAVAILABLE:    "サービスが存在しない",
	}
}
