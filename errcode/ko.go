package errcode

func InitKo() map[int]string {
	return map[int]string{
		URI_NOT_FOUND:          "요청한 메서드가 존재하지 않습니다",
		SYSTEM_ERROR:           "시스템 이상이 있으니 고객 서비스 담당자에게 연락해 주십시오",
		DB_CONNECT_ERROR:       "데이터베이스 연결 오류",
		REQUEST_PARAMETER_LOST: "요청 인자가 비어 있으면 안 됩니다",
		DATA_NOT_FOUND:         "무수한 증거",
		USER_NOT_FOUND:         "이 사용자 없음",
		PASSWORD_ERROR:         "암호가 잘못되었습니다",
		VERIFY_CODE_ERROR:      "인증 코드 오류",
		TOKEN_ERROR:            "잘못된 토큰",
		AUTHENTICATION_FAILURE: "보안 인증 실패",
		SERVICE_UNAVAILABLE:    "서비스가 존재하지 않음",
	}
}
