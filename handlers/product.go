package handlers

import (
	"encoding/json"
	productdto "fundamental-golang/dto/product"
	dto "fundamental-golang/dto/result"
	"fundamental-golang/models"
	"fundamental-golang/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerProduct struct {
  ProductRepository repositories.ProductRepository
}

func HandlerProduct(ProductRepository repositories.ProductRepository) *handlerProduct {
  return &handlerProduct{ProductRepository}
}

func (h *handlerProduct) FindProducts(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  products, err := h.ProductRepository.FindProducts()
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
    return
  }

  w.WriteHeader(http.StatusOK)
  response := dto.SuccessResult{Code: http.StatusOK, Data: products}
  json.NewEncoder(w).Encode(response)
}

func (h *handlerProduct) GetProduct(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  id, _ := strconv.Atoi(mux.Vars(r)["id"])

  var product models.Product
  product, err := h.ProductRepository.GetProduct(id)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
    return
  }

  w.WriteHeader(http.StatusOK)
  response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseProduct(product)}
  json.NewEncoder(w).Encode(response)
}

func (h *handlerProduct) CreateProduct(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  request := new(productdto.ProductRequest)
  if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
    w.WriteHeader(http.StatusBadRequest)
    response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
    return
  }

  validation := validator.New()
  err := validation.Struct(request)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
    return
  }

  product := models.Product{
    Title:  request.Title,
    Price:  request.Price,
    Image:  request.Image,
    UserID: request.UserID,
  }

  product, err = h.ProductRepository.CreateProduct(product)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
    json.NewEncoder(w).Encode(response)
    return
  }

  product, _ = h.ProductRepository.GetProduct(product.ID)

  w.WriteHeader(http.StatusOK)
  response := dto.SuccessResult{Code: http.StatusOK, Data: product}
  json.NewEncoder(w).Encode(response)
}

func convertResponseProduct(u models.Product) models.ProductResponse {
  return models.ProductResponse{
    Title:    u.Title,
    Price:    u.Price,
    Image:    u.Image,
    UserID: u.UserID,
    User:     u.User,
    Category: u.Category,
  }
}

func (h *handlerProduct) UpdateProduct(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  request := new(productdto.ProductRequest)
  if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
    w.WriteHeader(http.StatusBadRequest)
    response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}  
    json.NewEncoder(w).Encode(response)
    return
  }

  id, _ := strconv.Atoi(mux.Vars(r)["id"])

  product, _ := h.ProductRepository.GetProduct(id)

  if request.Title != "" {
    product.Title = request.Title
  }

  if request.Price != 0 {
    product.Price = request.Price
  }

  if request.Image != "" {
    product.Image = request.Image
  }

  data, err := h.ProductRepository.UpdateProduct(product,id)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()} 
    json.NewEncoder(w).Encode(response)
    return
    }

  w.WriteHeader(http.StatusOK)
  response := dto.SuccessResult{Code: http.StatusOK, Data:convertResponseProduct(data)}
  json.NewEncoder(w).Encode(response)
}

func (h *handlerProduct) DeleteProduct(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  id, _ := strconv.Atoi(mux.Vars(r)["id"])

  product, err := h.ProductRepository.GetProduct(id)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}  
    json.NewEncoder(w).Encode(response)
    return
  }

  data, err := h.ProductRepository.DeleteProduct(product, id)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()} 
    json.NewEncoder(w).Encode(response)
    return
  }

  w.WriteHeader(http.StatusOK)
  response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseProduct(data)}
  json.NewEncoder(w).Encode(response)
}


