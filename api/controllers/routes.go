package controllers

import "github.com/LucasStarlingdePaulaSalles/insprTasks/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

}
