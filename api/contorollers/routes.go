package contorollers

import (
	"net/http"

	"github.com/tahahmmcgl/kullanici_api/api/middlewares"
)

type Route struct {
	Uri          string
	Method       string
	Handler      func(w http.ResponseWriter, r *http.Request)
	AuthRequired bool
}

//rotaları yüklemek için kullanılır
func (s *Server) initializeRoutes() {
	for _, route := range Load(s) {
		if route.AuthRequired {
			s.Router.HandleFunc(route.Uri,
				middlewares.SetMiddlewareLogger(
					middlewares.SetMilddlewareJson(
						middlewares.SetMiddlewareAuthentication(route.Handler))),
			).Methods(route.Method)
		} else {
			s.Router.HandleFunc(route.Uri,
				middlewares.SetMiddlewareLogger(
					middlewares.SetMilddlewareJson(route.Handler)),
			).Methods(route.Method)
		}
	}
}

// rotaları yüklemek için kullanılır
func Load(s *Server) []Route {
	routes := HomeRoutes(s)
	routes = append(routes, LoginRoutes(s)...)
	routes = append(routes, UserRoutes(s)...)
	routes = append(routes, TaskRoutes(s)...)
	return routes
}
func HomeRoutes(s *Server) []Route {
	routes := []Route{
		Route{
			Uri:          "/",
			Method:       http.MethodGet,
			Handler:      s.Home,
			AuthRequired: false,
		},
	}
	return routes
}
func LoginRoutes(s *Server) []Route {
	routes := []Route{
		Route{
			Uri:          "/login",
			Method:       http.MethodPost,
			Handler:      s.Login,
			AuthRequired: false,
		},
	}
	return routes
}
func UserRoutes(s *Server) []Route {
	{

	}
	routes := []Route{
		Route{
			Uri:          "/users",
			Method:       http.MethodGet,
			Handler:      s.GetUsers,
			AuthRequired: true,
		},
		Route{
			Uri:          "/users",
			Method:       http.MethodPost,
			Handler:      s.CreateUser,
			AuthRequired: true,
		},
		Route{
			Uri:          "/users/{id}",
			Method:       http.MethodGet,
			Handler:      s.GetUser,
			AuthRequired: true,
		},
		Route{
			Uri:          "/users/{id}",
			Method:       http.MethodPut,
			Handler:      s.UpdateUser,
			AuthRequired: true,
		},
		Route{
			Uri:          "/users/{id}",
			Method:       http.MethodDelete,
			Handler:      s.DeleteUser,
			AuthRequired: true,
		},
	}
	return routes
}
func TaskRoutes(s *Server) []Route {
	routes := []Route{
		Route{
			Uri:          "/tasks",
			Method:       http.MethodPost,
			Handler:      s.CreateTask,
			AuthRequired: false,
		},
		Route{
			Uri:          "/tasks/{id}",
			Method:       http.MethodGet,
			Handler:      s.GetTask,
			AuthRequired: false,
		},
		Route{
			Uri:          "/tasks/{id}",
			Method:       http.MethodPut,
			Handler:      s.UpdateTask,
			AuthRequired: true,
		},
		Route{
			Uri:          "/tasks/{id}",
			Method:       http.MethodDelete,
			Handler:      s.DeleteTask,
			AuthRequired: true,
		},
		Route{
			Uri:          "/tasks/user/{id}",
			Method:       http.MethodGet,
			Handler:      s.GetUserTasks,
			AuthRequired: false,
		},
	}
	return routes
}
