package web

import (
	"fmt"
	"os"
	"net/http"
	"net/http/httptest"
	"database/sql"
	"time"

	"main/models"
	"main/utils"
	"main/database"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/go-redis/redis"
)

func Serve(files http.FileSystem, configuration models.Configuration) {

	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		configuration.PgHost, configuration.PgPort, configuration.PgUser, configuration.PgPass, configuration.PgDbName)
	fmt.Println(dbinfo);
	db, err := sql.Open("postgres", dbinfo)
	utils.CheckAndHandleErr(err)
	defer db.Close()

	fs := http.FileServer(files)

	router := mux.NewRouter()

	rCache := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		DB:       0,
	})

	router.Use(cacheMiddleware)

	// Routes
	api := router.PathPrefix("/api").Subrouter()

	api.HandleFunc("/metrics", getMetrics).Methods("GET", "OPTIONS")

	api.HandleFunc("/appsettings", func(rw http.ResponseWriter, r *http.Request) {
		utils.RespondWithJson(rw, models.NewApiResponse("OK", models.NewApiData("This should be a list of appsettings!")))
	}).Methods("GET", "OPTIONS")

	api.HandleFunc("/links", func(rw http.ResponseWriter, r *http.Request) {
		// TODO - These should come from a graphql request to postgres
		links := []models.StoredLink {
		}
		utils.RespondWithJson(rw, models.NewApiResponse("OK", links))
	}).Methods("GET", "OPTIONS")

	api.HandleFunc("/skills", func(rw http.ResponseWriter, r *http.Request) {
		// TODO - These should come from a graphql request to postgres
		skills := []models.Skill {

		}
		utils.RespondWithJson(rw, models.NewApiResponse("OK", skills))
	}).Methods("GET", "OPTIONS")

	graphqlHandler := database.GetGraphqlHandler(db, err)

	api.Handle("/data", graphqlHandler).Methods("POST", "OPTIONS")

	router.PathPrefix("/").Handler(fs)

	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})
	cors := handlers.AllowedOrigins([]string{"*"})

	srv := handlers.CombinedLoggingHandler(os.Stdout, router)
	port := fmt.Sprintf(":%s", configuration.ServerPort)
	server := &http.Server {
		Addr: port,
		Handler: handlers.CORS(header, methods, cors)(srv),
	}
	fmt.Println("Listening on: ", port)
	err = server.ListenAndServe()
	fmt.Println(err)
}

func getMetrics(rw http.ResponseWriter, r *http.Request) {
	utils.RespondWithJson(rw, models.NewApiResponse("OK", "Metrics go here"))
}


func cacheMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			next.ServeHTTP(w, r)
			return
		}
		fmt.Println(r.RequestURI)
		content, err := rCache.Get(r.RequestURI).Result()
		if err != nil {
			rr := httptest.NewRecorder()
			next.ServeHTTP(rr, r)
			content = rr.Body.String()
			fmt.Println(content)

			err = rCache.Set(r.RequestURI, content, 10*time.Minute).Err()
			if err != nil {
				utils.RepondWithError(w, http.StatusInternalServerError, models.NewApiResponse("OK",err.Error()))
			}
			utils.RespondWithJson(w, models.NewApiResponse("OK",content))
			return
		}else {
			utils.RespondWithJson(w, models.NewApiResponse("OK",content))
			return
		}
	})
}
