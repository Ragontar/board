package server

import (
	"api/testdata"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// TODO NOT IMPLEMENTED
func MessagesGroupGroupIdGet(w http.ResponseWriter, r *http.Request) {
	log.Printf("[MessagesGroupGroupIdGet] group-id: %s", mux.Vars(r)["group-id"])

	body, err := json.Marshal(testdata.RESPONSES.MessagesGroupGroupIdGetResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(body)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
