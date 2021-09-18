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

	api.HandleFunc("/metrics", getMetrics).Methods("GET", "OPTIONS")

	api.HandleFunc("/appsettings", func(rw http.ResponseWriter, r *http.Request) {
		utils.RespondWithJson(rw, models.NewApiResponse("OK", models.NewApiData("This should be a list of appsettings!")))
	}).Methods("GET", "OPTIONS")

	api.HandleFunc("/links", func(rw http.ResponseWriter, r *http.Request) {
		// TODO - These should come from a graphql request to postgres
		links := []models.StoredLink {
			models.NewStoredLink("twitter", "https://twitter.com/JimiofEden"),
			models.NewStoredLink("resume", "./AH-Resume_0721-linkedin.pdf"),
			models.NewStoredLink("github", "https://github.com/jimiofeden"),
			models.NewStoredLink("email", "mailto:jimiofeden@gmail.com"),
		}
		utils.RespondWithJson(rw, models.NewApiResponse("OK", links))
	}).Methods("GET", "OPTIONS")

	api.HandleFunc("/skills", func(rw http.ResponseWriter, r *http.Request) {
		// TODO - These should come from a graphql request to postgres
		skills := []models.Skill {

			models.NewSkill("Go", "Backend", "https://github.com/JimiofEden/PersonalSite-GoReact/tree/main/src/serverapp", "This site is being served by Go!"),
			models.NewSkill("GraphQL", "Backend", "https://github.com/JimiofEden/PersonalSite-GoReact", "This data is being retrieved using GraphQL!"),
			models.NewSkill("C#", "Backend", "", ""),
			models.NewSkill(".NET", "Backend", "", ""),
			models.NewSkill("Node", "Backend", "", ""),
			models.NewSkill("Python", "Backend", "", ""),
			models.NewSkill("Ruby", "Backend", "", ""),

			models.NewSkill("React", "Frontend", "https://github.com/JimiofEden/PersonalSite-GoReact/tree/main/src/clientapp", "This site is being rendered with React!"),
			models.NewSkill("Angular", "Frontend", "", ""),
			models.NewSkill("Knockout", "Frontend", "", ""),

			models.NewSkill("Postgres", "Database", "https://github.com/JimiofEden/PersonalSite-GoReact", "This data is being stored in Postgres!"),
			models.NewSkill("Graphql", "Database", "", ""),
			models.NewSkill("SQL Server", "Database", "", ""),
			models.NewSkill("MySQL", "Database", "", ""),

			models.NewSkill("Docker", "ServerSide", "", ""),
			models.NewSkill("Azure", "ServerSide", "", ""),
			models.NewSkill("IIS", "ServerSide", "", ""),
			models.NewSkill("Apache", "ServerSide", "", ""),

			models.NewSkill("Continuous Integration", "Misc.", "", "Teamcity, Octoput Deploy"),
			models.NewSkill("Version Control", "Misc.", "", "Git, SVN"),
			models.NewSkill("Unit Testing", "Misc.", "", "Jest, NUnit"),
			models.NewSkill("Project Management", "Misc.", "", "Agile, Scrum"),
			models.NewSkill("Translating Business Needs into Actionable Goals", "Misc.", "", ""),
		}
		utils.RespondWithJson(rw, models.NewApiResponse("OK", skills))
	}).Methods("GET", "OPTIONS")

	
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