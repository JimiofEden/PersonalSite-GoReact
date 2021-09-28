package web

import (
	"fmt"
	"os"
	"net/http"
	"net/http/httptest"
	"database/sql"
	"time"
	"io"
	"io/ioutil"
	"strings"
	"encoding/json"
	"bytes"

	"main/models"
	"main/utils"
	"main/database"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/go-redis/redis"
)

type App struct {
	Router *mux.Router
	Api *mux.Router
	Db *sql.DB
	Cache *redis.Client
}

var err error

func (a *App) Serve(files http.FileSystem, configuration models.Configuration) {

	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		configuration.PgHost, configuration.PgPort, configuration.PgUser, configuration.PgPass, configuration.PgDbName)
	fmt.Println(dbinfo);
	a.Db, err = sql.Open("postgres", dbinfo)
	utils.CheckAndHandleErr(err)
	defer a.Db.Close()

	a.Cache = redis.NewClient(&redis.Options{
		Addr: configuration.RedisAddress,
		DB: 0,
	})

	fs := http.FileServer(files)

	// Router prep
	a.Router = mux.NewRouter()
	a.Api = a.Router.PathPrefix("/api").Subrouter()
	a.Router.Use(a.cacheMiddleware)
	a.Router.PathPrefix("/").Handler(fs)
	graphqlHandler := database.GetGraphqlHandler(a.Db, err)
	
	// Routes
	a.Api.Handle("/data", graphqlHandler).Methods("POST", "OPTIONS")

	// Cors
	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})
	cors := handlers.AllowedOrigins([]string{"*"})

	// Server
	srv := handlers.CombinedLoggingHandler(os.Stdout, a.Router)
	port := fmt.Sprintf(":%s", configuration.ServerPort)
	server := &http.Server {
		Addr: port,
		Handler: handlers.CORS(header, methods, cors)(srv),
	}
	err = server.ListenAndServe()
}


func (a *App) cacheMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" && r.Method != "POST"{
			next.ServeHTTP(w, r)
			return
		}

		// Update with POST request for proper caching
		headerContentType := r.Header.Get("Content-Type")
		body, postErr := ioutil.ReadAll(r.Body)
		utils.CheckAndHandleErr(postErr)
		key := r.RequestURI + headerContentType + string(body)
		//fmt.Println(key)
		content, err := a.Cache.Get(key).Result()
		if err != nil {
			newR := httptest.NewRequest(r.Method, r.RequestURI, strings.NewReader(string(body)))
			newR.Header.Set("Content-Type", "application/json")
			rr := httptest.NewRecorder()
			next.ServeHTTP(rr, newR)
			resp := rr.Result()
			body, _ = io.ReadAll(resp.Body)
			buffer := new(bytes.Buffer)
			err := json.Compact(buffer, body);
			content := buffer.String()

			//fmt.Println(content)
			err = a.Cache.Set(key, content, 10*time.Minute).Err()
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
