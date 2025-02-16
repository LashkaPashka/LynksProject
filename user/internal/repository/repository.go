package repository

import (
	"Lynks/user/configs"
	"Lynks/user/internal/model"
	"Lynks/user/pkg/jwt"
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	MongoClient *mongo.Client
	Token *jwt.JWT
}

const (
	databaseName = "UserDb"
	collectionName = "data"
)

var ctx = context.Background()

func NewUserRepository() *UserRepository {
	conf, _ := configs.LoadConfig()
	
	repo := &UserRepository{
		Token: jwt.NewJWT(conf.Secret),
	}

	mongoOpts := options.Client().ApplyURI(conf.DSN.DSN)
	client, err := mongo.Connect(ctx, mongoOpts)
	if err != nil {
		log.Fatal(err)
	}

	repo.MongoClient = client
	
	return repo
}

func (repo *UserRepository) InsertDocs(c *mongo.Client, email, password, name string) error {	
	// Verifying the user's existence
	_, isExist, _ := repo.GetByEmail(c, email)
	if !isExist {
		return errors.New("Пользователь существует!")
	}

	// Generate hashed_password
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// // Init model user
	user := &model.User{
		ID: int(uuid.New().ID()),
		Email: email,
		Name: name,
		Password: string(hashed_password),

	}

	defer c.Disconnect(ctx)

	//Added model user to Collection
	collection := c.Database(databaseName).Collection(collectionName)
	if _, err = collection.InsertOne(ctx, user); err != nil {
		return err
	}

	return nil
}

func (repo *UserRepository) GetByEmail(c *mongo.Client, email string) (*model.User, bool, error) {
	collection := c.Database(databaseName).Collection(collectionName)
	filter := bson.M{"email": email}
	
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, false,  err
	}

	var user model.User
	for cur.Next(ctx) {
		if err := cur.Decode(&user); err != nil {
			return nil, false, err
		}
	}

	isExist := user == model.User{}

	return &user, isExist,  nil
}