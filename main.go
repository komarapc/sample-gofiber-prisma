package main

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"goprisma/db"
	"goprisma/lib"
	"goprisma/middleware"
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

func fiberConfig() fiber.Config {
	return fiber.Config{
		Prefork:           true,
		IdleTimeout:       10 * time.Second,
		EnablePrintRoutes: true,
		JSONEncoder:       json.Marshal,
		JSONDecoder:       json.Unmarshal,
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
	defer func(prisma *db.PrismaClient) {
		err := lib.DisconnectFromDatabase(prisma)
		if err != nil {
			log.Fatal(err)
		}
	}(prisma)
	app := fiber.New(fiberConfig())
	app.Use(cors.New(cors.Config{AllowOrigins: "*"}))
	app.Use(middleware.RateLimiter(60, 30))
	app.Use(middleware.Cache(5))
	routes.SetupRoutes(app, prisma)
	if err := app.Listen(":" + port); err != nil {
		log.Panic(err)
	}
}
