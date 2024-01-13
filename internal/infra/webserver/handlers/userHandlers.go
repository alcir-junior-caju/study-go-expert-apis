package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/alcir-junior-caju/study-go-expert-client-apis/internal/dto"
	"github.com/alcir-junior-caju/study-go-expert-client-apis/internal/entity"
	"github.com/alcir-junior-caju/study-go-expert-client-apis/internal/infra/database"
	"github.com/go-chi/jwtauth"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandlerStruct struct {
	UserDatabase database.UserInterface
	// Load configs by struct
	// JWT          *jwtauth.JWTAuth
	// JWTExpiresIn int
}

func UserHandler(database database.UserInterface /*, JWT *jwtauth.JWTAuth, JWTExpiresIn int*/) *UserHandlerStruct {
	return &UserHandlerStruct{
		UserDatabase: database,
		// JWT:          JWT,
		// JWTExpiresIn: JWTExpiresIn,
	}
}

// CreateUser godoc
// @Summary      Create user
// @Description  Create user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request     body      dto.CreateUserInput  true  "user request"
// @Success      201
// @Failure      500         {object}  Error
// @Router       /users [post]
func (userHandler *UserHandlerStruct) CreateUser(writer http.ResponseWriter, request *http.Request) {
	var user dto.CreateUserInput
	errorDecode := json.NewDecoder(request.Body).Decode(&user)
	if errorDecode != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(errorDecode.Error())
		return
	}
	userCreated, errorCreated := entity.User(user.Name, user.Email, user.Password)
	if errorCreated != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(errorCreated.Error())
		return
	}
	errorDatabase := userHandler.UserDatabase.Create(userCreated)
	if errorDatabase != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(errorDatabase.Error())
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

// GetJWT godoc
// @Summary      Get a user JWT
// @Description  Get a user JWT
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request   body     dto.GetJWTInput  true  "user credentials"
// @Success      200  {object}  dto.GetJWTOutput
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /users/login [post]
func (userHandler *UserHandlerStruct) GetJWT(writer http.ResponseWriter, request *http.Request) {
	jwt := request.Context().Value("JWT").(*jwtauth.JWTAuth)
	jwtExpiresIn := request.Context().Value("JWTExpiresIn").(int)
	var user dto.GetJWTInput
	errorDecode := json.NewDecoder(request.Body).Decode(&user)
	if errorDecode != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(errorDecode.Error())
		return
	}
	userExists, errorExists := userHandler.UserDatabase.FindByEmail(user.Email)
	if errorExists != nil {
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(errorExists.Error())
		return
	}
	if !userExists.ValidatePassword(user.Password) {
		writer.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(writer).Encode("invalid password")
		return
	}
	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		"sub": userExists.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})
	accessToken := dto.GetJWTOutput{AccessToken: tokenString}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(accessToken)
}
