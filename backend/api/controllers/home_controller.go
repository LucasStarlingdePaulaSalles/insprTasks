package controllers

import (
	"net/http"

	"github.com/LucasStarlingdePaulaSalles/insprTasks/backend/api/responses"
)

// Home is a basic endpoint for testing api's connection
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To InsprTasks!")
}