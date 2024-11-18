package handlers

import (
	"api-gateway/shared/utils"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

var farmsServiceUrl = userServiceURL + "/farms"

func ListFarmsHandler(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest(http.MethodGet, farmsServiceUrl, nil)
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

func CreateFarmHandler(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest(http.MethodPost, farmsServiceUrl, r.Body)
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

func GetFarmByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	farmId := vars["id"]

	url := farmsServiceUrl + "/" + farmId

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

func ListFarmsByFarmerIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	farmId := vars["id"]

	url := farmsServiceUrl + "/farmer/" + farmId

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
