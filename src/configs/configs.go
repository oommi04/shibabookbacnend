package configs

import (
	"github.com/joho/godotenv"
	"os"
)

type Configs struct {
	GoogleMapKey                 string
	PORT                         string
	ChanelSecretLineBot          string
	ChanelTokenLineBot           string
	AdminIdLineBot               string
	TestNotificationAdminLineBot string
}

func New() *Configs {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	return &Configs{
		os.Getenv("GOOGLEMAPKEY"),
		os.Getenv("PORT"),
		os.Getenv("CHANELSECRETLINEBOT"),
		os.Getenv("CHANELTOKENLINEBOT"),
		os.Getenv("ADMINIDLINEBOT"),
		os.Getenv("TESTNOTIFICATIONADMINLINEBOT"),
	}
}
