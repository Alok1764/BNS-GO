package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Env  string
}

// must pattern
func MustLoad() Config {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		panic("Port is req")

	}

	env := os.Getenv("ENV")
	if env == "" {
		panic("Env is req")

	}

	return Config{
		Port: port,
		Env:  env,
	}

}
