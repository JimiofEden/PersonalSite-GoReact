package web

import (
	"fmt"
	"os"
	"net/http"
	"database/sql"
	"main/models"
	"main/utils"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	_ "github.com/lib/pq"
)

func Serve(files http.FileSystem, configuration models.Configuration) {
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
			// models.NewStoredLink("twitter", "https://twitter.com/JimiofEden"),
			// models.NewStoredLink("resume", "./AH-Resume_0721-linkedin.pdf"),
			// models.NewStoredLink("github", "https://github.com/jimiofeden"),
			// models.NewStoredLink("email", "mailto:jimiofeden@gmail.com"),
		}
		utils.RespondWithJson(rw, models.NewApiResponse("OK", links))
	}).Methods("GET", "OPTIONS")

	api.HandleFunc("/skills", func(rw http.ResponseWriter, r *http.Request) {
		// TODO - These should come from a graphql request to postgres
		skills := []models.Skill {

			// models.NewSkill("Go", "Backend", "https://github.com/JimiofEden/PersonalSite-GoReact/tree/main/src/serverapp", "This site is being served by Go!"),
			// models.NewSkill("GraphQL", "Backend", "https://github.com/JimiofEden/PersonalSite-GoReact", "This data is being retrieved using GraphQL!"),
			// models.NewSkill("C#", "Backend", "", ""),
			// models.NewSkill(".NET", "Backend", "", ""),
			// models.NewSkill("Node", "Backend", "", ""),
			// models.NewSkill("Python", "Backend", "", ""),
			// models.NewSkill("Ruby", "Backend", "", ""),

			// models.NewSkill("React", "Frontend", "https://github.com/JimiofEden/PersonalSite-GoReact/tree/main/src/clientapp", "This site is being rendered with React!"),
			// models.NewSkill("Angular", "Frontend", "", ""),
			// models.NewSkill("Knockout", "Frontend", "", ""),

			// models.NewSkill("Postgres", "Database", "https://github.com/JimiofEden/PersonalSite-GoReact/tree/main/src/database", "This data is being stored in Postgres!"),
			// // TODO - Put in specific Graphql references
			// models.NewSkill("Graphql", "Database", "https://github.com/JimiofEden/PersonalSite-GoReact/", "This data is being queried by Graphql!"),
			// models.NewSkill("SQL Server", "Database", "", ""),
			// models.NewSkill("MySQL", "Database", "", ""),

			// models.NewSkill("Docker", "ServerSide", "https://github.com/JimiofEden/PersonalSite-GoReact/tree/main/src/server", "This site's host was built with a docker image!"),
			// models.NewSkill("Azure", "ServerSide", "", ""),
			// models.NewSkill("IIS", "ServerSide", "", ""),
			// models.NewSkill("Apache", "ServerSide", "", ""),

			// models.NewSkill("Continuous Integration", "Misc.", "", "Teamcity, Octoput Deploy"),
			// models.NewSkill("Version Control", "Misc.", "", "Git, SVN"),
			// models.NewSkill("Unit Testing", "Misc.", "", "Jest, NUnit"),
			// models.NewSkill("Project Management", "Misc.", "", "Agile, Scrum"),
			// models.NewSkill("Translating Business Needs into Actionable Goals", "Misc.", "", ""),
		}
		utils.RespondWithJson(rw, models.NewApiResponse("OK", skills))
	}).Methods("GET", "OPTIONS")

	h := getGraphqlHandler(configuration)

	api.Handle("/graphql", h).Methods("GET", "POST")

	router.PathPrefix("/").Handler(fs)
	cors := handlers.AllowedOrigins([]string{"*"})

	srv := handlers.CombinedLoggingHandler(os.Stdout, router)
	port := fmt.Sprintf(":%s", configuration.ServerPort)
	server := &http.Server {
		Addr: port,
		Handler: handlers.CORS(cors)(srv),
	}
	fmt.Println("Listening on: ", port)
	err := server.ListenAndServe()
	fmt.Println(err)
}

func getMetrics(rw http.ResponseWriter, r *http.Request) {
	utils.RespondWithJson(rw, models.NewApiResponse("OK", "Metrics go here"))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getGraphqlHandler(configuration models.Configuration) http.Handler {
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		configuration.PgHost, configuration.PgPort, configuration.PgUser, configuration.PgPass, configuration.PgDbName)
	fmt.Println(dbinfo);
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)

	gqlSkillType := graphql.NewObject(graphql.ObjectConfig{
		Name:        "SkillType",
		Description: "A skilltype",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "The identity identifier of the skillType.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if skillType, ok := p.Source.(*models.SkillType); ok {
						return skillType.Id, nil
					}

					return nil, nil
				},
			},
			"skillTypeId": 	&graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "The lookup id of the skill type.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if skillType, ok := p.Source.(*models.SkillType); ok {
						return skillType.SkillTypeId, nil
					}

					return nil, nil
				},
			},
			"skillTypeName": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The name of the skill type.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if skillType, ok := p.Source.(*models.SkillType); ok {
						return skillType.SkillTypeName, nil
					}

					return nil, nil
				},
			},
		},
	})

	gqlSkill := graphql.NewObject(graphql.ObjectConfig{
		Name:        "Skill",
		Description: "A Skill",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "The identifier of the Skill.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if skill, ok := p.Source.(*models.Skill); ok {
						return skill.Id, nil
					}

					return nil, nil
				},
			},
			"skillName": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The name of the skill.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if skill, ok := p.Source.(*models.Skill); ok {
						return skill.SkillName, nil
					}

					return nil, nil
				},
			},
			"skillTypeId": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "The Skill Type Id of the Skill.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if skill, ok := p.Source.(*models.Skill); ok {
						return skill.SkillTypeId, nil
					}

					return nil, nil
				},
			},
			"skillType": &graphql.Field{
				Type: 	gqlSkillType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if skill, ok := p.Source.(*models.Skill); ok {
						skillType := &models.SkillType{}
						err = db.QueryRow("select id, SkillTypeId, SkillTypeName from dbo.SkillType where SkillTypeId = $1", skill.SkillTypeId).Scan(&skillType.Id, &skillType.SkillTypeId, &skillType.SkillTypeName)
						checkErr(err)

						return skillType, nil
					}

					return nil, nil
				},
			},
			"url": &graphql.Field{
				Type:        graphql.String,
				Description: "The url (if any) for the skill.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if skill, ok := p.Source.(*models.Skill); ok {
						return skill.Url, nil
					}

					return nil, nil
				},
			},
			"comment": &graphql.Field{
				Type:        graphql.String,
				Description: "The comment (if any) for the skill.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if skill, ok := p.Source.(*models.Skill); ok {
						return skill.Comment, nil
					}

					return nil, nil
				},
			},
			
		},
	})

	gqlStoredLink := graphql.NewObject(graphql.ObjectConfig{
		Name:        "StoredLink",
		Description: "A link",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "The identity identifier of the link.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if storedLink, ok := p.Source.(*models.StoredLink); ok {
						return storedLink.Id, nil
					}

					return nil, nil
				},
			},
			"name": 		&graphql.Field{
				Type:        graphql.NewNonNull(graphql.Int),
				Description: "The link's name.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if storedLink, ok := p.Source.(*models.StoredLink); ok {
						return storedLink.Name, nil
					}

					return nil, nil
				},
			},
			"url": &graphql.Field{
				Type:        graphql.NewNonNull(graphql.String),
				Description: "The link's url.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if storedLink, ok := p.Source.(*models.StoredLink); ok {
						return storedLink.Url, nil
					}

					return nil, nil
				},
			},
		},
	})

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"skillType": &graphql.Field{
				Type:        gqlSkillType,
				Description: "Get a skilltype lookup.",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(int)

					skillType := &models.SkillType{}
					err = db.QueryRow("select Id, SkillTypeId, SkillTypeName from dbo.SkillType where Id = $1", id).Scan(&skillType.Id, &skillType.SkillTypeId, &skillType.SkillTypeName)
					checkErr(err)

					return skillType, nil
				},
			},
			"skillTypes": &graphql.Field{
				Type:        graphql.NewList(gqlSkillType),
				Description: "List of skill types.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					rows, err := db.Query("SELECT Id, SkillTypeId, SkillTypeName FROM dbo.SkillType")
					checkErr(err)
					var skillTypes []*models.SkillType

					for rows.Next() {
						skillType := &models.SkillType{}

						err = rows.Scan(&skillType.Id, &skillType.SkillTypeId, &skillType.SkillTypeName)
						checkErr(err)
						skillTypes = append(skillTypes, skillType)
					}

					return skillTypes, nil
				},
			},
			"skill": &graphql.Field{
				Type:        gqlSkill,
				Description: "Get a skill lookup.",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(int)

					skill := &models.Skill{}
					err = db.QueryRow("select Id, SkillName, SkillTypeId, Url, Comment from dbo.Skill where Id = $1", id).Scan(&skill.Id, &skill.SkillName, &skill.SkillTypeId, &skill.Url, &skill.Comment)
					checkErr(err)

					return skill, nil
				},
			},
			"skills": &graphql.Field{
				Type:        graphql.NewList(gqlSkill),
				Description: "List of skill types.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					rows, err := db.Query("SELECT Id, SkillName, SkillTypeId, Url, Comment FROM dbo.Skill")
					checkErr(err)
					var skills []*models.Skill

					for rows.Next() {
						skill := &models.Skill{}

						err = rows.Scan(&skill.Id, &skill.SkillName, &skill.SkillTypeId, &skill.Url, &skill.Comment)
						checkErr(err)
						skills = append(skills, skill)
					}

					return skills, nil
				},
			},
			"storedLink": &graphql.Field{
				Type:        gqlStoredLink,
				Description: "Get a storedLink lookup.",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, _ := params.Args["id"].(int)

					storedLink := &models.StoredLink{}
					err = db.QueryRow("select Id, Name, Url from dbo.StoredLink where Id = $1", id).Scan(&storedLink.Id, &storedLink.Name, &storedLink.Url)
					checkErr(err)

					return storedLink, nil
				},
			},
			"storedLinks": &graphql.Field{
				Type:        graphql.NewList(gqlStoredLink),
				Description: "List of storedlink types.",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					rows, err := db.Query("SELECT Id, Name, Url FROM dbo.StoredLink")
					checkErr(err)
					var storedLinks []*models.StoredLink

					for rows.Next() {
						storedLink := &models.StoredLink{}

						err = rows.Scan(&storedLink.Id, &storedLink.Name, &storedLink.Url)
						checkErr(err)
						storedLinks = append(storedLinks, storedLink)
					}

					return storedLinks, nil
				},
			},
		},
	})

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
	})

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	// Return handler to be served by main function
	return h
}
