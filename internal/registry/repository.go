package registry

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	RegisterService(ctx context.Context, service *Service) error
	GettServices(ctx context.Context) ([]*Service, error)
	DeregisterService(ctx context.Context, id string) error
	PingDatabase(ctx context.Context) error
}

type MongoRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(db *mongo.Database, collectionName string) Repository {
	return &MongoRepository{
		collection: db.Collection(collectionName),
	}
}

func (r *MongoRepository) RegisterService(ctx context.Context, service *Service) error {
	service.RegistredAt = time.Now()
	_, err := r.collection.InsertOne(ctx, service)

	return err
}

func (r *MongoRepository) GettServices(ctx context.Context) ([]*Service, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var services []*Service
	if err := cursor.All(ctx, &services); err != nil {
		return nil, err
	}

	return services, nil
}

func (r *MongoRepository) DeregisterService(ctx context.Context, id string) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})

	return err
}

func (r *MongoRepository) PingDatabase(ctx context.Context) error {
	return r.collection.Database().Client().Ping(ctx, nil)
}
