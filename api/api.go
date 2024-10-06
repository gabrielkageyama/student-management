package api

import (
	"fmt"
	"net/http"

	"github.com/gabrielkageyama/api_teste1/db"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type API struct {
	Echo *echo.Echo
	DB   *gorm.DB
}

func NewServer() *API {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db := db.Init()

	return &API{
		Echo: e,
		DB:   db,
	}
}

func (api *API) Start() error {
	return api.Echo.Start(":8080")
}

func (api *API) ConfigureRoutes() {
	// Routes
	api.Echo.GET("/students", getStudents)
	api.Echo.POST("/students", createStudent)
	api.Echo.GET("/students/:id", getStudentInfo)
	api.Echo.PUT("/students/:id", updateStudent)
	api.Echo.DELETE("/students/:id", deleteStudent)
}

// Handler
func getStudents(c echo.Context) error {

	students, err := db.GetStudents()
	if err != nil {
		return c.String(http.StatusNotFound, "Failed to get Students")
	}

	return c.JSON(http.StatusOK, students)
}

func createStudent(c echo.Context) error {
	student := db.Student{}

	if err := c.Bind(&student); err != nil {
		return err
	}

	if err := db.AddStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Error to create student")
	}

	return c.String(http.StatusOK, "Create student")
}

func getStudentInfo(c echo.Context) error {
	id := c.Param("id")
	getStud := fmt.Sprintf("Get %s student", id)
	return c.String(http.StatusOK, getStud)
}

func updateStudent(c echo.Context) error {
	id := c.Param("id")
	updateStud := fmt.Sprintf("Update %s student", id)
	return c.String(http.StatusOK, updateStud)
}

func deleteStudent(c echo.Context) error {
	id := c.Param("id")
	deleteStud := fmt.Sprintf("Delete %s student", id)
	return c.String(http.StatusOK, deleteStud)
}
