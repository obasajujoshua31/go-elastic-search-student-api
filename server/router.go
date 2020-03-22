package server

import (
	"github.com/gorilla/mux"
	"go-elastic-search-student-api/config"
)

type Router struct {
	*mux.Router
	AppConfig config.Config
}

func (r *Router) InitializeRoutes() {
	r.HandleFunc("/", HandleGetHome())
	r.HandleFunc("/actors", HandleGetAllActors(r.AppConfig))
	r.HandleFunc("/actors/{id}", HandleGetOneActor(r.AppConfig))
	r.HandleFunc("/movies", HandleGetAllMovies(r.AppConfig))
	r.HandleFunc("/movies/{id}", HandleGetOneMovie(r.AppConfig))
	r.HandleFunc("/directors", HandleGetAllDirectors(r.AppConfig))
	r.HandleFunc("/directors/{id}", HandleGetOneDirector(r.AppConfig))
	r.HandleFunc("/search/actors/{search}", HandleSearchActorOrDirector())
	r.HandleFunc("/search/movies/{search}", HandleSearchMovie())
}

func NewRouter(appConfig config.Config) *Router {
	return &Router{mux.NewRouter(), appConfig}
}
