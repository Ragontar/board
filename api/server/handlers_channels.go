package server

import (
	"api/testdata"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// TODO NOT IMPLEMENTED
func ChannelListUserIdGet(w http.ResponseWriter, r *http.Request) {
	log.Printf("[ChannelListUserIdGet] user-id: %s", mux.Vars(r)["user-id"])

	body, err := json.Marshal(testdata.RESPONSES.ChannelListUserIdGetResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(body)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
