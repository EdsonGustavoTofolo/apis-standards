package main

import (
	"github.com/EdsonGustavoTofolo/apis-standards/configs"
	_ "github.com/EdsonGustavoTofolo/apis-standards/docs"
	"github.com/EdsonGustavoTofolo/apis-standards/internal/entity"
	"github.com/EdsonGustavoTofolo/apis-standards/internal/infra/database"
	"github.com/EdsonGustavoTofolo/apis-standards/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
)

// @title Go Expert API Example
// @version 1.o
// @description Product API with authentication
// @termsOfService http://swagger.io/terms/

// @contact.name Edson Gustavo Tofolo
// @contact.url http//github.com/EdsonGustavoTofolo
// @contact.email edsontofolo@gmail.com

// @license.name Edson License
// @license.url http//github.com/

// @host localhost:8000
// @BasePath /
// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	config := configs.LoadConfig("./cmd/server/")

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.WithValue("jwt", config.TokenAuth))
	router.Use(middleware.WithValue("jwtExpiresIn", config.JWTExpiresIn))
	//router.Use(LogRequest)

	router.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	router.Post("/users/token", userHandler.GetJwt)
	router.Post("/users", userHandler.CreateUser)

	router.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	http.ListenAndServe(":8000", router)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
