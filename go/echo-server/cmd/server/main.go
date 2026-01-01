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


	e.GET("/students/:id", database.GetStudentHandler)
    e.GET("/all_class_schedule", database.GetAllScheduleHandler)
    e.GET("/schedule/group/:id", database.GetScheduleByGroupHandler)
	  e.Start(":8080")
}