package router

import (
	"github.com/NominalTrajectory/nt-precision-rest-api/handlers/auth"
	"github.com/NominalTrajectory/nt-precision-rest-api/handlers/home"
	"github.com/NominalTrajectory/nt-precision-rest-api/handlers/okr"
	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter()

	/* ROUTES AND HANDLERS */
	r.HandleFunc("/", home.Home)
	r.HandleFunc("/objectives", okr.GetAllObjectives)

	r.HandleFunc("/register", auth.Register).Methods("POST")
	r.HandleFunc("/login", auth.Login).Methods("POST")
	r.HandleFunc("/logout", auth.Logout).Methods("POST")
	r.HandleFunc("/user", auth.User).Methods("GET")

	/*                     */

	return r
}
