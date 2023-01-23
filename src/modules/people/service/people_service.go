package service

import (
	"encoding/json"
	"go-learn-api/src/modules/people/model"
	"go-learn-api/src/modules/people/repository"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type peopleService struct {
	repository repository.PeopleRepository
}

func NewPeopleService(repository repository.PeopleRepository) peopleService {
	return peopleService{repository}
}

func (s peopleService) FindAllPeople(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	peoples, err := repository.PeopleRepository.FindAll(s.repository)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		errMessage := map[string]string{}
		errMessage["error"] = err.Error()

		errMessageJson, err := json.Marshal(errMessage)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{"error": "Error unmarshalling error message"}`))
			return
		}
		response.Write([]byte(errMessageJson))
		return
	}

	result, err := json.Marshal(peoples)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error marshalling data"}`))
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write(result)
}

func (s peopleService) FindPeopleById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)
	paramId := params["id"]
	id, err := strconv.Atoi(paramId)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error get paramater id"}`))
		return
	}

	people, err := repository.PeopleRepository.FindById(s.repository, id)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		errMessage := map[string]string{}
		errMessage["error"] = err.Error()

		errMessageJson, err := json.Marshal(errMessage)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{"error": "Error unmarshalling error message"}`))
			return
		}
		response.Write([]byte(errMessageJson))
		return
	}

	result, err := json.Marshal(people)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error marshalling data"}`))
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write(result)
}

func (s peopleService) SavePeople(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var people *model.People
	err := json.NewDecoder(request.Body).Decode(&people)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error unmarshalling data"}`))
		return
	}

	err = repository.PeopleRepository.Save(s.repository, people)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		errMessage := map[string]string{}
		errMessage["error"] = err.Error()

		errMessageJson, err := json.Marshal(errMessage)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{"error": "Error unmarshalling error message"}`))
			return
		}
		response.Write([]byte(errMessageJson))
		return
	}

	response.WriteHeader(http.StatusCreated)
	response.Write([]byte(`{"message": "Success save new data"}`))
}

func (s peopleService) UpdatePeople(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var people *model.People

	params := mux.Vars(request)
	paramId := params["id"]
	id, err := strconv.Atoi(paramId)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error get paramater id"}`))
		return
	}

	err = json.NewDecoder(request.Body).Decode(&people)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error unmarshalling data"}`))
		return
	}

	err = repository.PeopleRepository.Update(s.repository, id, people)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		errMessage := map[string]string{}
		errMessage["error"] = err.Error()

		errMessageJson, err := json.Marshal(errMessage)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{"error": "Error unmarshalling error message"}`))
			return
		}
		response.Write([]byte(errMessageJson))
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write([]byte(`{"message": "Success update data"}`))
}

func (s peopleService) DeletePeople(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)
	paramId := params["id"]
	id, err := strconv.Atoi(paramId)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error get paramater id"}`))
		return
	}

	err = repository.PeopleRepository.Delete(s.repository, id)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		errMessage := map[string]string{}
		errMessage["error"] = err.Error()

		errMessageJson, err := json.Marshal(errMessage)
		if err != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{"error": "Error unmarshalling error message"}`))
			return
		}
		response.Write([]byte(errMessageJson))
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write([]byte(`{"message": "Success delete data"}`))
}
