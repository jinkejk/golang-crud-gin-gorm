package main

import (
	"github.com/rs/zerolog/log"
	_ "golang-crud-gin/docs"
	"golang-crud-gin/helper"
	"golang-crud-gin/router"
	"net/http"
)

func main() {

	log.Info().Msg("Started Server on port 8888 ...")

	// Router
	routes := router.InitRouter()

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
