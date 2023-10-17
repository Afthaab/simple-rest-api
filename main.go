package main

import (
	"log"
	"net/http"

	"github.com/afthaab/simple-rest-api/db"
	"github.com/afthaab/simple-rest-api/handler"
	"github.com/afthaab/simple-rest-api/middleware"
)

func main() {
	// add the details to the database
	db.AddStudentsData()

	// endpoints and handler functions
	http.HandleFunc("/students/viewall", middleware.Auth(handler.GetAllStudents))
	http.HandleFunc("/students/view", middleware.Auth(handler.GetStudentByID))

	// start the server
	log.Println("Server started and running at localhost:8080")
	panic(http.ListenAndServe(":8080", nil))
}
