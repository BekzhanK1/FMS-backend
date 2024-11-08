package routes

import (
	"api-gateway/internal/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func Routes(router *mux.Router) {
	router.HandleFunc("/health-check", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods(http.MethodGet)

	router = router.PathPrefix("/api/v1").Subrouter()

	// usersRouter := router.PathPrefix("/users").Subrouter()
	// usersRouter.HandleFunc("", handlers.GetUsersHandler).Methods(http.MethodGet)
	// usersRouter.HandleFunc("", handlers.CreateUserHandler).Methods(http.MethodPost)
	// usersRouter.HandleFunc("/search", handlers.GetUserByQueryHandler).Methods(http.MethodGet)
	// usersRouter.HandleFunc("/{id}", handlers.GetUserByIDHandler).Methods(http.MethodGet)
	// usersRouter.HandleFunc("/{id}", handlers.UpdateUserHandler).Methods(http.MethodPut)
	// usersRouter.HandleFunc("/{id}", handlers.DeleteUserHandler).Methods(http.MethodDelete)


	productsRouter := router.PathPrefix("/documents").Subrouter()
	productsRouter.HandleFunc("", handlers.GetDocumentsHandler).Methods(http.MethodGet)
	productsRouter.HandleFunc("", handlers.CreateDocumentHandler).Methods(http.MethodPost)
	productsRouter.HandleFunc("/{id}", handlers.GetDocumentHandler).Methods(http.MethodGet)

}