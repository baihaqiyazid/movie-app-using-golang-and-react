package main

import (
	_ "github.com/lib/pq"
	"server/cmd/app"
	"server/cmd/controller"
	"server/cmd/repository"
	"server/cmd/service"
)

func main() {

	db, _ := app.ConnectDB()

	movieRepository := repository.NewMovieRepository(db)
	movieService := service.NewMovieService(db, movieRepository)
	movieController := controller.NewMovieController(movieService)

	app.Route(movieController)
}

// func main() {
// 	var cfg Config

// 	flag.IntVar(&cfg.port, "port", 4001, "Server listening on ")
// 	flag.StringVar(&cfg.env, "env", "development", "Application environment (development | production)")
// 	flag.StringVar(&cfg.db.dsn, "postgres",
// 				   "postgres://postgres:postgres@localhost:5432/goreactmovies?sslmode=disable",
// 				   "Postgres Connection Config")
// 	flag.Parse()

// 	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

// 	db, err := DBConnect(&cfg)
// 	if err != nil{
// 		logger.Fatal(err)
// 	}

// 	app := &Application{
// 		config: cfg,
// 		logger: logger,
// 		models: models.NewModels(db),
// 	}

// 	fmt.Println("server is running...")

// 	serve := &http.Server{
// 		Addr:         fmt.Sprintf(":%d", cfg.port),
// 		Handler:      app.Route(),
// 		IdleTimeout:  time.Minute,
// 		ReadTimeout:  10 * time.Second,
// 		WriteTimeout: 30 * time.Second,
// 	}
// 	logger.Printf("Starting server on port: %d", cfg.port)
// 	err = serve.ListenAndServe()
// 	if err != nil {
// 		log.Println(err)
// 	}
// }
