package main

import (
	"ecommerce/controllers"
	"ecommerce/database"
	"ecommerce/middlewares"
	"ecommerce/sugar"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	pool := database.Connect()
	defer pool.Close()

	server := sugar.New(sugar.Config{
		Host: "localhost:8080",
		Timeout: 15 * time.Second,
		Database: pool,
	})

	server.Middleware("/api/*", middlewares.RoutesProtection)

	// Auth
	server.Post("/api/v1/auth/register", controllers.AuthRegister)
	server.Post("/api/v1/auth/login", controllers.AuthLogin)
	server.Get("/api/v1/auth/validate", controllers.AuthValidate)

	server.Get("/api/v1/me", controllers.GetMe)
	server.Post("api/v1/me/update", controllers.EditMe)

	server.Get("/api/v1/categories", controllers.Categories)
	server.Get("/api/v1/categories/:id", controllers.CategoriesById)
	server.Post("/api/v1/categories", controllers.AddCategory)
	server.Delete("/api/v1/categories/:id", controllers.DeleteCategory)
	server.Patch("/api/v1/categories/:id", controllers.EditCategory)

	server.Get("/api/v1/products", controllers.Products)
	server.Get("/api/v1/products/:id", controllers.ProductById)
	server.Post("/api/v1/products", controllers.AddProduct)
	server.Delete("/api/v1/products/:id", controllers.DeleteProduct)
	server.Patch("/api/v1/products/:id", controllers.EditProduct)

	// Cart
	server.Get("/api/v1/cart", controllers.GetCartItems)
	server.Delete("/api/v1/cart/clear", controllers.ClearCartItems)
	server.Delete("/api/v1/cart/:id", controllers.DeleteCartItem)
	server.Post("/api/v1/cart", controllers.AddCartItem)
	
	server.Listen()
}