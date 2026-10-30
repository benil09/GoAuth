package config

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

func load(){
	err := godotenv.Load();
	if err != nil {
		// Log the error if .env file is not found
		fmt.Println("Error loading .env file", err);
	}

}

func GetString (key string , fallback string) string {
	load();

	value , ok := os.LookupEnv(key)

	if !ok {
		return fallback;
	}
	return value;

}
