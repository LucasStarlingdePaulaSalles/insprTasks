package controllers

import "github.com/LucasStarlingdePaulaSalles/insprTasks/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	//Task endpoints
	s.Router.HandleFunc("/task", middlewares.SetMiddlewareJSON(s.CreateTask)).Methods("POST")
	s.Router.HandleFunc("/tasks", middlewares.SetMiddlewareJSON(s.GetTasks)).Methods("GET")
}
