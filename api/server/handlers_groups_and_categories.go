package server

import (
	"api/testdata"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// TODO NOT IMPLEMENTED
func GroupListUserIdCategoryCategoryIdPut(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GroupListUserIdCategoryCategoryIdPut] user-id: %s", mux.Vars(r)["user-id"])
	log.Printf("[GroupListUserIdGroupGroupIdPut] category-id: %s", mux.Vars(r)["category-id"])

	body, err := json.Marshal(testdata.RESPONSES.GroupListUserIdCategoryCategoryIdPutResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(body)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

// TODO NOT IMPLEMENTED
func GroupListUserIdCategoryPost(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GroupListUserIdCategoryPost] user-id: %s", mux.Vars(r)["user-id"])

	body, err := json.Marshal(testdata.RESPONSES.GroupListUserIdCategoryPostResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(body)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

// TODO NOT IMPLEMENTED
func GroupListUserIdGet(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GroupListUserIdGet] user-id: %s", mux.Vars(r)["user-id"])

	body, err := json.Marshal(testdata.RESPONSES.GroupListUserIdGetResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(body)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

// TODO NOT IMPLEMENTED
func GroupListUserIdGroupGroupIdPut(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GroupListUserIdGroupGroupIdPut] user-id: %s", mux.Vars(r)["user-id"])
	log.Printf("[GroupListUserIdGroupGroupIdPut] group-id: %s", mux.Vars(r)["group-id"])

	body, err := json.Marshal(testdata.RESPONSES.GroupListUserIdGroupGroupIdPutResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(body)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

// TODO NOT IMPLEMENTED
func GroupListUserIdGroupPost(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GroupListUserIdGroupPost] user-id: %s", mux.Vars(r)["user-id"])

	body, err := json.Marshal(testdata.RESPONSES.GroupListUserIdGroupPostResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(body)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
