package handlers

import (
	"encoding/json"
	"github.com/EdsonGustavoTofolo/apis-standards/internal/dto"
	"github.com/EdsonGustavoTofolo/apis-standards/internal/entity"
	"github.com/EdsonGustavoTofolo/apis-standards/internal/infra/database"
	"github.com/go-chi/jwtauth"
	"net/http"
	"time"
)

type UserHandler struct {
	UserDB database.UserRepository
}

func NewUserHandler(db database.UserRepository) *UserHandler {
	return &UserHandler{
		UserDB: db,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.UserDB.Create(u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) GetJwt(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("jwtExpiresIn").(int)

	var body dto.GetJwtInput

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.UserDB.FindByEmail(body.Email)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !user.ValidatePassword(body.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	_, token, _ := jwt.Encode(map[string]interface{}{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})

	accesToken := struct {
		AccesToken string `json:"access_token"`
	}{
		AccesToken: token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accesToken)
}
