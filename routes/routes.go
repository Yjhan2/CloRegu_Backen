package routes // 定义路由

import (
	"myproject/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to the Home Page!"))
	}).Methods("GET")
	router.HandleFunc("/user/register", controllers.Register).Methods("POST")
	router.HandleFunc("/user/login", controllers.Login).Methods("POST")

	router.HandleFunc("/user/upload", controllers.Upload).Methods("POST")
	router.HandleFunc("/user/download", controllers.Upload).Methods("POST")

	router.HandleFunc("/user", controllers.GetUsers).Methods("GET")

	return router
}
