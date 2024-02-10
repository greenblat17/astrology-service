package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	// "github.com/joho/godotenv"
	"github.com/greenblat17/worker/domain"
	"github.com/greenblat17/worker/repository"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	// if err := godotenv.Load(); err != nil {
	// 	log.Fatalf("error loading env variables: %s", err.Error())
	// }

	client := initClient()

	req, err := http.NewRequest("GET", "https://api.nasa.gov/planetary/apod", nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("api_key", "B3tYLXowtRCwboU4yOuNWjHlzenYJULLXfb50LGJ")
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Parse response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var apod domain.APOD
	err = json.Unmarshal(body, &apod)
	if err != nil {
		log.Fatal(err)
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: "1412",
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	repo := repository.NewStarPostgres(db)

	star := domain.Star{
		Name:        apod.Title,
		Description: apod.Explanation,
		Date:        apod.Date,
		ImageURL:    apod.URL,
	}

	repo.Save(star)
}

func initClient() *http.Client {
	return &http.Client{
		Timeout: 10 * time.Second,
	}
}

func initConfig() error {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
