package main

import (
	"LionChallenge/driver"
	"LionChallenge/login/handler"
	"LionChallenge/login/repo"
	"LionChallenge/login/usecase"
	"LionChallenge/middleware"
	"log"
	"net/http"

	usrH "LionChallenge/user/handler"
	usrR "LionChallenge/user/repo"
	usrU "LionChallenge/user/usecase"

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

	userRepo := usrR.CreateUserRepoImpl(db)
	userUsecase := usrU.CreateUserUsecaseImpl(userRepo)
	usrH.CreateUserHandler(router, userUsecase)

	loginRepo := repo.CreateLoginRepoImpl(db)
	loginUsecase := usecase.CreateLoginUsecaseImpl(loginRepo)
	handler.CreateLoginHandler(router, loginUsecase)

	headerOK := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originOK := handlers.AllowedOrigins([]string{"*"})
	methodOK := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"})

	log.Println("Server start at server http://localhost" + viper.GetString("server.port"))
	log.Fatalln(http.ListenAndServe(viper.GetString("server.port"), handlers.CORS(headerOK, originOK, methodOK)(router)))

}
