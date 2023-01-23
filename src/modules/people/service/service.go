package service

import (
	"net/http"
)

type PeopleService interface {
	FindAllPeople(http.ResponseWriter, *http.Request)
	FindPeopleById(http.ResponseWriter, *http.Request)
	SavePeople(http.ResponseWriter, *http.Request)
	UpdatePeople(http.ResponseWriter, *http.Request)
	DeletePeople(http.ResponseWriter, *http.Request)
}
