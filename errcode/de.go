package errcode

func InitDe() map[int]string {
	return map[int]string{
		URI_NOT_FOUND:          "Die gesuchtmethode ist nicht möglich",
		SYSTEM_ERROR:           "Systemanomalie, bitte nehmen sie den kundendienst",
		DB_CONNECT_ERROR:       "Fehler in der datenbank",
		REQUEST_PARAMETER_LOST: "Bitte keine als parameter angeben",
		DATA_NOT_FOUND:         "Eine theorie gibt es nicht.",
		USER_NOT_FOUND:         "Kein anwender",
		PASSWORD_ERROR:         "Falsches passwort Julius cäsar",
		VERIFY_CODE_ERROR:      "Authentifizierungsfehler bestätigt",
		TOKEN_ERROR:            "Keine wirkung token",
		AUTHENTICATION_FAILURE: "Das sicherheitssystem ist fehlgeschlagen",
		SERVICE_UNAVAILABLE:    "Es gibt aber noch andere dinge",
	}
}
