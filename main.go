package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
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
	prisma = db.NewClient()
	err := lib.ConnectToDatabase(prisma)
	if err != nil {
		log.Fatal(err)
	}
	defer lib.DisconnectFromDatabase(prisma)
	app := fiber.New(fiber.Config{Prefork: true, IdleTimeout: 10 * time.Second})
	app.Use(cache.New())
	routes.SetupRoutes(app, prisma)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Default port
	}
	err = app.Listen(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
