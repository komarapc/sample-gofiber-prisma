package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"goprisma/db"
	"goprisma/lib"
	"goprisma/routes"
	"log"
	"os"
	"time"
)

var prisma *db.PrismaClient

func init() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

}

func main() {
	//seeder.Run()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Default port
	}
	prisma = db.NewClient()
	err := lib.ConnectToDatabase(prisma)
	if err != nil {
		log.Fatal(err)
	}
	defer lib.DisconnectFromDatabase(prisma)
	app := fiber.New(fiber.Config{Prefork: true, IdleTimeout: 10 * time.Second, EnablePrintRoutes: true})
	app.Use(cors.New(cors.Config{AllowOrigins: "*"}))
	//app.Use(middleware.RateLimiter(60, 30))
	app.Use(cache.New(cache.Config{Expiration: time.Duration(30) * time.Second}))
	routes.SetupRoutes(app, prisma)
	if err := app.Listen(":" + port); err != nil {
		log.Panic(err)
	}
}
