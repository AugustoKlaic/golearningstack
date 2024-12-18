package repository

import (
	"context"
	. "github.com/AugustoKlaic/golearningstack/pkg/configuration"
	"github.com/AugustoKlaic/golearningstack/pkg/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type UserCredentialsRepository struct {
	collection *mongo.Collection
	timeout    time.Duration
}

func NewUserCredentialsRepository(db *mongo.Database) *UserCredentialsRepository {
	return &UserCredentialsRepository{
		collection: db.Collection(Props.Mongo.Dbname),
		timeout:    5 * time.Second,
	}
}

func (r *UserCredentialsRepository) contextWithTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), r.timeout)
}

func (r *UserCredentialsRepository) Create(entity *entity.UserCredentials) (*mongo.InsertOneResult, error) {
	ctx, cancel := r.contextWithTimeout()
	defer cancel()

	return r.collection.InsertOne(ctx, entity)
}

func (r *UserCredentialsRepository) FindByID(id string) (*entity.UserCredentials, error) {
	ctx, cancel := r.contextWithTimeout()
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(id)
	var userCredentials entity.UserCredentials
	err := r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&userCredentials)
	if err != nil {
		return nil, err
	}
	return &userCredentials, nil
}

func (r *UserCredentialsRepository) Update(id string, entity *entity.UserCredentials) (*mongo.UpdateResult, error) {
	ctx, cancel := r.contextWithTimeout()
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(id)
	update := bson.M{"$set": entity}
	return r.collection.UpdateByID(ctx, objID, update)
}

func (r *UserCredentialsRepository) Delete(id string) (*mongo.DeleteResult, error) {
	ctx, cancel := r.contextWithTimeout()
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(id)
	return r.collection.DeleteOne(ctx, bson.M{"_id": objID})
}
