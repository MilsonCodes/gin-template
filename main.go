package main

import (
	"fmt"

	"github.com/MilsonCodes/gin-template/controllers"
	"github.com/MilsonCodes/gin-template/models"
	"github.com/MilsonCodes/gin-template/routes"
	"github.com/gin-gonic/gin"
)

var routeSlice = []routes.Route{
	{Type: "GET", Path: "/users", Controller: controllers.FindUsers},
	{Type: "POST", Path: "/users", Controller: controllers.CreateUser},
	{Type: "GET", Path: "/users/:id", Controller: controllers.FindUser},
	{Type: "PATCH", Path: "/users/:id", Controller: controllers.UpdateUser},
	{Type: "DELETE", Path: "/users/:id", Controller: controllers.DeleteUser},
}

func main() {
	r := gin.Default()

	models.ConnectDataBase()

	err := routes.Router(r, routeSlice)
	if err != nil {
		fmt.Println(err)
		panic("Error loading router!")
	}

	r.Run()
}
