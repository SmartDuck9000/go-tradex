package main

import (
	"github.com/SmartDuck9000/go-tradex/src/config"
	"github.com/SmartDuck9000/go-tradex/src/service"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	var loader config.ConfigurationLoader = config.EnvLoader{}
	conf := loader.ReadConfig()
	var api service.StatServiceInterface = service.CreateServer(*conf)
	err := api.Run()

	if err != nil {
		log.Print(err.Error())
	}
}
