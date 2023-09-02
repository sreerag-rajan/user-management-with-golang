package repository

import (
	"context"
	"errors"
	"log"

	"github.com/sreerag-rajan/user-management-with-go/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	collection   *mongo.Collection
	userRoleRepo *UserRoleRepository
}

func NewUserRepository(db *mongo.Database, userRoleRepo *UserRoleRepository) *UserRepository {
	collection := db.Collection("user")
	return &UserRepository{
		collection:   collection,
		userRoleRepo: userRoleRepo,
	}
}

func (r *UserRepository) Create(user *model.User) (*model.User, error) {

	result, err := r.collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Println("failed to create a new user", err)
		return nil, errors.New("failed to create a new user")
	}

	createdUser := &model.User{
		Id:        result.InsertedID.(primitive.ObjectID),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
		Role:      user.Role,
	}

	return createdUser, nil
}

func (r *UserRepository) Update(user *model.User) (*string, error) {
	update := bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{Key: "firstName", Value: user.FirstName},
				{Key: "lastName", Value: user.LastName},
				{Key: "email", Value: user.Email},
				{Key: "role", Value: user.Role},
				{Key: "password", Value: user.Password},
			},
		},
	}

	_, err := r.collection.UpdateOne(context.Background(), bson.M{"_id": user.Id}, update)
	if err != nil {
		log.Println("failed to update the user", err)
		return nil, errors.New("failed to update the user")
	}

	successMessage := "Successfully Updated"

	return &successMessage, nil
}

func (r *UserRepository) UpdateById(userID primitive.ObjectID, updateData bson.M) error {
	update := bson.D{{Key: "$set", Value: updateData}}

	_, err := r.collection.UpdateOne(context.Background(), bson.M{"_id": userID}, update)
	if err != nil {
		log.Println("failed to update the user by ID", err)
		return errors.New("failed to update the user by ID")
	}

	return nil
}

func (r *UserRepository) Delete(userId primitive.ObjectID) (*string, error) {
	_, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": userId})

	if err != nil {
		log.Println("Failed to delete the user", err)
		return nil, err
	}

	sucessMessage := "Successfully Deleted"
	return &sucessMessage, nil
}

func (r *UserRepository) Get(userId primitive.ObjectID) (*model.User, error) {
	var user model.User

	err := r.collection.FindOne(context.Background(), bson.M{"_id": userId}).Decode(&user)

	if err != nil {
		log.Println("Error in get User", err)
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetAll(page int64, pageSize int64) ([]*model.User, error) {
	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * pageSize)
	findOptions.SetLimit(pageSize)

	cursor, err := r.collection.Find(context.Background(), bson.M{}, findOptions)
	if err != nil {
		log.Println("ERROR IN GET ALL USERS", err)
		return nil, err
	}

	defer cursor.Close(context.Background())

	users := []*model.User{}
	for cursor.Next(context.Background()) {
		var user model.User
		err := cursor.Decode(&user)
		if err != nil {
			log.Println("ERROR IN DECODING USER", err)
			continue
		}
		users = append(users, &user)
	}
	if err := cursor.Err(); err != nil {
		log.Println("cursor error while fetching users", err)
		return nil, err
	}

	return users, nil
}
