package repositories

import (
	"context"
	"time"

	"github.com/vimalkumar-2124/sample-authentication/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	db *mongo.Database
}

func NewInstanceOfUserRepo(db *mongo.Database) UserRepo {
	return UserRepo{db: db}
}

func (u *UserRepo) SaveUser(user models.Users) error {
	_, err := u.db.Collection("auth").InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil

}

func (u *UserRepo) GetUserByEmail(email string) (bool, models.Users, error) {
	var users models.Users
	filter := bson.M{"email": email}
	count, err := u.db.Collection("auth").CountDocuments(context.TODO(), filter)
	if err != nil {
		return false, models.Users{}, err
	}
	if count != 1 {
		return false, models.Users{}, nil
	}
	err = u.db.Collection("auth").FindOne(context.TODO(), filter).Decode(&users)
	if err != nil {
		return false, models.Users{}, err
	}
	return true, users, nil
}

func (u *UserRepo) SaveSession(session models.Session) (string, error) {
	insertResult, err := u.db.Collection("sessions").InsertOne(context.TODO(), session)
	if err != nil {
		return "", err
	}
	return insertResult.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (u *UserRepo) GetSessinById(token string) (bool, models.Session, error) {
	docId, err := primitive.ObjectIDFromHex(token)
	if err != nil {
		return false, models.Session{}, err
	}
	var session models.Session
	filter := bson.M{
		"_id": docId,
		"expiryAt": bson.M{
			"$gte": time.Now(),
		},
	}
	count, err := u.db.Collection("sessions").CountDocuments(context.TODO(), filter)
	if err != nil {
		return false, models.Session{}, err
	}
	if count != 1 {
		return false, models.Session{}, nil
	}
	err = u.db.Collection("sessions").FindOne(context.TODO(), filter).Decode(&session)
	if err != nil {
		return false, models.Session{}, err
	}
	return true, session, nil

}

func (u *UserRepo) MarkSessionAsExpired(authToken string) error {
	docId, err := primitive.ObjectIDFromHex(authToken)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": docId}
	update := bson.M{"$set": bson.M{"expiryAt": time.Now()}}
	_, err = u.db.Collection("sessions").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}
