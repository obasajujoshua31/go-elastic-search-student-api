package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"go-elastic-search-student-api/config"
	"go-elastic-search-student-api/services"
	"log"
	"net/http"
	"strconv"
)

const (
	couldNotConnnect                          = "Could not connect to Database"
	couldNotGetAllActors                      = "Could not get all actors"
	jsonMarshalError                          = "Could not marshal json"
	invalidIDParameter                        = "Invalid id parameter"
	couldNotGetActor                          = "Could not get one actor"
	couldNotGetMovies                         = "Could not get movies"
	couldNotGetOneMovie                       = "Could not get one movie"
	couldNotGetDirectors                      = "Could not get directors"
	couldNotGetOneDirector                    = "Could not get one director"
	couldNotConnectToElasticSearch            = "Could not connect to elastic search"
	couldNotGetValidResponseFromElasticDaemon = "Could not get valid response from elastic demon"
)

func HandleGetHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to movies API")
	}
}

func HandleGetAllActors(config config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db, err := services.ConnectToDatabase(config)
		log.Println(err)
		if err != nil {
			http.Error(w, couldNotConnnect, http.StatusInternalServerError)
			return
		}

		db.SetPool()
		defer db.Close()

		actors, err := db.GetAllActors()
		if err != nil {
			http.Error(w, couldNotGetAllActors, http.StatusInternalServerError)
			return
		}

		writeOutput(w, actors)
	}
}

func HandleGetOneActor(config config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db, err := services.ConnectToDatabase(config)
		log.Println(err)
		if err != nil {
			http.Error(w, couldNotConnnect, http.StatusInternalServerError)
			return
		}

		db.SetPool()
		defer db.Close()

		id := mux.Vars(r)["id"]

		parsedId, err := strconv.ParseInt(id, 10, 10)
		if err != nil {
			http.Error(w, invalidIDParameter, http.StatusBadRequest)
			return
		}

		actor, err := db.GetOneActor(int(parsedId))
		if gorm.IsRecordNotFoundError(err) {
			http.Error(w, couldNotGetActor, http.StatusNotFound)
			return
		}

		if err != nil {
			http.Error(w, couldNotGetActor, http.StatusInternalServerError)
			return
		}

		writeOutput(w, actor)
	}
}

func HandleGetAllMovies(config config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db, err := services.ConnectToDatabase(config)
		log.Println(err)
		if err != nil {
			http.Error(w, couldNotConnnect, http.StatusInternalServerError)
			return
		}

		db.SetPool()
		defer db.Close()

		movies, err := db.GetAllMovies()
		if err != nil {
			http.Error(w, couldNotGetMovies, http.StatusInternalServerError)
			return
		}

		writeOutput(w, movies)
	}
}

func HandleGetOneMovie(config config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db, err := services.ConnectToDatabase(config)
		log.Println(err)
		if err != nil {
			http.Error(w, couldNotConnnect, http.StatusInternalServerError)
			return
		}

		db.SetPool()
		defer db.Close()

		id := mux.Vars(r)["id"]

		parsedId, err := strconv.ParseInt(id, 10, 10)
		if err != nil {
			http.Error(w, invalidIDParameter, http.StatusBadRequest)
			return
		}

		movie, err := db.GetOneActor(int(parsedId))
		if gorm.IsRecordNotFoundError(err) {
			http.Error(w, couldNotGetOneMovie, http.StatusNotFound)
			return
		}

		if err != nil {
			http.Error(w, couldNotGetOneMovie, http.StatusInternalServerError)
			return
		}

		writeOutput(w, movie)
	}
}

func HandleGetAllDirectors(config config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db, err := services.ConnectToDatabase(config)
		log.Println(err)
		if err != nil {
			http.Error(w, couldNotConnnect, http.StatusInternalServerError)
			return
		}

		db.SetPool()
		defer db.Close()

		directors, err := db.GetAllDirectors()
		if err != nil {
			http.Error(w, couldNotGetDirectors, http.StatusInternalServerError)
			return
		}

		writeOutput(w, directors)
	}
}

func HandleGetOneDirector(config config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db, err := services.ConnectToDatabase(config)
		log.Println(err)
		if err != nil {
			http.Error(w, couldNotConnnect, http.StatusInternalServerError)
			return
		}

		db.SetPool()
		defer db.Close()

		id := mux.Vars(r)["id"]

		parsedId, err := strconv.ParseInt(id, 10, 10)
		if err != nil {
			http.Error(w, invalidIDParameter, http.StatusBadRequest)
			return
		}

		director, err := db.GetOneDirector(int(parsedId))
		if gorm.IsRecordNotFoundError(err) {
			http.Error(w, couldNotGetOneDirector, http.StatusNotFound)
			return
		}

		if err != nil {
			http.Error(w, couldNotGetOneDirector, http.StatusInternalServerError)
			return
		}

		writeOutput(w, director)
	}
}

func HandleSearchActorOrDirector() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		client, err := services.ConnectToESClient()

		if err != nil {
			http.Error(w, couldNotConnectToElasticSearch, http.StatusInternalServerError)
			return
		}

		searchParams := mux.Vars(r)["search"]

		var actor services.Actor

		actors, err := client.SearchClient(&actor, "name", searchParams)

		if err != nil {
			http.Error(w, couldNotGetValidResponseFromElasticDaemon, http.StatusInternalServerError)
			return
		}

		writeOutput(w, actors)
	}
}

func HandleSearchMovie() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		client, err := services.ConnectToESClient()

		if err != nil {
			http.Error(w, couldNotConnectToElasticSearch, http.StatusInternalServerError)
			return
		}

		searchParams := mux.Vars(r)["search"]

		var movie services.Movie

		movies, err := client.SearchClient(movie, "title", searchParams)

		if err != nil {
			fmt.Println(err)
			http.Error(w, couldNotGetValidResponseFromElasticDaemon, http.StatusInternalServerError)
			return
		}

		writeOutput(w, movies)
	}
}

func writeOutput(w http.ResponseWriter, data interface{}) {
	dataByte, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		http.Error(w, jsonMarshalError, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(dataByte)
}
