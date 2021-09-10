package web

import (
	"fmt"
	"net/http"
	"main/models"
	"main/utils"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Serve(files http.FileSystem, addr string) {
	fs := http.FileServer(files)

	router := mux.NewRouter()

	// Routes
	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/data", func(rw http.ResponseWriter, r *http.Request) {
		utils.RespondWithJson(rw, models.NewApiResponse("OK", models.NewApiData("Hello, world!")))
	}).Methods("GET", "OPTIONS")

	api.HandleFunc("/metrics", getMetrics).Methods("GET", "OPTIONS")

	api.HandleFunc("/data", func(rw http.ResponseWriter, r *http.Request) {
		utils.RespondWithJson(rw, models.NewApiResponse("OK", "Successfully Saved!"))
	}).Methods("POST", "OPTIONS")

	router.PathPrefix("/").Handler(fs)
	cors := handlers.AllowedOrigins([]string{"*"})

	srv := handlers.CombinedLoggingHandler(os.Stdout, router)
	server := &http.Server {
		Addr: addr,
		Handler: handlers.CORS(cors)(srv),
	}
	fmt.Println("Listening on: ", addr)
	err := server.ListenAndServe()
	fmt.Println(err)
}

func getMetrics(rw http.ResponseWriter, r *http.Request) {
	utils.RespondWithJson(rw, models.NewApiResponse("OK", "Metrics go here"))
}