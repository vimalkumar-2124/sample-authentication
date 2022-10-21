package repositories

import (
	"context"
	"log"
	"time"

	"github.com/vimalkumar-2124/sample-authentication/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (u *UserRepo) GetUserById(id string) (bool, models.Users, error) {
	var users models.Users
	docId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, models.Users{}, err
	}
	filter := bson.M{"_id": docId}
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

// Used MongoDB ID as a token
// func (u *UserRepo) SaveSession(session models.Session) (string, error) {
// 	insertResult, err := u.db.Collection("sessions").InsertOne(context.TODO(), session)
// 	if err != nil {
// 		return "", err
// 	}
// 	return insertResult.InsertedID.(primitive.ObjectID).Hex(), nil
// }

// Implemented JWT and storing it in sessions collection
func (u *UserRepo) SaveSession(session models.Session) error {
	_, err := u.db.Collection("sessions").InsertOne(context.TODO(), session)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepo) GetSessinById(token string) (bool, models.Session, error) {
	// docId, err := primitive.ObjectIDFromHex(token)
	// if err != nil {
	// 	return false, models.Session{}, err
	// }
	var session models.Session
	filter := bson.M{
		"token": token,
		"expiryAt": bson.M{
			"$gte": time.Now().Unix(),
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
	// docId, err := primitive.ObjectIDFromHex(authToken)
	// if err != nil {
	// 	return err
	// }
	filter := bson.M{"token": authToken}
	update := bson.M{"$set": bson.M{"expiryAt": time.Now()}}
	_, err := u.db.Collection("sessions").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepo) UpdateUser(user models.SignInBody) error {
	// docId, err := primitive.ObjectIDFromHex(authToken)
	// if err != nil {
	// 	return err
	// }
	filter := bson.M{"email": user.Email}
	update := bson.M{"$set": bson.M{"password": user.Password}}
	_, err := u.db.Collection("auth").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepo) DoesUserExist(email string) (bool, error) {
	exist, _, err := u.GetUserByEmail(email)
	if err != nil {
		return false, err
	}
	return exist, nil

}

func (u *UserRepo) ShowUsers() ([]models.AllUser, error) {
	log.Println("All user repo started...")
	findOptions := options.Find()
	var users []models.AllUser
	cur, err := u.db.Collection("auth").Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		return []models.AllUser{}, err
	}
	for cur.Next(context.TODO()) {
		var ele models.AllUser
		err := cur.Decode(&ele)
		if err != nil {
			return []models.AllUser{}, err
		}
		users = append(users, ele)
	}
	if err := cur.Err(); err != nil {
		return []models.AllUser{}, err
	}
	cur.Close(context.TODO())
	return users, nil
}
