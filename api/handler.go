package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gabrielkageyama/api_teste1/schemas"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// getStudents godoc
//
// @Summary Get a list of entities
// @Description Retrieve entities details
// @Tags students
// @Accept json
// @produce json
// @Param register    path    int    false    "Registration"
// @Sucess 200 {object} schemas.StudentResponse
// @Failure 404
// @Router /students [get]
func (api *API) getStudents(c echo.Context) error {

	students, err := api.DB.GetStudents()
	if err != nil {
		return c.String(http.StatusNotFound, "Failed to get Students")
	}

	active := c.QueryParam("active")
	if active != "" {
		act, err := strconv.ParseBool(active)
		if err != nil {
			log.Error().Err(err).Msgf("[api] error to parse boolean")
			return c.String(http.StatusInternalServerError, "Failed to parse boolean")
		}

		students, err = api.DB.GetFilteredStudents(act)
	}

	listOfStudents := map[string][]schemas.StudentResponse{"students": schemas.NewResponse(students)}

	return c.JSON(http.StatusOK, listOfStudents)
}

// createStudent godoc
//
// @Summary Create entity
// @Description Create a new entity with all parameters of the student struct
// @Tags students
// @Accept json
// @produce json
// @Param register    path    int    false    "Registration"
// @Sucess 200 {object} schemas.StudentResponse
// @Failure 404
// @Router /students [post]
func (api *API) createStudent(c echo.Context) error {
	studentReq := StudentRequest{}
	if err := c.Bind(&studentReq); err != nil {
		return err
	}

	if err := studentReq.Validate(); err != nil {
		log.Error().Err(err).Msgf("[api] error validating struct")
		return c.String(http.StatusBadRequest, "Error validating student")
	}

	student := schemas.Student{
		Name:   studentReq.Name,
		Email:  studentReq.Email,
		CPF:    studentReq.CPF,
		Age:    studentReq.Age,
		Active: *studentReq.Active,
	}

	if err := api.DB.AddStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Error to create student")
	}

	return c.JSON(http.StatusOK, student)
}

// getStudentInfo godoc
//
// @Summary Get entity by ID
// @Description Retrive information about a specific entity
// @Tags students
// @Accept json
// @produce json
// @Param register    path    int    false    "Registration"
// @Sucess 200 {object} schemas.StudentResponse
// @Failure 404
// @Router /students/:id [get]
func (api *API) getStudentInfo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get Student ID")
	}

	student, err := api.DB.GetStudent(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNotFound, "Student not found")
	}
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get student")
	}

	return c.JSON(http.StatusOK, student)
}

// updateStudent godoc
//
// @Summary Update entity
// @Description Update new information about a specific entity
// @Tags students
// @Accept json
// @produce json
// @Param register    path    int    false    "Registration"
// @Sucess 200 {object} schemas.StudentResponse
// @Failure 404
// @Router /students/:id [put]
func (api *API) updateStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get Student ID")
	}

	requestedStudent := schemas.Student{}
	if err := c.Bind(&requestedStudent); err != nil {
		return err
	}

	updateStudent, err := api.DB.GetStudent(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNotFound, "Student not found")
	}
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get student")
	}

	student := updateStudentInfo(requestedStudent, updateStudent)

	if err := api.DB.UpdateStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to save student")
	}

	return c.JSON(http.StatusOK, student)
}

func updateStudentInfo(requestedStudent, updateStudent schemas.Student) schemas.Student {
	if requestedStudent.Name != "" {
		updateStudent.Name = requestedStudent.Name
	}
	if requestedStudent.Email != "" {
		updateStudent.Email = requestedStudent.Email
	}
	if requestedStudent.CPF > 0 {
		updateStudent.CPF = requestedStudent.CPF
	}
	if requestedStudent.Age > 0 {
		updateStudent.Age = requestedStudent.Age
	}
	if requestedStudent.Active != updateStudent.Active {
		updateStudent.Active = requestedStudent.Active
	}

	return updateStudent
}

// deleteStudent godoc
//
// @Summary Delete entity
// @Description Delete a specific entity
// @Tags students
// @Accept json
// @produce json
// @Param register    path    int    false    "Registration"
// @Sucess 200 {object} schemas.StudentResponse
// @Failure 404
// @Router /students/:id [delete]
func (api *API) deleteStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get Student ID")
	}

	student, err := api.DB.GetStudent(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNotFound, "Student not found")
	}
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get student")
	}

	if err := api.DB.DeleteStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to delete student")
	}

	return c.JSON(http.StatusOK, student)
}
