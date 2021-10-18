package main

import (
	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
)

type Environment struct {
	Database struct {
		ConnectionUri string `env:"DB_CONNECTION_URI,required=true"`
	}

	Extras env.EnvSet
}

func InitEnvVars() *Environment {
	if err := godotenv.Load(); err != nil {
		panic("Could not load dotenv file: " + err.Error())
	}

	var envVars *Environment = &Environment{}
	extras, err := env.UnmarshalFromEnviron(envVars)
	if err != nil {
		panic("Error loading environment variables: " + err.Error())
	}

	envVars.Extras = extras

	return envVars
}
