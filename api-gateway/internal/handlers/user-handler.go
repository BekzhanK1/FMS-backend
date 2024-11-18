package handlers

import (
	"api-gateway/internal/config"
	"io"
	"net/http"

	"api-gateway/shared/utils"
	"github.com/gorilla/mux"
)

var userServiceURL = config.Envs.Users_url + "/users"

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	url := userServiceURL + "/" + userID

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

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	url := userServiceURL + "/login"

	req, err := http.NewRequest(http.MethodPost, url, r.Body)
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

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	url := userServiceURL + "/register"

	req, err := http.NewRequest(http.MethodPost, url, r.Body)
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

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	userServiceURL = userServiceURL + "/" + userID

	req, err := http.NewRequest(http.MethodPut, userServiceURL, r.Body)
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

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	userServiceURL = userServiceURL + "/" + userID

	req, err := http.NewRequest(http.MethodDelete, userServiceURL, nil)
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

func ActivateUserHandler(w http.ResponseWriter, r *http.Request) {
	url := userServiceURL + "/activate"

	req, err := http.NewRequest(http.MethodPost, url, r.Body)
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
