package routes

import (
	"api-go-gustavo/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Requests() {

	router := mux.NewRouter()

	router.HandleFunc("/", controllers.Home).Methods("Get")
	router.HandleFunc("/ex01", controllers.Ex01).Methods("Get")
	router.HandleFunc("/ex02", controllers.Ex02).Methods("Get")
	router.HandleFunc("/ex03", controllers.Ex03).Methods("Get")
	router.HandleFunc("/ex04", controllers.Ex04).Methods("Get")
	router.HandleFunc("/ex05", controllers.Ex05).Methods("Get")
	router.HandleFunc("/ex06", controllers.Ex06).Methods("Get")
	router.HandleFunc("/ex07", controllers.Ex07).Methods("Get")
	router.HandleFunc("/ex08", controllers.Ex08).Methods("Get")
	router.HandleFunc("/ex09", controllers.Ex09).Methods("Get")
	router.HandleFunc("/ex10", controllers.Ex10).Methods("Get")
	router.HandleFunc("/ex11", controllers.Ex11).Methods("Get")
	router.HandleFunc("/ex12", controllers.Ex12).Methods("Get")
	router.HandleFunc("/ex13", controllers.Ex13).Methods("Get")
	router.HandleFunc("/ex14", controllers.Ex14).Methods("Get")
	router.HandleFunc("/ex15", controllers.Ex15).Methods("Get")

	log.Fatal(http.ListenAndServe(":8080", router))
}
