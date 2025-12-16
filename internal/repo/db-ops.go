package repo

import (
	"chatapp/internal/model"

	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

type DBOps interface {
	Insert(context.Context, model.Employee) error
	List(context.Context) ([]model.Employee, error)
}

type dbOps struct {
	client mongo.Client
}

func New(client mongo.Client) DBOps {
	return &dbOps{
		client: client,
	}
}

// Insert inserts to collection
func (d dbOps) Insert(ctx context.Context, emp model.Employee) error {

	log.Printf("inserting employee data: %v", emp)
	res, err := d.client.Database("admin").Collection("employees").InsertOne(ctx, emp)
	if err != nil {
		log.Printf("insert failed: %s", err.Error())
		return err
	}
	log.Printf("insert resp: %v", res)

	return nil
}

func (d dbOps) List(ctx context.Context) ([]model.Employee, error) {
	return nil, nil
}
