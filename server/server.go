package server

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/NominalTrajectory/nt-precision-rest-api/handlers/auth"
	"github.com/NominalTrajectory/nt-precision-rest-api/handlers/home"
	"github.com/NominalTrajectory/nt-precision-rest-api/handlers/okr"
	"github.com/rs/cors"

	"github.com/gorilla/mux"
)

var Server *http.Server

func InitializeServer(listenAddress string, dbConnectionString string) {

	router := setupRouter()

	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences: []tls.CurveID{
			tls.CurveP256,
			tls.X25519,
		},

		MinVersion: tls.VersionTLS12,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
	}

	// TODO: Move cors setting to the config file and import as env vars
	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},
		AllowedHeaders: []string{
			"*",
		},
		AllowCredentials: true,
	})

	handler := corsOptions.Handler(router)

	server := &http.Server{
		Addr:         listenAddress,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		TLSConfig:    tlsConfig,
		Handler:      handler,
	}

	Server = server
}

func setupRouter() *mux.Router {
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
