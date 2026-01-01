package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Student struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Major string  `json:"major"`
	Year  int     `json:"year"`
	GPA   float32 `json:"gpa"`
}

func main() {
	e := echo.New()
	PORT := 8080

	students := map[string]Student{
		"202312045": {ID: "202312045", Name: "Alinur", Major: "CS", Year: 2, GPA: 3.4},
		"202312046": {ID: "202312046", Name: "Ali", Major: "ROBT", Year: 2, GPA: 3.9},
	}

	msg := fmt.Sprintf("Server is running on %d", PORT)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, msg)
	})

	e.GET("/students/:studentId", func(c echo.Context) error {
		id := c.Param("studentId")
		student, success := students[id]
		if !success {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "student was not found"})
		}
		return c.JSON(http.StatusOK, student)

	})
	e.Logger.Fatal(e.Start(":8080"))

}
