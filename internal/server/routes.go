package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/markbates/goth/gothic"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", s.HelloWorldHandler)
	r.Get("/health", s.healthHandler)
	r.Get("/auth/{provider}/callback", s.GetAuthCallback)
	r.Get("/auth/{provider}", s.GetAuth)
	r.Get("/logout", s.Logout)

	return r
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}

func (s *Server) GetAuthCallback(w http.ResponseWriter, r *http.Request) {
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		log.Fatalf("error completing user auth. Err: %v", err)
	}

	jsonResp, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	fmt.Println(string(jsonResp))

	http.Redirect(w, r, "http://localhost:5173/", http.StatusTemporaryRedirect)
}

func (s *Server) GetAuth(w http.ResponseWriter, r *http.Request) {
	gothic.BeginAuthHandler(w, r)
}

func (s *Server) Logout(w http.ResponseWriter, r *http.Request) {
	gothic.Logout(w, r)
}
