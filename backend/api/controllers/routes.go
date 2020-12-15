package controllers

import "github.com/LucasStarlingdePaulaSalles/insprTasks/backend/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	//Task endpoints
	s.Router.HandleFunc("/task", middlewares.SetMiddlewareJSON(s.CreateTask)).Methods("POST")
	s.Router.HandleFunc("/tasks", middlewares.SetMiddlewareJSON(s.GetTasks)).Methods("GET")
	s.Router.HandleFunc("/work/{id}", middlewares.SetMiddlewareJSON(s.WorkOnATask)).Methods("PATCH")
	s.Router.HandleFunc("/stop", middlewares.SetMiddlewareJSON(s.StopWork)).Methods("PATCH")
	s.Router.HandleFunc("/status/{id}", middlewares.SetMiddlewareJSON(s.ChangeStatus)).Methods("PATCH")
	s.Router.HandleFunc("/tasks/date", middlewares.SetMiddlewareJSON(s.GetTasksByDate)).Methods("GET")
	s.Router.HandleFunc("/tasks/value", middlewares.SetMiddlewareJSON(s.GetTasksByValue)).Methods("GET")
}
