package httpserver

import (
	"log"
	"net/http"
	"senggol/service"

	"github.com/rs/cors"
)

func StartServer(port string, services service.Services) {
	router := makeRouter(services)

	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodHead,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,

		Debug: true,
	}).Handler(router)

	log.Fatal(http.ListenAndServe(":"+port, handler))
}
