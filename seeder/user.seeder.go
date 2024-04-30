package seeder

import (
	"context"
	"goprisma/db"
	"goprisma/lib"
	"log"

	gonanoid "github.com/matoous/go-nanoid"

	"time"

	"github.com/jaswdr/faker"
)

type User struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

func GenerateUserSeed() {
	length := 10
	users := make([]User, length)
	f := faker.New()
	for i := 0; i < length; i++ {
		id, _ := gonanoid.ID(21)
		hashPassword, _ := lib.HashPassword("password")
		users[i] = User{
			ID:        id,
			Name:      f.Person().Name(),
			Email:     f.Internet().Email(),
			Password:  hashPassword,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: nil,
		}
	}

	prisma := db.NewClient()
	err := lib.ConnectToDatabase(prisma)
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range users {
		_, _ = prisma.User.CreateOne(
			db.User.Name.Set(user.Name),
			db.User.Email.Set(user.Email),
			db.User.Password.Set(user.Password),
			db.User.CreatedAt.Set(user.CreatedAt),
			db.User.UpdatedAt.Set(user.UpdatedAt),
		).Exec(context.Background())
	}
	_ = prisma.Disconnect()
}
