package handler

import (
	mdw "LionChallenge/middleware"
	"LionChallenge/model"
	"LionChallenge/user"
	"LionChallenge/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

// UserHandler struct
type UserHandler struct {
	userUsecase user.UserUsecase
}

// CreateUserHandler constructor use for routing
func CreateUserHandler(r *mux.Router, userUsecase user.UserUsecase) {
	handler := UserHandler{userUsecase}

	r.HandleFunc("/user", mdw.TokenVerifyMiddleware(handler.createUser)).Methods(http.MethodPost)
	r.HandleFunc("/user", mdw.TokenVerifyMiddleware(handler.readAllUser)).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}", mdw.TokenVerifyMiddleware(handler.readUser)).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}", mdw.TokenVerifyMiddleware(handler.updateUser)).Methods(http.MethodPut)
	r.HandleFunc("/user/{id}", mdw.TokenVerifyMiddleware(handler.deleteUser)).Methods(http.MethodDelete)
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.Response(w, utils.Message(false, "sorry, url not found"))
		w.WriteHeader(http.StatusNotFound)
		return
	})
}

func (conn *UserHandler) createUser(w http.ResponseWriter, r *http.Request) {
	user := new(model.User)

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.Response(w, utils.Message(false, "Invalid Request "+err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if user.Username == "" {
		utils.Response(w, utils.Message(false, "Username must be filled in"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if user.Email == "" {
		utils.Response(w, utils.Message(false, "Email must be filled in"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if user.Password == "" {
		utils.Response(w, utils.Message(false, "Password must be filled in"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		utils.LogError("createUser", "Error when generate password", err)
		return
	}

	user.Password = string(hash)

	response, err := conn.userUsecase.Create(user)
	if err != nil {
		utils.Response(w, response)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	utils.Response(w, response)
	return
}

func (conn *UserHandler) readUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		utils.Response(w, utils.Message(false, "please provide valid id"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := conn.userUsecase.Read(id)
	if err != nil {
		utils.Response(w, response)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	utils.Response(w, response)
	return
}

func (conn *UserHandler) readAllUser(w http.ResponseWriter, r *http.Request) {
	response, err := conn.userUsecase.ReadAll()
	if err != nil {
		utils.Response(w, response)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	utils.Response(w, response)
	return
}

func (conn *UserHandler) updateUser(w http.ResponseWriter, r *http.Request) {
	user := new(model.User)
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		utils.Response(w, utils.Message(false, "please provide valid id"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		utils.Response(w, utils.Message(false, "Invalid Request "+err.Error()))
		return
	}

	response, err := conn.userUsecase.Update(id, user)
	if err != nil {
		utils.Response(w, response)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	utils.Response(w, response)
	return

}

func (conn *UserHandler) deleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		utils.Response(w, utils.Message(false, "please provide valid id"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	response, err := conn.userUsecase.Delete(id)
	if err != nil {
		utils.Response(w, response)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	utils.Response(w, response)
	return
}
