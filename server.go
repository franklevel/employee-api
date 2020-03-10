package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type (
	employee struct {
		ID       int    `json: "id"`
		Name     string `json: "name"`
		Position string `json: "position"`
	}
)

var (
	employees = map[int]*employee{}
	seq       = 1
)

// ---------
// Handlers
// ---------

func createEmployee(c echo.Context) error {
	y := &employee{
		ID: seq,
	}

	if err := c.Bind(y); err != nil {
		return err
	}

	employees[y.ID] = y
	seq++
	return c.JSON(http.StatusCreated, y)
}

func getEmployee(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, employees[id])
}

func updateEmployee(c echo.Context) error {
	y := new(employee)
	if err := c.Bind(y); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	employees[id].Name = y.Name
	return c.JSON(http.StatusOK, employees[id])
}

func deleteEmployee(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(employees, id)
	return c.NoContent(http.StatusNoContent)
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/employees", createEmployee)
	e.GET("/employees/:id", getEmployee)
	e.PUT("/employees/:id", updateEmployee)
	e.DELETE("/employees", deleteEmployee)

	// Serve and listen
	e.Logger.Fatal(e.Start(":1323"))
}
