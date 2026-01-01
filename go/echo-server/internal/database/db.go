package database

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

type Student struct{
    FullName  string `json:"full_name"`
    Gender    string `json:"gender"`
    BirthDate string `json:"birth_date"`
    GroupID   int    `json:"group_id"`
    GroupName string `json:"group_name"`
}

type Group struct{
    ID        int    `json:"id"`
    GroupName string `json:"group_name"`
    FacultyID int    `json:"faculty_id"`
}
type Schedule struct{
    ID          int    `json:"id"`
    SubjectName string `json:"subject_name"`
    TimeSlot    string `json:"time_slot"`
    GroupID     int    `json:"group_id"`
    GroupName   string `json:"group_name"`
}

var pool *pgxpool.Pool

func InitDB() *pgxpool.Pool {
    connStr := os.Getenv("DATABASE_URL")
    if connStr == "" {
        log.Fatal("DATABASE_URL is not set in .env file")
    }

    config, err := pgxpool.ParseConfig(connStr)
    if err != nil {
        log.Fatalf("Unable to parse connection string: %v", err)
    }

    pool, err = pgxpool.NewWithConfig(context.Background(), config)
    if err != nil {
        log.Fatalf("Unable to create connection pool: %v", err)
    }

    err = pool.Ping(context.Background())
    if err != nil {
        log.Fatalf("Could not ping database: %v", err)
    }

    fmt.Println("Connected to PostgreSQL")
    return pool
}

func GetStudentHandler(c echo.Context) error{
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid student ID"})
    }

    student, err := getStudentById(pool, id)
    if err != nil {
        if err == pgx.ErrNoRows {
            return c.JSON(http.StatusNotFound, map[string]string{"error": "Student not found"})
        }
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database error"})
    }
    return c.JSON(http.StatusOK, student)
}

func getStudentById(pool *pgxpool.Pool, id int) (*Student, error) {
    student := &Student{}

    query := `
        SELECT s.full_name, s.gender, s.birth_date::text, s.group_id, sg.group_name 
        FROM students s
        LEFT JOIN student_groups sg ON s.group_id = sg.id
        WHERE s.id = $1
    `
    row := pool.QueryRow(context.Background(), query, id)

    err := row.Scan(&student.FullName, &student.Gender, &student.BirthDate, &student.GroupID, &student.GroupName)
    if err != nil {
        return nil, err
    }

    return student, nil
}






func GetAllScheduleHandler(c echo.Context) error{
    schedules, err := getAllSchedules(pool)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database error"})
    }
    return c.JSON(http.StatusOK, schedules)
}

func getAllSchedules(pool *pgxpool.Pool) ([]Schedule, error) {
    query := `
        SELECT s.id, s.subject_name, s.time_slot, s.group_id, sg.group_name
        FROM schedule s
        LEFT JOIN student_groups sg ON s.group_id = sg.id
        ORDER BY s.id
    `
    rows, err := pool.Query(context.Background(), query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var schedules []Schedule
    for rows.Next() {
        var schedule Schedule
        err := rows.Scan(&schedule.ID, &schedule.SubjectName, &schedule.TimeSlot, &schedule.GroupID, &schedule.GroupName)
        if err != nil {
            return nil, err
        }
        schedules = append(schedules, schedule)
    }

    return schedules, nil
}

func GetScheduleByGroupHandler(c echo.Context) error{
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid group ID"})
    }

    schedules, err := getScheduleByGroupId(pool, id)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database error"})
    }
    return c.JSON(http.StatusOK, schedules)
}

func getScheduleByGroupId(pool *pgxpool.Pool, groupId int) ([]Schedule, error) {
    query := `
        SELECT s.id, s.subject_name, s.time_slot, s.group_id, sg.group_name
        FROM schedule s
        LEFT JOIN student_groups sg ON s.group_id = sg.id
        WHERE s.group_id = $1
    `
    rows, err := pool.Query(context.Background(), query, groupId)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var schedules []Schedule
    for rows.Next() {
        var schedule Schedule
        err := rows.Scan(&schedule.ID, &schedule.SubjectName, &schedule.TimeSlot, &schedule.GroupID, &schedule.GroupName)
        if err != nil {
            return nil, err
        }
        schedules = append(schedules, schedule)
    }

    return schedules, nil
}