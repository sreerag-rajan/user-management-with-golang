package repository

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/sreerag-rajan/user-management-with-go/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRoleRepository struct {
	collection *mongo.Collection
}

func NewUserRoleRepository(db *mongo.Database) *UserRoleRepository {
	collection := db.Collection("userRoles")
	return &UserRoleRepository{
		collection: collection,
	}
}

func (r *UserRoleRepository) Create(userRole *model.RolePayload) (*model.UserRoleModel, error) {
	fmt.Println("ENTERING REPOSITORY")
	result, err := r.collection.InsertOne(context.Background(), userRole)
	if err != nil {
		log.Println("Failed to create a new Role", err)
		return nil, errors.New("failed to create new role")
	}

	createdUserRole := &model.UserRoleModel{
		Id:    result.InsertedID.(primitive.ObjectID),
		Label: userRole.Label,
		Code:  userRole.Code,
	}
	return createdUserRole, nil
}
