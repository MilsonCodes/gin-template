package routes

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// Route - a struct mapping to route initializations
type Route struct {
	Type       string
	Path       string
	Controller func(c *gin.Context)
}

// Router - function to create the endpoints provided in routes
func Router(r *gin.Engine, routes []Route) (err error) {
	for _, route := range routes {
		switch route.Type {
		case "GET":
			r.GET(route.Path, route.Controller)
		case "POST":
			r.POST(route.Path, route.Controller)
		default:
			err = errors.New("route type '" + route.Type + "' is not yet supported")
			return
		}

	}
	return
}
