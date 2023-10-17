package db

import (
	"errors"

	"github.com/afthaab/simple-rest-api/models"
)

var dataBase = make(map[uint64]models.Student)

var ErrNoRecords = errors.New("no student data is present")
var ErrRecordNotFound = errors.New("could not find the student")

func AddStudentsData() {
	dataBase[101] = models.Student{
		Name:  "Mohammed Afthab",
		Age:   22,
		Grade: 9.7,
	}
	dataBase[102] = models.Student{
		Name:  "Purvi",
		Age:   23,
		Grade: 9.8,
	}
	dataBase[103] = models.Student{
		Name:  "Jeevan",
		Age:   22,
		Grade: 9.9,
	}
}

func ViewAllStudens() (map[uint64]models.Student, error) {
	if len(dataBase) == 0 {
		return map[uint64]models.Student{}, ErrNoRecords
	}
	return dataBase, nil
}

func ViewStudentById(uid uint64) (models.Student, error) {
	userData, ok := dataBase[uid]
	if !ok {
		return models.Student{}, ErrRecordNotFound
	}
	return userData, nil
}
