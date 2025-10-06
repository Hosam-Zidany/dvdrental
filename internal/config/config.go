package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Env  string

	DBHost string
	DBUser string
	DBPass string
	DBName string
	DBPort string

	JWTSecret string
	JWTExpHrs int
}

var AppConfig Config

func LoadConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error while loading env")
	}
	cfg := Config{
		Port: getEnv("PORT", "8080"),
		Env:  getEnv("ENV", "env"),

		DBHost: getEnv("DBHost", "localhost"),
		DBUser: getEnv("DBUser", ""),
		DBPass: getEnv("DBPass", ""),
		DBName: getEnv("DBName", ""),
		DBPort: getEnv("DBPort", "5432"),

		JWTSecret: getEnv("JWTSecret", ""),
	}
	jwtexphrs := 24
	strHours := getEnv("JWTExpHrs", "24")
	if strHours == "" {
		strHours = "24"
	}
	if val, err := strconv.Atoi(strHours); err == nil {
		jwtexphrs = val
	}
	cfg.JWTExpHrs = jwtexphrs

	AppConfig = cfg

	return cfg
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
