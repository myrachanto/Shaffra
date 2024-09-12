package users

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Userrepository repository
const (
	collectionName = "user"
)

var (
	Userrepository UserrepoInterface = &userrepository{}
	ctx                              = context.TODO()
	Userrepo                         = userrepository{}
)

type Key struct {
	EncryptionKey string `mapstructure:"EncryptionKey"`
}

type UserrepoInterface interface {
	Create(user *User) (*User, error)
	GetOne(code string) (user *User, errors error)
	GetAll() ([]User, error)
	Update(code string, user *User) (*User, error)
	Delete(code string) (string, error)
}
type userrepository struct {
	db *mongo.Database
}

func NewUserRepo(db *mongo.Database) UserrepoInterface {
	return &userrepository{
		db: db,
	}
}

func (r *userrepository) Create(user *User) (*User, error) {
	if err1 := user.Validate(); err1 != nil {
		return nil, err1
	}
	ok := r.emailexist(user.Email)
	if ok {

		return nil, fmt.Errorf("that email exist in the our system")
	}
	collection := r.db.Collection(collectionName)
	result1, errd := collection.InsertOne(ctx, &user)
	if errd != nil {
		return nil, fmt.Errorf("invalid input")
	}
	user.ID = result1.InsertedID.(primitive.ObjectID)
	return user, nil

}

func (r *userrepository) GetOne(code string) (user *User, errors error) {
	if len(code) == 0 {
		return nil, fmt.Errorf("the code is empty")
	}
	// Convert the string ID to a MongoDB ObjectID
	objectID, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format")
	}
	collection := r.db.Collection(collectionName)
	filter := bson.M{"_id": objectID}

	err = collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (r *userrepository) GetAll() ([]User, error) {
	collection := r.db.Collection(collectionName)
	results := []User{}
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("no records found")
	}
	if err = cursor.All(ctx, &results); err != nil {
		return nil, fmt.Errorf("error decoding")
	}
	return results, nil

}

func (r *userrepository) Update(code string, user *User) (*User, error) {
	if len(code) == 0 {
		return nil, fmt.Errorf("the code is empty")
	}
	// Convert the string ID to a MongoDB ObjectID
	objectID, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format")
	}
	collection := r.db.Collection(collectionName)
	filter := bson.M{"_id": objectID}
	uuser := &User{}
	err = collection.FindOne(ctx, filter).Decode(&uuser)
	if err != nil {
		return nil, fmt.Errorf("invalid input")
	}

	if user.Name == "" {
		user.Name = uuser.Name
	}
	if user.Email == "" {
		user.Email = uuser.Email
	}
	if user.Age == 0 {
		user.Age = uuser.Age
	}
	update := bson.M{"$set": user}
	_, errs := collection.UpdateOne(ctx, filter, update)
	if errs != nil {
		return nil, fmt.Errorf("Error updating!")
	}
	return user, nil
}

func (r userrepository) Delete(code string) (string, error) {
	if len(code) == 0 {
		return "", fmt.Errorf("the code is empty")
	}

	// Convert the string ID to a MongoDB ObjectID
	objectID, err := primitive.ObjectIDFromHex(code)
	if err != nil {
		return "", fmt.Errorf("invalid ID format")
	}
	collection := r.db.Collection(collectionName)
	filter := bson.M{"_id": objectID}
	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		return "", fmt.Errorf("invalid input")
	}
	return "deleted successfully", nil

}
func (r userrepository) emailexist(email string) bool {
	if len(email) == 0 {
		return false
	}
	collection := r.db.Collection(collectionName)
	result := &User{}
	filter := bson.M{"email": email}
	err1 := collection.FindOne(ctx, filter).Decode(&result)
	return err1 == nil
}
