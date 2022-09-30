package server

import (
	grpcclients "api/grpcClients"
	"api/models"
	"encoding/json"
	"io"
	"net/http"
)

// TODO NOT IMPLEMENTED
func LinkTelegramUserIdConfirmPut(w http.ResponseWriter, r *http.Request) {
	var tcc models.TelegramConfirmationCode

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if err := json.Unmarshal(body, &tcc); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	status, err := grpcclients.GetAccountServiceClient().LinkTelegramConfirmCode(tcc)

	if status.Ok && status.PasswordRequired {
		// TODO NOT IMPLEMENTED
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// TODO NOT IMPLEMENTED
func LinkTelegramUserIdPut(w http.ResponseWriter, r *http.Request) {
	ts := models.TelegramSession{}
	var userID string
	for _, cookie := range r.Cookies() {
		if cookie.Name == "user_id" {
			userID = cookie.Value
		}
	}

	tcc, err := grpcclients.GetAccountServiceClient().LinkTelegram(ts, userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	body, err := json.Marshal(tcc)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
