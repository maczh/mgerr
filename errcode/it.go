package errcode

func InitIt() map[int]string {
	return map[int]string{
		URI_NOT_FOUND:          "Il metodo della richiesta non esiste",
		SYSTEM_ERROR:           "Il sistema è anomalo. Si prega di contattare il servizio clienti",
		DB_CONNECT_ERROR:       "Errore di connessione",
		REQUEST_PARAMETER_LOST: "I parametri di richiesta non sono liberi",
		DATA_NOT_FOUND:         "Dati non disponibili",
		USER_NOT_FOUND:         "Nessun utente",
		PASSWORD_ERROR:         "Errore di password",
		VERIFY_CODE_ERROR:      "Errori del captatore",
		TOKEN_ERROR:            "Invalidità totale",
		AUTHENTICATION_FAILURE: "Fallimento della certificazione di sicurezza",
		SERVICE_UNAVAILABLE:    "Servizi non disponibili",
	}
}
