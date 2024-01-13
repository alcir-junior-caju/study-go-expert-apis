package main

import (
	"log"
	"net/http"

	"github.com/alcir-junior-caju/study-go-expert-client-apis/configs"
	_ "github.com/alcir-junior-caju/study-go-expert-client-apis/docs"
	"github.com/alcir-junior-caju/study-go-expert-client-apis/internal/entity"
	"github.com/alcir-junior-caju/study-go-expert-client-apis/internal/infra/database"
	"github.com/alcir-junior-caju/study-go-expert-client-apis/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// @title           Go Expert API Example
// @version         1.0
// @description     Product API with auhtentication
// @termsOfService  http://swagger.io/terms/

// @contact.name   Caju
// @contact.url    http://www.github.com/alcir-junio-caju
// @contact.email  junior@cajucomunica.com.br

// @license.name   Caju License
// @license.url    http://www.github.com/alcir-junio-caju

// @host      localhost:8080
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	configs, errorConfig := configs.LoadConfig(".")
	if errorConfig != nil {
		panic(errorConfig)
	}
	databaseConnection, errorDatabase := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if errorDatabase != nil {
		panic(errorDatabase)
	}
	databaseConnection.AutoMigrate(&entity.UserStruct{}, &entity.ProductStruct{})
	userDatabase := database.UserDatabase(databaseConnection)
	userHandler := handlers.UserHandler(userDatabase /*, configs.TokenAuth, configs.JWTExpiresIn*/)
	productDatabase := database.ProductDatabase(databaseConnection)
	productHandler := handlers.ProductHandler(productDatabase)
	routes := chi.NewRouter()
	routes.Use(middleware.Logger)
	routes.Use(middleware.Recoverer)
	routes.Use(middleware.WithValue("JWT", configs.TokenAuth))
	routes.Use(middleware.WithValue("JWTExpiresIn", configs.JWTExpiresIn))
	routes.Use(CustomMiddleware)
	routes.Route("/users", func(routes chi.Router) {
		routes.Post("/", userHandler.CreateUser)
		routes.Post("/login", userHandler.GetJWT)
	})
	routes.Route("/products", func(routes chi.Router) {
		routes.Use(jwtauth.Verifier(configs.TokenAuth))
		routes.Use(jwtauth.Authenticator)
		routes.Post("/", productHandler.CreateProduct)
		routes.Get("/", productHandler.FindProducts)
		routes.Get("/{id}", productHandler.FindProduct)
		routes.Put("/{id}", productHandler.UpdateProduct)
		routes.Delete("/{id}", productHandler.DeleteProduct)
	})
	routes.Route("/docs", func(routes chi.Router) {
		routes.Get("/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/docs/doc.json")))
	})
	http.ListenAndServe(":8080", routes)
}

func CustomMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("Custom log middleware: %s %s", request.Method, request.URL)
		next.ServeHTTP(writer, request)
	})
}
