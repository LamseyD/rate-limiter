package mongo

import (
	"context"
	"ratelimiter/pkg/models"
	"ratelimiter/pkg/static"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

// TODO this is probably not necessary if we use another repo for the worker for loading the rules into cache
type Repository interface {
	AddUser(ctx context.Context, user models.User) error
}

type repository struct {
	client *mongo.Client
	db     *mongo.Database
	logger *zap.Logger
}

func NewRepository(conf Config, logger *zap.Logger) (Repository, error) {

	uri := conf.GetConnectionString()
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Ping the MongoDB server to check if the connection is successful
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	db := client.Database(conf.Database)

	return &repository{
		client: client,
		db:     db,
		logger: logger,
	}, nil
}

// AddUser implements Repository.
func (r *repository) AddUser(ctx context.Context, user models.User) error {
	collection := r.db.Collection(static.MONGO_REPO_USERS_COLLECTION)
	// Insert the person document into the collection
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}

	return nil
}
