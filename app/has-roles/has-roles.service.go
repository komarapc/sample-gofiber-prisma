package has_roles

import (
	"github.com/gofiber/fiber/v2"
	"goprisma/app/roles"
	"goprisma/app/users"
	"goprisma/db"
	"goprisma/lib"
)

func GetHasRoleById(id string, p *db.PrismaClient) lib.ResponseData {
	hasRole := FindById(id, p)
	if hasRole == nil {
		message := "Has role not found"
		return lib.ResponseError(lib.ResponseProps{Code: fiber.StatusNotFound, Message: message})
	}
	return lib.ResponseSuccess(lib.ResponseProps{Code: fiber.StatusOK, Data: hasRole})
}

func GetHasRolesByUserId(userId string, p *db.PrismaClient) lib.ResponseData {
	hasRoles := FindByUserId(userId, p)
	if hasRoles == nil {
		message := "Has roles not found"
		return lib.ResponseError(lib.ResponseProps{Code: fiber.StatusNotFound, Message: message})
	}
	return lib.ResponseSuccess(lib.ResponseProps{Code: fiber.StatusOK, Data: hasRoles})
}

func StoreHasRoles(req HasRolesRequest, p *db.PrismaClient) lib.ResponseData {
	userExist, userDeletedAt := users.CheckUserExistAndDeletedAt(req.UserId, p)
	if userExist == nil || userDeletedAt != nil {
		message := "User not found"
		return lib.ResponseError(lib.ResponseProps{Code: fiber.StatusNotFound, Message: message})
	}
	roleExist, roleDeletedAt := roles.CheckRoleExistAndDeletedAt(req.RoleId, p)
	if roleExist == nil || roleDeletedAt != nil {
		message := "Role not found"
		return lib.ResponseError(lib.ResponseProps{Code: fiber.StatusNotFound, Message: message})
	}

	existingHasRoles := FindByUserIdAndRoleId(req.UserId, req.RoleId, p)
	if existingHasRoles != nil {
		message := "User already has this role"
		return lib.ResponseError(lib.ResponseProps{Code: fiber.StatusConflict, Message: message})
	}
	hasRoles := CreateHasRoles(req, p)
	return lib.ResponseSuccess(lib.ResponseProps{Code: fiber.StatusCreated, Data: hasRoles})
}

func UpdateService(id string, req HasRolesRequest, p *db.PrismaClient) lib.ResponseData {
	existHasRoles := FindById(id, p)
	if existHasRoles == nil {
		message := "Has role not found"
		return lib.ResponseError(lib.ResponseProps{Code: fiber.StatusNotFound, Message: message})
	}
	duplicateHasRoles := FindByUserIdAndRoleId(req.UserId, req.RoleId, p)
	if duplicateHasRoles != nil {
		message := "User already has this role"
		return lib.ResponseError(lib.ResponseProps{Code: fiber.StatusConflict, Message: message})
	}
	updateHasRoles := UpdateHasRoles(id, req, p)
	return lib.ResponseSuccess(lib.ResponseProps{Code: fiber.StatusOK, Data: updateHasRoles})
}

func DestroyService(id string, p *db.PrismaClient) lib.ResponseData {
	hasRole := FindById(id, p)
	if hasRole == nil {
		message := "Data not found"
		return lib.ResponseError(lib.ResponseProps{Code: fiber.StatusNotFound, Message: message})
	}
	_ = DestroyHasRoles(id, p)
	return lib.ResponseSuccess(lib.ResponseProps{Code: fiber.StatusOK, Message: "Has role deleted successfully"})
}
