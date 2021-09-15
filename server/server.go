package server

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/NominalTrajectory/nt-precision-rest-api/handlers/home"
	"github.com/NominalTrajectory/nt-precision-rest-api/handlers/okr"

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

	server := &http.Server{
		Addr:         listenAddress,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		TLSConfig:    tlsConfig,
		Handler:      router,
	}

	Server = server
}

func setupRouter() *mux.Router {
	r := mux.NewRouter()

	/* ROUTES AND HANDLERS */
	r.HandleFunc("/", home.Home)
	r.HandleFunc("/objectives", okr.GetAllObjectives)

	/*                     */

	return r
}
