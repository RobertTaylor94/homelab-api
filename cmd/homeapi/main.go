package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jessevdk/go-flags"
	_ "github.com/lib/pq"
	"murvoth.co.uk/homeapi/handlers"
)

type EnvOptions struct {
	PgHost string `long:"pg-hosy" env:"POSTGRES_HOST" description:"postgres db host IP" default:"postgres" required:"true"`
	PgUser string `long:"pg-user" env:"POSTGRES_USER" description:"postgres db user name" default:"" required:"true"`
	PgPass string `long:"pg-pass" env:"POSTGRES_PASS" description:"postgres db password" default:"" required:"true"`
	PgDb   string `long:"pg-db" env:"POSTGRES_DB" description:"postgres db name" default:"" required:"true"`
	Port   string `long:"port" env:"API_PORT" default:"" required:"true"`
}

var opts EnvOptions

func init() {
	parser := flags.NewParser(&opts, flags.Default)
	if _, err := parser.Parse(); err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", opts.PgUser, opts.PgPass, opts.PgHost, opts.PgDb)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}

	fmt.Println("Connected to the database successfully!")

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.Use(gin.Logger())

	s := http.Server{
		Addr:    fmt.Sprintf(":%v", opts.Port),
		Handler: r,
	}

	r.GET("/", handlers.Ping())

	// Home Endpoints
	r.POST("/api/homekit", handlers.HomeKitPost(db))

	// HealthKit Endpoints
	r.POST("/api/healthkit/heart", handlers.HealthKitHeartPost(db))
	r.POST("/api/healthkit/steps", handlers.HealthKitStepsPost(db))
	r.POST("/api/healthkit/energy", handlers.HealthKitEnergyPost(db))

	fmt.Printf("Starting HomeAPI on Port: %v\n", opts.Port)
	err = s.ListenAndServe()
	if err != nil {
		log.Fatalf("Error starting HomeAPI: %v", err)
	}
}
