package roles

import (
	"context"
	"fmt"
	"goprisma/db"
	"time"
)

type RoleQueryRequest struct {
	Name    string `json:"name"`
	Page    int    `json:"page"`
	PerPage int    `json:"per_page"`
}

type RoleResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RoleRequest struct {
	Name string `json:"name"`
}

func (r RoleResponse) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":         r.ID,
		"name":       r.Name,
		"created_at": r.CreatedAt,
		"updated_at": r.UpdatedAt,
	}
}

func (r RoleResponse) ToResponse() map[string]interface{} {
	return map[string]interface{}{
		"data": r.ToMap(),
	}
}

func FindManyRoles(query RoleQueryRequest, prisma *db.PrismaClient) []db.RolesModel {
	offset := (query.Page - 1) * query.PerPage
	roles, err := prisma.Roles.FindMany(
		db.Roles.Name.Contains(query.Name),
		db.Roles.Or(
			db.Roles.And(db.Roles.DeletedAt.IsNull()),
			db.Roles.And(db.Roles.DeletedAt.Equals(time.Time{})), // for zero time
		),
	).OrderBy(
		db.Roles.Name.Order(db.ASC),
	).Skip(offset).Take(query.PerPage).Exec(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	return roles
}

func FindById(id string, prisma *db.PrismaClient) *db.RolesModel {
	role, _ := prisma.Roles.FindUnique(
		db.Roles.ID.Equals(id),
	).Exec(context.Background())
	return role
}

func FindByName(name string, prisma *db.PrismaClient) *db.RolesModel {
	role, _ := prisma.Roles.FindFirst(
		db.Roles.Name.Equals(name),
	).Exec(context.Background())
	return role
}

func CreateOne(role RoleRequest, prisma *db.PrismaClient) *db.RolesModel {
	newRole, _ := prisma.Roles.CreateOne(
		db.Roles.Name.Set(role.Name),
	).Exec(context.Background())
	return newRole
}

func UpdateRole(id string, role RoleRequest, prisma *db.PrismaClient) *db.RolesModel {
	updatedRole, _ := prisma.Roles.FindUnique(
		db.Roles.ID.Equals(id),
	).Update(
		db.Roles.Name.Set(role.Name),
	).Exec(context.Background())
	return updatedRole
}

func DeleteRole(id string, prisma *db.PrismaClient) *db.RolesModel {
	deletedRole, _ := prisma.Roles.FindUnique(
		db.Roles.ID.Equals(id),
	).Update(db.Roles.DeletedAt.Set(time.Now())).Exec(context.Background())
	return deletedRole
}

func RestoreRole(id string, prisma *db.PrismaClient) *db.RolesModel {
	restoredRole, _ := prisma.Roles.FindUnique(
		db.Roles.ID.Equals(id),
	).Update(db.Roles.DeletedAt.Set(time.Time{})).Exec(context.Background())
	return restoredRole
}
