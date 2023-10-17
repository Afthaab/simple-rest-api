package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/afthaab/simple-rest-api/db"
)

func GetStudentByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userId := r.URL.Query().Get("student_id")

	uid, err := strconv.ParseUint(userId, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		w.Write([]byte(string("could not process the request"))) //generic err msg to the client
		return
	}

	studentData, err := db.ViewStudentById(uid)
	if err != nil {
		errMap := map[string]string{
			"message": err.Error(),
			"status":  http.StatusText(http.StatusNotFound),
		}
		res, err := json.Marshal(errMap)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)
			w.Write([]byte(string("could not process the request"))) //generic err msg to the client
			return
		}

		w.WriteHeader(http.StatusNotFound)
		w.Write(res)
		return
	}

	res, err := json.Marshal(studentData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		w.Write([]byte(string("could not process the request"))) //generic err msg to the clien
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)

	// r.Context().Done()

}

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	studentsData, err := db.ViewAllStudens() // database call
	if err != nil {
		errMap := map[string]string{
			"message": err.Error(),
			"status":  http.StatusText(http.StatusNotFound),
		}

		res, err := json.Marshal(errMap)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)
			w.Write([]byte(string("could not process the request"))) //generic err msg to the client
			return
		}

		w.WriteHeader(http.StatusNotFound)
		w.Write(res)
		return

	}

	res, err := json.Marshal(studentsData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		w.Write([]byte(string("could not process the request"))) //generic err msg to the clien
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(res)

	// r.Context().Done()
}
