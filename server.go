package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

type Server struct {
	db *gorm.DB
}

func NewServer(db *gorm.DB) *Server {
	return &Server{db: db}
}

func (s *Server) RegisterRouter(router *httprouter.Router) {
	router.GET("/ping", s.ping)
}

func (s *Server) ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}
