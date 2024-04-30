package lib

import (
	"goprisma/db"
	"log"
)

func ConnectToDatabase(prisma *db.PrismaClient) error {
	//connect to the database
	err := prisma.Connect()
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func DisconnectFromDatabase(prisma *db.PrismaClient) error {
	//disconnect from the database
	err := prisma.Disconnect()
	if err != nil {
		log.Fatal(err)
	}
	return err
}
