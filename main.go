package main

import (
	"fmt"
	"go-learn-api/config"
	"go-learn-api/src/modules/people/repository"
	"go-learn-api/src/modules/people/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("GO-LEARN-API")

	db, err := config.GetPostgresDB()

	if err != nil {
		log.Panic(err)
	}

	err = db.Ping()

	if err != nil {
		log.Panic(err)
	}

	peopleRepository := repository.NewPeopleRepository(db)
	peopleService := service.NewPeopleService(peopleRepository)

	router := mux.NewRouter()
	const port string = ":6060"

	router.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Println(response, "Up and Running ...")
	})
	router.HandleFunc("/people/all", peopleService.FindAllPeople).Methods("GET")
	router.HandleFunc("/people/{id}/detail", peopleService.FindPeopleById).Methods("GET")
	router.HandleFunc("/people/create", peopleService.SavePeople).Methods("POST")
	router.HandleFunc("/people/{id}/update", peopleService.UpdatePeople).Methods("PUT")
	router.HandleFunc("/people/{id}/destroy", peopleService.DeletePeople).Methods("DELETE")

	log.Println("Server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))

}
