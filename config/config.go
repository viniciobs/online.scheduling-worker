package config

import (
	"os"
)

var Schedule_Api_Uri = "http://localhost:8080"

func GetMongoUri() string {
	return os.Getenv("MONGO_URI")
}

func GetDBName() string {
	return os.Getenv("DB_NAME")
}

func GetMessengerBroker() string {
	return os.Getenv("MESSENGER_BROKER")
}
