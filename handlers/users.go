package handlers

import (
	"encoding/json"
	dto "fundamental-golang/dto/result"
	usersdto "fundamental-golang/dto/users"
	"fundamental-golang/models"
	"fundamental-golang/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handler struct {
  UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handler {
  return &handler{UserRepository} 
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  request := new(usersdto.CreateUserRequest) 
  if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
    w.WriteHeader(http.StatusBadRequest)
    response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}  
    json.NewEncoder(w).Encode(response)
    return
  }

  validation := validator.New()
  err := validation.Struct(request)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}  
    json.NewEncoder(w).Encode(response)
    return
  }

  user := models.User{
    Fullname:     request.Fullname,
    Email:    request.Email,
    Password:    request.Password,
    Phone: request.Phone,
    Location: request.Location,
    Image: request.Image,
    Role: request.Role,
  }

  data, err := h.UserRepository.CreateUser(user)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    json.NewEncoder(w).Encode(err.Error())
  }

  w.WriteHeader(http.StatusOK)
  response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)}
  json.NewEncoder(w).Encode(response)
}

func convertResponse(u models.User) usersdto.UserResponse {
  return usersdto.UserResponse{
    ID:       u.ID,
    Fullname: u.Fullname,
    Email: u.Email,
    Password: u.Password,
    Phone: u.Phone,
    Location: u.Location,
    Image: u.Image,
    Role: u.Role,
    
  }
}

// previous code...

func (h *handler) FindUsers(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  users, err := h.UserRepository.FindUsers()
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
  }

  w.WriteHeader(http.StatusOK)
  response := dto.SuccessResult{Code: http.StatusOK, Data: users}
  json.NewEncoder(w).Encode(response)
}

func (h *handler) GetUser(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  id, _ := strconv.Atoi(mux.Vars(r)["id"])

  user, err := h.UserRepository.GetUser(id)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
    return
  }

  w.WriteHeader(http.StatusOK)
  response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(user)}
  json.NewEncoder(w).Encode(response)
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  request := new(usersdto.UpdateUserRequest)
  if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
    w.WriteHeader(http.StatusBadRequest)
    response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}  
    json.NewEncoder(w).Encode(response)
    return
  }

  id, _ := strconv.Atoi(mux.Vars(r)["id"])

  user := models.User{}

  if request.Fullname != "" {
    user.Fullname = request.Fullname
  }

  if request.Email != "" {
    user.Email = request.Email
  }

  if request.Phone != "" {
    user.Phone = request.Phone
  }

  if request.Location != ""{
    user.Location = request.Location
  }

  if request.Image != "" {
    user.Image = request.Image
  }

  if request.Role != "" {
    user.Role = request.Role
  }

  data, err := h.UserRepository.UpdateUser(user,id)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()} 
    json.NewEncoder(w).Encode(response)
    return
    }

  w.WriteHeader(http.StatusOK)
  response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)}
  json.NewEncoder(w).Encode(response)
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  id, _ := strconv.Atoi(mux.Vars(r)["id"])

  user, err := h.UserRepository.GetUser(id)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}  
    json.NewEncoder(w).Encode(response)
    return
  }

  data, err := h.UserRepository.DeleteUser(user,id)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()} 
    json.NewEncoder(w).Encode(response)
    return
  }

  w.WriteHeader(http.StatusOK)
  response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)}
  json.NewEncoder(w).Encode(response)
}