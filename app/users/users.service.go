package users

import (
	"goprisma/db"
	"goprisma/lib"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GetAllUsersService(query UserQueryRequest, prisma *db.PrismaClient) lib.ResponseData {
	users := GetAllUsers(query, prisma)
	return lib.ResponseSuccess(lib.ResponseProps{Code: fiber.StatusOK, Data: users})
}

func GetUserByIdService(id string, prisma *db.PrismaClient) lib.ResponseData {
	user := GetUserById(id, prisma)

	if user == nil {
		return lib.ResponseError(lib.ResponseProps{Code: fiber.StatusNotFound})
	}
	deletedAt, ok := user.DeletedAt()
	var timeDeletedAt *time.Time
	if ok {
		timeDeletedAt = &deletedAt
	}
	if timeDeletedAt != nil {
		return lib.ResponseError(lib.ResponseProps{Code: fiber.StatusNotFound})
	}
	response := UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: timeDeletedAt,
	}
	return lib.ResponseSuccess(lib.ResponseProps{Code: fiber.StatusOK, Data: response})
}

func CreateOneService(user UserRequest, prisma *db.PrismaClient) lib.ResponseData {
	// validate user input
	errors := validateStoreRequest(user)
	if len(errors) > 0 {
		return lib.ResponseError(lib.ResponseProps{Code: fiber.StatusBadRequest, Message: &errors})
	}
	// check if email already exist
	userEmailExist := GetByEmail(user.Email, prisma)
	if userEmailExist != nil {
		message := "Email already exist"
		return lib.ResponseError(lib.ResponseProps{Code: fiber.StatusConflict, Message: &message})
	}
	hashPasswrd, _ := lib.HashPassword(user.Password)
	newUser := CreateOne(UserRequest{
		Name:     user.Name,
		Email:    user.Email,
		Password: hashPasswrd,
	}, prisma)
	return lib.ResponseSuccess(lib.ResponseProps{Code: fiber.StatusCreated, Data: newUser})
}

func UpdateOneService(id string, user UserRequest, prisma *db.PrismaClient) lib.ResponseData {
	userExist := GetUserById(id, prisma)
	if userExist == nil {
		return lib.ResponseError(lib.ResponseProps{Code: fiber.StatusNotFound})
	}
	errors := validateStoreRequest(user)
	if len(errors) > 0 {
		return lib.ResponseError(lib.ResponseProps{Code: fiber.StatusBadRequest, Message: &errors})
	}
	// check if email already exist
	userEmailExist := GetByEmail(user.Email, prisma)
	if userEmailExist != nil && userEmailExist.ID != id {
		message := "Email already exist"
		return lib.ResponseError(lib.ResponseProps{Code: fiber.StatusConflict, Message: &message})
	}
	var hashPasswrd string
	if user.Password != "" {
		hashPasswrd, _ = lib.HashPassword(user.Password)
	} else {
		hashPasswrd = userExist.Password
	}
	dataUser := UserRequest{
		Name:     user.Name,
		Email:    user.Email,
		Password: hashPasswrd,
	}
	newUser := UpdateOne(id, dataUser, prisma)
	return lib.ResponseSuccess(lib.ResponseProps{Code: fiber.StatusOK, Data: newUser})
}

func DeleteOneService(id string, prisma *db.PrismaClient) lib.ResponseData {
	existUser := GetUserById(id, prisma)
	if existUser.ID == "" {
		return lib.ResponseError(lib.ResponseProps{Code: fiber.StatusNotFound})
	}
	userDeleted := DeleteOne(id, prisma)
	if userDeleted == nil {
		message := "Failed to delete user"
		return lib.ResponseError(lib.ResponseProps{Code: fiber.StatusInternalServerError, Message: &message})
	}
	return lib.ResponseSuccess(lib.ResponseProps{Code: fiber.StatusOK, Message: "User deleted"})
}

func validateStoreRequest(user UserRequest) []lib.ValidationResponse {
	// Define validation rules
	rules := lib.ValidationRules{
		"Name": func(value interface{}) bool {
			// Name must be a string and not empty
			name, ok := value.(string)
			return ok && name != ""
		},
		"Email": func(value interface{}) bool {
			// Email must be a string and not empty and must be a valid email
			email, ok := value.(string)
			return ok && email != "" && lib.ValidateEmail(email)
		},
		"Password": func(value interface{}) bool {
			// Password must be a string and not empty and must be at least 8 characters long
			password, ok := value.(string)
			return ok && password != "" && len(password) >= 8
		},
	}

	// Convert UserRequest to map
	userMap := map[string]interface{}{
		"Name":     user.Name,
		"Email":    user.Email,
		"Password": user.Password,
	}

	// Validate user input
	errors := lib.ValidateRequest(userMap, rules)

	return errors
}

func CheckUserExistAndDeletedAt(id string, prisma *db.PrismaClient) (*db.UserModel, *time.Time) {
	userExist := GetUserById(id, prisma)
	if userExist == nil {
		return nil, nil
	}
	deletedAtUserExist, _ := userExist.DeletedAt()
	var timeDeletedAt *time.Time
	if deletedAtUserExist.IsZero() == false {
		timeDeletedAt = &deletedAtUserExist
	}
	return userExist, timeDeletedAt
}
