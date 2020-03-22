package server

import (
	"fmt"
	"go-elastic-search-student-api/config"
	"net/http"
)

func StartServer() error {

	appConfig := config.GetConfig()

	router := NewRouter(*appConfig)

	router.InitializeRoutes()

	if err := http.ListenAndServe(fmt.Sprintf(":%s", appConfig.AppPort), router); err != nil {
		return err
	}

	return nil
}
