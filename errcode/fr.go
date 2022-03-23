package errcode

func InitFr() map[int]string {
	return map[int]string{
		URI_NOT_FOUND:          "La méthode demandée n’existe pas",
		SYSTEM_ERROR:           "Anomalie du système, veuillez contacter le service client",
		DB_CONNECT_ERROR:       "Connexion incorrecte à la base de données",
		REQUEST_PARAMETER_LOST: "Les paramètres de requête ne peuvent pas être vides",
		DATA_NOT_FOUND:         "Données non disponibles",
		USER_NOT_FOUND:         "Aucun utilisateur pour cet utilisateur",
		PASSWORD_ERROR:         "Mot de passe incorrect",
		VERIFY_CODE_ERROR:      "Code de vérification incorrect",
		TOKEN_ERROR:            "Token non valide",
		AUTHENTICATION_FAILURE: "Échec de l’authentification de sécurité",
		SERVICE_UNAVAILABLE:    "Le service n’existe pas",
	}
}
