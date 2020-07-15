package conf

import "github.com/joho/godotenv"

func Init() error {
	godotenv.Load()

	return nil
}
