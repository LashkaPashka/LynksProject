package repository

import (
	"Lynks/user/configs"
	"Lynks/user/internal/model"
	"context"
	"log"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	MongoClient *mongo.Client
}

const (
	databaseName = "UserDb"
	collectionName = "data"
)

var ctx = context.Background()

func NewUserRepository() *UserRepository {
	conf, _ := configs.LoadConfig()
	
	repo := &UserRepository{}

	mongoOpts := options.Client().ApplyURI(conf.DSN.DSN)
	client, err := mongo.Connect(ctx, mongoOpts)
	if err != nil {
		log.Fatal(err)
	}

	repo.MongoClient = client

	return repo
}

func (repo *UserRepository) InsertDocs(c *mongo.Client, email, password, name string) error {	
	
	// Generate hashed_password
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Init model user
	user := &model.User{
		ID: int(uuid.New().ID()),
		Email: email,
		Name: name,
		Password: string(hashed_password),

	}

	defer c.Disconnect(ctx)

	//Added model user to Collection
	collection := c.Database(databaseName).Collection(collectionName)
	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}

	return nil
}
