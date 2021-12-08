package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"techtrain-mission/src/presen/request"
	"techtrain-mission/src/presen/response"
	"techtrain-mission/src/usecase"
)

type UserHandler interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{userUsecase: userUsecase}
}

func (uh *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Println("Method not Found")
		return
	}

	var req request.UserCreateRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		log.Println(err)
	}

	createdUser, err := uh.userUsecase.Create(req.Name)
	if err != nil {
		log.Println(err)
	}

	res := response.UserCreateResponse{
		Token: createdUser.Token,
	}

	w.WriteHeader(http.StatusCreated)

	je := json.NewEncoder(w)
	if err := je.Encode(res); err != nil {
		log.Println(err)
	}
}

func (uh *userHandler) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Println("Method not Found")
		return
	}

	token := r.Header.Get("X-Token")

	if token == "" {
		log.Println("token not found")
	}

	targetUser, err := uh.userUsecase.Get(token)

	if err != nil {
		log.Println(err)
	}

	res := response.UserGetResponse{
		Name: targetUser.Name,
	}

	w.WriteHeader(http.StatusOK)

	je := json.NewEncoder(w)
	if err := je.Encode(res); err != nil {
		log.Println(err)
	}
}

func (uh *userHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Println("Method not Found")
		return
	}

	token := r.Header.Get("X-Token")

	if token == "" {
		log.Println("token not found")
	}

	var req request.UserUpdateRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)

	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusNoContent)

	_, err = uh.userUsecase.Update(req.Name, token)
	if err != nil {
		log.Println(err)
	}
}
