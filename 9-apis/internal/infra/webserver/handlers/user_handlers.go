package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/leandropiassetta/goexpert/9-apis/internal/dto"
	"github.com/leandropiassetta/goexpert/9-apis/internal/entity"
	database "github.com/leandropiassetta/goexpert/9-apis/internal/infra/database"
)

type UserHandler struct {
	UserDB database.UserInterface

	// JWT -> JSON Web Token is a standard used to create access tokens for an application, him is a string that has three parts: a header, a payload, and a signature.
	JWT           *jwtauth.JWTAuth
	JWTExperiesIn int
}

func NewUserHandler(userDB database.UserInterface, jwt *jwtauth.JWTAuth, JWTExperiesIn int) *UserHandler {
	return &UserHandler{
		UserDB:        userDB,
		JWT:           jwt,
		JWTExperiesIn: JWTExperiesIn,
	}
}

func (handler *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	var user dto.GetJWTInput

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := handler.UserDB.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !u.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, tokenString, _ := handler.JWT.Encode(map[string]interface{}{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(handler.JWTExperiesIn)).Unix(),
	})

	accessToken := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: tokenString,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accessToken)
	w.WriteHeader(http.StatusOK)
}

func (handler *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.UserDB.CreateUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}
