package util

import (
	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
	"log"
)

type Environment struct {
	DBUrl string `env:"DATABASE_URL"`
	Extras env.EnvSet
}

func LoadEnv() *Environment {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	var environment Environment
	es, err := env.UnmarshalFromEnviron(&environment)
	if err != nil {
		log.Fatal(err)
	}

	environment.Extras = es

	return &environment
}
