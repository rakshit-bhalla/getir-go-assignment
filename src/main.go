package main

import (
	"net/http"

	"rakshit.dev/m/src/configs"
	"rakshit.dev/m/src/controllers"
	"rakshit.dev/m/src/repositories"
	"rakshit.dev/m/src/routers"
	"rakshit.dev/m/src/services"
	"rakshit.dev/m/src/utils"
)

func main() {
	mux := http.NewServeMux()

	// in-memory api configuration
	inMemoryRepository := repositories.NewInMemoryService()
	inMemoryService := services.NewInMemoryService(inMemoryRepository)
	inMemoryController := controllers.NewInMemoryController(inMemoryService)
	inMemoryRouter := routers.NewInMemoryRouter(inMemoryController)
	mux.Handle(utils.IN_MEMORY_BASE_URL, inMemoryRouter)
	mux.Handle(utils.IN_MEMORY_BASE_URL+"/", inMemoryRouter)

	// records api configuration
	recordsRepository := repositories.NewRecordsService()
	recordsService := services.NewRecordsService(recordsRepository)
	recordsController := controllers.NewRecordsController(recordsService)
	recordsRouter := routers.NewRecordsRouter(recordsController)
	mux.Handle(utils.RECORDS_BASE_URL, recordsRouter)
	mux.Handle(utils.RECORDS_BASE_URL+"/", recordsRouter)

	http.ListenAndServe("localhost:"+configs.PORT, mux)
}
