package initializers

import "github.com/joho/godotenv"

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Environment variable error: ", err)
	}
}
