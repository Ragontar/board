package server

import (
	grpcclients "api/grpcClients"
	"encoding/json"
	"net/http"
)

func RegistrationPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	s, err := grpcclients.GetAccountServiceClient().RegisterUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respBody, err := json.Marshal(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	http.SetCookie(w, &http.Cookie{
		Name:  "Token",
		Value: *s.Token,
	})
	w.Write(respBody)
}

func AuthorizationPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	s, err := grpcclients.GetAccountServiceClient().AuthorizeUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respBody, err := json.Marshal(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	http.SetCookie(w, &http.Cookie{
		Name:  "Token",
		Value: *s.Token,
	})
	w.Write(respBody)
}
