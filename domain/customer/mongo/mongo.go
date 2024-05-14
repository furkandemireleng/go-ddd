package mongo

import (
	"context"
	"github.com/furkandemireleng/go-ddd/aggregate"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoRepository struct {
	db       *mongo.Database
	customer *mongo.Collection
}

// MongoCustomer is a internal type that is used to stroe a Customer Aggregate inside the repository

type MongoCustomer struct {
	ID   uuid.UUID `bson:"id"`
	Name string    `bson:"name"`
}

func NewFromCustomer(c aggregate.Customer) MongoCustomer {
	return MongoCustomer{
		ID:   c.GetId(),
		Name: c.GetName(),
	}
}

func (m MongoCustomer) ToAggregate() aggregate.Customer {
	c := aggregate.Customer{}

	c.SetID(m.ID)
	_ = c.SetName(m.Name)

	return c
}

func New(ctx context.Context, connectionString string) (*MongoRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}

	db := client.Database("test")
	customers := db.Collection("customers")

	return &MongoRepository{
		db:       db,
		customer: customers,
	}, nil

}

func (mr *MongoRepository) Get(uuid uuid.UUID) (aggregate.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result := mr.customer.FindOne(ctx, bson.M{"id": uuid})

	var customer MongoCustomer
	err := result.Decode(&customer)
	if err != nil {
		return aggregate.Customer{}, err
	}

	return customer.ToAggregate(), nil

}
func (mr *MongoRepository) Add(c aggregate.Customer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	internal := NewFromCustomer(c)

	_, err := mr.customer.InsertOne(ctx, internal)
	if err != nil {
		return err
	}
	return nil
}
func (mr *MongoRepository) Update(c aggregate.Customer) error {
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	panic("implement me")

}

func (mr *MongoRepository) Delete(uuid uuid.UUID) error {
	panic("implement me")

}
