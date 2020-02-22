package main

import (
	"LionChallenge/driver"
	"LionChallenge/middleware"
	"log"
	"net/http"

	"LionChallenge/user/handler"
	"LionChallenge/user/repo"
	"LionChallenge/user/usecase"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func main() {
	db := driver.Config()

	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	router := mux.NewRouter()
	router.Use(middleware.Logging)

	userRepo := repo.CreateUserRepoImpl(db)
	userUsecase := usecase.CreateUserUsecaseImpl(userRepo)
	handler.CreateUserHandler(router, userUsecase)

	headerOK := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originOK := handlers.AllowedOrigins([]string{"*"})
	methodOK := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"})

	log.Println("Server start at server http://localhost" + viper.GetString("server.port"))
	log.Fatalln(http.ListenAndServe(viper.GetString("server.port"), handlers.CORS(headerOK, originOK, methodOK)(router)))

}
