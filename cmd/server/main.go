package main

import (
	"encoding/json"
	"github.com/EdsonGustavoTofolo/apis-standards/internal/dto"
	"github.com/EdsonGustavoTofolo/apis-standards/internal/entity"
	"github.com/EdsonGustavoTofolo/apis-standards/internal/infra/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	//config := configs.LoadConfig("./cmd/server/")

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)

	productHandler := NewProductHandler(productDB)

	http.HandleFunc("/products", productHandler.CreateProduct)

	http.ListenAndServe(":8000", nil)
}

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
