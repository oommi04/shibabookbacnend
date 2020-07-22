package configs

import (
	"github.com/joho/godotenv"
	"os"
)

type Configs struct {
	PORT                         string
}

func New() *Configs {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	return &Configs{
		os.Getenv("PORT"),
	}
}
