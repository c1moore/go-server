package main

import (
	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
)

type Environment struct {
	MongoDB struct {
		ConnectionURI string `env:"MONGODB_URI,required=true"`

		AuthMechanism string `env:"MONGODB_AUTH_MECHANISM,required=true"`
		AuthDatabase  string `env:"MONGODB_AUTH_DATABASE,required=true"`
		Username      string `env:"MONGODB_USERNAME,required=true"`
		Password      string `env:"MONGODB_PASSWORD,required=true"`

		DatabaseName string `env:"MONGODB_DATABASE,required=true"`
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
