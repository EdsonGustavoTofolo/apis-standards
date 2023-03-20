package handlers

import (
	"encoding/json"
	"github.com/EdsonGustavoTofolo/apis-standards/internal/dto"
	"github.com/EdsonGustavoTofolo/apis-standards/internal/entity"
	"github.com/EdsonGustavoTofolo/apis-standards/internal/infra/database"
	"net/http"
)

type ProductHandler struct {
	ProductDB database.ProductRepository
}

func NewProductHandler(db database.ProductRepository) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p, err := entity.NewProduct(product.Name, product.Price)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = h.ProductDB.Create(p); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
