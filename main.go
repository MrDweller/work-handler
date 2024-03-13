package main

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	serviceModels "github.com/MrDweller/service-registry-connection/models"
	"github.com/MrDweller/work-handler/database"
	_ "github.com/MrDweller/work-handler/docs"
	workhandler "github.com/MrDweller/work-handler/work-handler"
	"github.com/joho/godotenv"
)

// @title          Work Handler
// @version        1.0
// @description    This page shows the REST interfaces offered by the Work Handler.
// @contact.url    https://github.com/MrDweller/work-handler
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	dbConnectionString := os.Getenv("MONGO_DB_CONNECTION_STRING")
	err = database.InitDatabase(dbConnectionString)
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}

	address := os.Getenv("ADDRESS")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Panic(err)
	}
	systemName := os.Getenv("SYSTEM_NAME")

	serviceRegistryAddress := os.Getenv("SERVICE_REGISTRY_ADDRESS")
	serviceRegistryPort, err := strconv.Atoi(os.Getenv("SERVICE_REGISTRY_PORT"))
	if err != nil {
		log.Panic(err)
	}

	workHandler, err := workhandler.New(address, port, systemName, serviceRegistryAddress, serviceRegistryPort, []serviceModels.ServiceDefinition{
		{
			ServiceDefinition: "create-work",
			ServiceUri:        "/create-work",
		},
		{
			ServiceDefinition: "assign-worker",
			ServiceUri:        "/assign-worker",
		},
	})
	if err != nil {
		log.Panic(err)
	}

	go func() {
		err = workHandler.Run()
		log.Panic(err)

	}()

	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan
	log.Printf("Stopping the work handler!")

	err = workHandler.Stop()
	if err != nil {
		log.Panic(err)
	}
}
