package handler

import (
	"LionChallenge/login"
	"LionChallenge/model"
	"LionChallenge/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// LoginHandler struct
type LoginHandler struct {
	loginUsecase login.LoginUsecaseInterface
}

// CreateLoginHandler constructor
func CreateLoginHandler(r *mux.Router, loginUsecase login.LoginUsecaseInterface) {
	loginHandler := LoginHandler{loginUsecase}

	r.HandleFunc("/login", loginHandler.loginUser).Methods(http.MethodPost)
}

func (h *LoginHandler) loginUser(w http.ResponseWriter, r *http.Request) {
	user := new(model.User)

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.LogError("CreateLoginHandler", "Error when login", err)
		utils.Response(w, utils.Message(false, "Input is Empty"))
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

	token, err := h.loginUsecase.Login(user)
	if err != nil {
		utils.LogError("CreateLoginHandler", "Error when login", err)
		utils.Response(w, utils.Message(false, "Email or Password is Wrong"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tokenMap := map[string]string{
		"token": token,
	}

	utils.TokenResponse(w, tokenMap)
}
