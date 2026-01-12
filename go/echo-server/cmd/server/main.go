package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/yungkhann/echo-server/internal/database"
)







func main() {
	godotenv.Load()

	db := database.InitDB()
	defer db.Close()
    e := echo.New()


	e.GET("/student/:id", database.GetStudentHandler)
    e.GET("/all_class_schedule", database.GetAllScheduleHandler)
    e.GET("/schedule/group/:id", database.GetScheduleByGroupHandler)
    e.POST("/attendance/subject", database.PostAttendanceHandler)
    e.GET("/attendanceByStudentId/:id", database.GetAttendanceByStudentIdHandler)
    e.GET("/attendanceBySubjectId/:id", database.GetAttendanceBySubjectIdHandler)
	  e.Start(":8080")
}