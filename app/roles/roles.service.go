package roles

import (
	"goprisma/db"
	"goprisma/lib"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetAllRoles(query RoleQueryRequest, prisma *db.PrismaClient) lib.ResponseData {
	roles := FindManyRoles(query, prisma)
	if len(roles) == 0 {
		return lib.ResponseError(lib.ResponseProps{Code: fiber.StatusNotFound, Message: "Roles not found"})
	}
	return lib.ResponseSuccess(lib.ResponseProps{Code: fiber.StatusOK, Data: roles})
}

func GetSingleById(id string, prisma *db.PrismaClient) lib.ResponseData {
	roles := FindById(id, prisma)
	if roles == nil {
		return lib.ResponseError(lib.ResponseProps{Code: fiber.StatusNotFound, Message: "Role not found"})
	}
	var timeDeletedAt *time.Time
	deletedAt, ok := roles.DeletedAt()
	if ok {
		timeDeletedAt = &deletedAt
	}
	zeroTime := lib.IsZeroTime(deletedAt)
	if timeDeletedAt != nil && !zeroTime {
		return lib.ResponseError(lib.ResponseProps{Code: fiber.StatusNotFound, Message: "Role not found"})
	}
	roleResponse := RoleResponse{
		ID:        roles.ID,
		Name:      roles.Name,
		CreatedAt: roles.CreatedAt,
		UpdatedAt: roles.UpdatedAt,
	}
	return lib.ResponseSuccess(lib.ResponseProps{Code: fiber.StatusOK, Data: roleResponse})
}

func Store(role RoleRequest, prisma *db.PrismaClient) lib.ResponseData {
	// validate role input
	errors := validateStoreRequest(role)
	if len(errors) > 0 {
		return lib.ResponseError(lib.ResponseProps{Code: fiber.StatusBadRequest, Message: &errors})
	}
	// check if role already exist
	roleExist := FindByName(role.Name, prisma)
	if roleExist != nil {
		message := "Role already exist"
		return lib.ResponseError(lib.ResponseProps{Code: fiber.StatusConflict, Message: &message})
	}
	newRole := CreateOne(role, prisma)
	return lib.ResponseSuccess(lib.ResponseProps{Code: fiber.StatusOK, Data: newRole})
}

func Update(id string, role RoleRequest, prisma *db.PrismaClient) lib.ResponseData {
	// validate role input
	errors := validateStoreRequest(role)
	if len(errors) > 0 {
		return lib.ResponseError(lib.ResponseProps{Code: fiber.StatusBadRequest, Message: &errors})
	}
	// check if role exist
	roleExist := FindById(id, prisma)
	messageNotFound := "Role not found"
	if roleExist == nil {
		return lib.ResponseError(lib.ResponseProps{Code: fiber.StatusNotFound, Message: &messageNotFound})
	}
	var timeDeletedAt *time.Time
	deletedAt, ok := roleExist.DeletedAt()
	if ok {
		timeDeletedAt = &deletedAt
	}
	zeroTime := lib.IsZeroTime(deletedAt)
	if timeDeletedAt != nil && !zeroTime {
		return lib.ResponseError(lib.ResponseProps{Code: fiber.StatusNotFound, Message: &messageNotFound})
	}
	updatedRole := UpdateRole(id, role, prisma)
	return lib.ResponseSuccess(lib.ResponseProps{Code: fiber.StatusOK, Data: updatedRole})
}

func Delete(id string, prisma *db.PrismaClient) lib.ResponseData {
	roleExist := FindById(id, prisma)
	if roleExist == nil {
		message := "Role not found"
		return lib.ResponseError(lib.ResponseProps{Code: fiber.StatusNotFound, Message: &message})
	}
	deletedAt, _ := roleExist.DeletedAt()
	tDelete := lib.TimeDeletedAt(deletedAt)
	if tDelete != nil {
		message := "Failed. Record already deleted"
		return lib.ResponseError(lib.ResponseProps{Code: fiber.StatusBadRequest, Message: &message})
	}
	_ = DeleteRole(id, prisma)
	return lib.ResponseSuccess(lib.ResponseProps{Code: fiber.StatusOK, Message: "Role deleted successfully"})
}

func UndeleteRole(id string, p *db.PrismaClient) lib.ResponseData {
	existRoles := FindById(id, p)
	if existRoles == nil {
		return lib.ResponseError(lib.ResponseProps{Code: fiber.StatusNotFound, Message: "Role not found"})
	}
	deletedAt, _ := existRoles.DeletedAt()
	tDelete := lib.TimeDeletedAt(deletedAt)
	if tDelete == nil {
		message := "Failed. Record not deleted"
		return lib.ResponseError(lib.ResponseProps{Code: fiber.StatusBadRequest, Message: &message})
	}
	undeleteRole := RestoreRole(id, p)
	return lib.ResponseSuccess(lib.ResponseProps{Code: fiber.StatusOK, Data: undeleteRole})
}

func CheckRoleExistAndDeletedAt(id string, p *db.PrismaClient) (*db.RolesModel, *time.Time) {
	roleExist := FindById(id, p)
	if roleExist == nil {
		return nil, nil
	}
	deletedAt, _ := roleExist.DeletedAt()
	zTime := lib.IsZeroTime(deletedAt)
	if zTime {
		return roleExist, nil
	}
	return roleExist, &deletedAt
}

func validateStoreRequest(role RoleRequest) []lib.ValidationResponse {
	rules := lib.ValidationRules{
		"Name": func(value interface{}) bool {
			name, ok := value.(string)
			return ok && name != ""
		},
	}
	roleMaps := map[string]interface{}{
		"Name": role.Name,
	}
	errors := lib.ValidateRequest(roleMaps, rules)
	return errors
}
