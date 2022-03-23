package errcode

func InitEnUs() map[int]string {
	return map[int]string{
		URI_NOT_FOUND:          "Requested method does not exist",
		SYSTEM_ERROR:           "System error. Please contact customer service",
		DB_CONNECT_ERROR:       "Database connection error",
		REQUEST_PARAMETER_LOST: "Request parameter is empty",
		DATA_NOT_FOUND:         "Data not found",
		USER_NOT_FOUND:         "User not found",
		PASSWORD_ERROR:         "Invalid password",
		VERIFY_CODE_ERROR:      "Invalid code error",
		TOKEN_ERROR:            "Invalid token",
		AUTHENTICATION_FAILURE: "Authentication failure",
		SERVICE_UNAVAILABLE:    "Service unavailable",
	}
}
