package handlers

import (
	"api-gateway/internal/config"
	"io"
	"net/http"

	"api-gateway/shared/utils"

	"github.com/gorilla/mux"
)

var orderServiceURL = config.Envs.Users_url + "/users"

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID := vars["id"]

	url := orderServiceURL + "/" + orderID

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}
}
