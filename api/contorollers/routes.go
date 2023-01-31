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
			AuthRequired: false,
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
