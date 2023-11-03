package main

import (
	"fmt"
	"log"
	"net/http"
	"refactoring/internal/app/controller"
	"refactoring/internal/app/middleware"
	"refactoring/internal/app/router"
	"refactoring/internal/config"
)


func main() {
	log := log.Default()

	log.Println("Сервер запускается")

	cfg := config.Load()

	userStore := middleware.NewStore(cfg.Store)
	userService := middleware.New(userStore, *log)
	userController := controller.New(userService)
	
	r := router.Router(userController)

	log.Println("Сервер запущен")

	if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), r); err != nil {
		fmt.Println("Ошибка записи сервиса: ", err.Error())
	}
}
