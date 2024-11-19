package handlers

import (
	"api-gateway/shared/utils"
	"io"
	"net/http"

	apiUtils "api-gateway/internal/utils"

	"github.com/gorilla/mux"
)


var applicationServiceUrl = userServiceURL + "/applications"

func ListAppsHandler(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest(http.MethodGet, applicationServiceUrl, nil)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	apiUtils.CopyHeaders(r, req)

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

func GetApplicationByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appID := vars["id"]

	url := applicationServiceUrl + "/" + appID

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	apiUtils.CopyHeaders(r, req)

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

func UpdateApplicationHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appID := vars["id"]

	url := applicationServiceUrl + "/" + appID

	req, err := http.NewRequest(http.MethodPut, url, r.Body)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	apiUtils.CopyHeaders(r, req)

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

func ListAppsByFarmerIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	appID := vars["id"]

	url := applicationServiceUrl + "/farmer/" + appID

	req, err := http.NewRequest(http.MethodGet, url, r.Body)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	apiUtils.CopyHeaders(r, req)

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
