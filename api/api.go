package api

import (
	"github.com/gabrielkageyama/api_teste1/db"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type API struct {
	Echo *echo.Echo
	DB   *db.StudentHandler
}

// @title Entities Register (students as example) API
// @version 1.0
// @description This is a sample server to registration of entities, students used as exemple.

// @host localhost:8080
// @BasePath /
// @schemes http

func NewServer() *API {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	database := db.Init()
	studentDB := db.NewStundentHandler(database)

	return &API{
		Echo: e,
		DB:   studentDB,
	}
}

func (api *API) Start() error {
	return api.Echo.Start(":8080")
}

func (api *API) ConfigureRoutes() {
	// Routes
	api.Echo.GET("/students", api.getStudents)
	api.Echo.POST("/students", api.createStudent)
	api.Echo.GET("/students/:id", api.getStudentInfo)
	api.Echo.PUT("/students/:id", api.updateStudent)
	api.Echo.DELETE("/students/:id", api.deleteStudent)
	api.Echo.GET("/swagger/*", echoSwagger.WrapHandler)
}
