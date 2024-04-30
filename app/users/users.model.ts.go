package users

import (
	"context"
	"goprisma/db"
	"log"
	"time"
)

type UserResponse struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type UserQueryRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Page    int    `json:"page"`
	PerPage int    `json:"per_page"`
}

func GetAllUsers(query UserQueryRequest, prisma *db.PrismaClient) []UserResponse {
	offset := (query.Page - 1) * query.PerPage
	users, _ := prisma.User.FindMany(
		db.User.DeletedAt.IsNull(),
		db.User.Name.Contains(query.Name),
		db.User.Email.Contains(query.Email),
	).OrderBy(
		db.User.CreatedAt.Order(db.DESC),
	).Select(
		db.User.ID.Field(),
		db.User.Name.Field(),
		db.User.Email.Field(),
		db.User.CreatedAt.Field(),
		db.User.UpdatedAt.Field(),
	).Skip(offset).Take(query.PerPage).Exec(context.Background())
	var response []UserResponse
	for _, user := range users {
		response = append(response, UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
	return response
}

func GetUserById(id string, prisma *db.PrismaClient) *db.UserModel {
	user, _ := prisma.User.FindUnique(
		db.User.ID.Equals(id),
	).Exec(context.Background())
	return user
}

func GetByEmail(email string, prisma *db.PrismaClient) *db.UserModel {
	user, _ := prisma.User.FindUnique(
		db.User.Email.Equals(email),
	).Exec(context.Background())
	return user
}

func CreateOne(user UserRequest, prisma *db.PrismaClient) UserResponse {
	newUser, _ := prisma.User.CreateOne(
		db.User.Name.Set(user.Name),
		db.User.Email.Set(user.Email),
		db.User.Password.Set(user.Password),
	).Exec(context.Background())
	return UserResponse{
		ID:        newUser.ID,
		Name:      newUser.Name,
		Email:     newUser.Email,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
	}
}
func UpdateOne(id string, data UserRequest, prisma *db.PrismaClient) UserResponse {
	user, _ := prisma.User.FindUnique(
		db.User.ID.Equals(id),
	).Update(
		db.User.Name.Set(data.Name),
		db.User.Email.Set(data.Email),
		db.User.Password.Set(data.Password),
	).Exec(context.Background())
	return UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func DeleteOne(id string, prisma *db.PrismaClient) *db.UserModel {
	deleted, err := prisma.User.FindUnique(
		db.User.ID.Equals(id),
	).Update(
		db.User.DeletedAt.Set(time.Now()),
	).Exec(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return deleted
}
