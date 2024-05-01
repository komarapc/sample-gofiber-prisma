package has_roles

import (
	"context"
	"goprisma/db"
	"time"
)

type HasRolesRequest struct {
	UserId string `json:"user_id"`
	RoleId string `json:"role_id"`
}

type rolesInResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type userInResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
type HasRolesWithRelations struct {
	ID        string          `json:"id"`
	RoleId    string          `json:"role_id"`
	UserId    string          `json:"user_id"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	Role      rolesInResponse `json:"role"`
	User      userInResponse  `json:"user"`
}

type HasRolesWithRole struct {
	ID        string          `json:"id"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	Roles     rolesInResponse `json:"role"`
}

func FindById(id string, prisma *db.PrismaClient) *db.HasRolesModel {
	hasRole, _ := prisma.HasRoles.FindUnique(
		db.HasRoles.ID.Equals(id),
	).Select(
		db.HasRoles.ID.Field(),
		db.HasRoles.CreatedAt.Field(),
		db.HasRoles.UpdatedAt.Field(),
		db.HasRoles.DeletedAt.Field(),
	).With(
		db.HasRoles.Role.Fetch().Select(
			db.Roles.ID.Field(),
			db.Roles.Name.Field(),
		),
		db.HasRoles.User.Fetch().Select(
			db.User.ID.Field(),
			db.User.Name.Field(),
			db.User.Email.Field(),
		),
	).Exec(context.Background())
	return hasRole
}

func FindByUserId(userId string, prisma *db.PrismaClient) []HasRolesWithRole {
	hasRoles, _ := prisma.HasRoles.FindMany(db.HasRoles.UserID.Equals(userId)).Select(
		db.HasRoles.ID.Field(),
		db.HasRoles.CreatedAt.Field(),
		db.HasRoles.UpdatedAt.Field(),
	).With(
		db.HasRoles.Role.Fetch().Select(
			db.Roles.ID.Field(),
			db.Roles.Name.Field(),
		),
	).Exec(context.Background())
	var response []HasRolesWithRole
	for _, hasRole := range hasRoles {
		response = append(response, HasRolesWithRole{
			ID:        hasRole.ID,
			CreatedAt: hasRole.CreatedAt,
			UpdatedAt: hasRole.UpdatedAt,
			Roles: rolesInResponse{
				ID:   hasRole.Role().ID,
				Name: hasRole.Role().Name,
			},
		})
	}
	return response
}

func FindByUserIdAndRoleId(userId string, roleId string, prisma *db.PrismaClient) *db.HasRolesModel {
	hasRole, _ := prisma.HasRoles.FindFirst(
		db.HasRoles.UserID.Equals(userId),
		db.HasRoles.RoleID.Equals(roleId),
	).Select(
		db.HasRoles.ID.Field(),
		db.HasRoles.CreatedAt.Field(),
		db.HasRoles.UpdatedAt.Field(),
	).With(
		db.HasRoles.Role.Fetch().Select(
			db.Roles.ID.Field(),
			db.Roles.Name.Field(),
		),
	).Exec(context.Background())
	return hasRole
}

func CreateHasRoles(hasRolesRequest HasRolesRequest, prisma *db.PrismaClient) HasRolesWithRelations {
	hasRole, _ := prisma.HasRoles.CreateOne(
		db.HasRoles.User.Link(
			db.User.ID.Equals(hasRolesRequest.UserId),
		),
		db.HasRoles.Role.Link(
			db.Roles.ID.Equals(hasRolesRequest.RoleId),
		),
	).With(
		db.HasRoles.Role.Fetch().Select(
			db.Roles.ID.Field(),
			db.Roles.Name.Field(),
		),
		db.HasRoles.User.Fetch().Select(
			db.User.ID.Field(),
			db.User.Name.Field(),
			db.User.Email.Field(),
		),
	).Exec(context.Background())
	return HasRolesWithRelations{
		ID:        hasRole.ID,
		RoleId:    hasRole.RoleID,
		UserId:    hasRole.UserID,
		CreatedAt: hasRole.CreatedAt,
		UpdatedAt: hasRole.UpdatedAt,
		Role: rolesInResponse{
			ID:   hasRole.Role().ID,
			Name: hasRole.Role().Name,
		},
		User: userInResponse{
			ID:    hasRole.User().ID,
			Name:  hasRole.User().Name,
			Email: hasRole.User().Email,
		},
	}
}

func UpdateHasRoles(id string, data HasRolesRequest, prisma *db.PrismaClient) HasRolesWithRelations {
	hasRoles, _ := prisma.HasRoles.FindUnique(
		db.HasRoles.ID.Equals(id),
	).With(
		db.HasRoles.Role.Fetch().Select(
			db.Roles.ID.Field(),
			db.Roles.Name.Field(),
		),
		db.HasRoles.User.Fetch().Select(
			db.User.ID.Field(),
			db.User.Name.Field(),
			db.User.Email.Field(),
		),
	).Update(
		db.HasRoles.User.Link(
			db.User.ID.Equals(data.UserId),
		),
		db.HasRoles.Role.Link(
			db.Roles.ID.Equals(data.RoleId),
		),
	).Exec(context.Background())

	return HasRolesWithRelations{
		ID:        hasRoles.ID,
		RoleId:    hasRoles.RoleID,
		UserId:    hasRoles.UserID,
		CreatedAt: hasRoles.CreatedAt,
		UpdatedAt: hasRoles.UpdatedAt,
		Role: rolesInResponse{
			ID:   hasRoles.Role().ID,
			Name: hasRoles.Role().Name,
		},
		User: userInResponse{
			ID:    hasRoles.User().ID,
			Name:  hasRoles.User().Name,
			Email: hasRoles.User().Email,
		},
	}
}

func DestroyHasRoles(id string, prisma *db.PrismaClient) *db.HasRolesModel {
	hasRole, _ := prisma.HasRoles.FindUnique(
		db.HasRoles.ID.Equals(id),
	).Delete().Exec(context.Background())
	return hasRole
}
