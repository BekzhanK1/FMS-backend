package http

import (
	"net/http"
	"user-service/internal/middleware"
	applService "user-service/internal/service/application"
	authService "user-service/internal/service/auth"
	farmService "user-service/internal/service/farms"
	userService "user-service/internal/service/user"

	"github.com/gorilla/mux"
)

type Handler struct {
	userService userService.Service
	authService authService.Service
	farmService farmService.FarmService
	applService applService.ApplicationService
}

func NewHanlder(userService userService.Service, authService authService.Service, farmService farmService.FarmService, applService applService.ApplicationService) *Handler {
	return &Handler{
		userService,
		authService,
		farmService,
		applService,
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	protected := router.PathPrefix("").Subrouter()
	protected.Use(middleware.AuthMiddleware)
	protected.HandleFunc("/profile", h.ProfileHandler).Methods(http.MethodGet)
	protected.HandleFunc("/switch-role", h.SwitchUserRoleHandler).Methods(http.MethodPut)

	protected.HandleFunc("/farms", h.CreateFarmHandler).Methods(http.MethodPost)
	protected.HandleFunc("/farms", h.ListFarms).Methods(http.MethodGet)
	protected.HandleFunc("/farms/{id}", h.GetFarmByID).Methods(http.MethodGet)
	protected.HandleFunc("/farms/farmer/{id}", h.ListFarmsByFarmerID).Methods(http.MethodGet)

	protected.HandleFunc("/applications", h.ListApplications).Methods(http.MethodGet)
	protected.HandleFunc("/applications/{id}", h.GetApplicationByID).Methods(http.MethodGet)
	protected.HandleFunc("/applications/{id}", h.UpdateApplication).Methods(http.MethodPut)
	protected.HandleFunc("/applications/farmer/{id}", h.ListApplicationsByFarmerID).Methods(http.MethodGet)

	router.HandleFunc("/register", h.CreateUserHandler).Methods(http.MethodPost)
	router.HandleFunc("/login", h.LoginHandler).Methods(http.MethodPost)
	router.HandleFunc("/{id}", h.GetUserHandler).Methods(http.MethodGet)
	router.HandleFunc("/{id}", h.UpdateUserHandler).Methods(http.MethodPut)
	router.HandleFunc("/{id}", h.DeleteUserHandler).Methods(http.MethodDelete)
	router.HandleFunc("/activate", h.ActivateUserHandler).Methods(http.MethodPost)
}
