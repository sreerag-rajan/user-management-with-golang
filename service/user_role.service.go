package service

import (
	"log"

	"github.com/sreerag-rajan/user-management-with-go/config/db"
	"github.com/sreerag-rajan/user-management-with-go/model"
	"github.com/sreerag-rajan/user-management-with-go/repository"
)

// type UserRoleService struct {
// 	userRoleRepository *repository.UserRoleRepository
// }

func CreateUserRoleService(role *model.RolePayload) (*model.UserRoleModel, error) {

	repo := repository.NewUserRoleRepository(&db.DbConnection)

	newRole, err := repo.Create(role)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return newRole, nil
}
