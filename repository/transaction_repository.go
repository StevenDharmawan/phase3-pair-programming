package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"pair-programming/model"
)

type TransactionRepository interface {
	CreateTransaction()
	GetTransactionById()
	GetTransactions()
	UpdateTransaction(transaction model.Transaction, transactionID primitive.ObjectID) error
	DeleteTransaction(transactionID primitive.ObjectID) error
	DeleteFirstTransaction() error
}

type TransactionRepositoryImpl struct {
	ctx        context.Context
	collection *mongo.Collection
}

func NewTransactionRepository(ctx context.Context, collection *mongo.Collection) *TransactionRepositoryImpl {
	return &TransactionRepositoryImpl{
		ctx:        ctx,
		collection: collection,
	}
}

func (repository *TransactionRepositoryImpl) CreateTransaction() {
	//TODO implement me
	panic("implement me")
}

func (repository *TransactionRepositoryImpl) GetTransactionById() {
	//TODO implement me
	panic("implement me")
}

func (repository *TransactionRepositoryImpl) GetTransactions() {
	//TODO implement me
	panic("implement me")
}

func (repository *TransactionRepositoryImpl) UpdateTransaction(transaction model.Transaction, transactionID primitive.ObjectID) error {
	filter := bson.M{"_id": transactionID}
	update := bson.M{
		"$set": transaction,
	}
	_, err := repository.collection.UpdateOne(context.Background(), filter, update)
	return err
}

func (repository *TransactionRepositoryImpl) DeleteTransaction(transactionID primitive.ObjectID) error {
	_, err := repository.collection.DeleteOne(repository.ctx, bson.M{"ID": transactionID})
	return err
}

func (repository *TransactionRepositoryImpl) DeleteFirstTransaction() error {
	_, err := repository.collection.DeleteOne(repository.ctx, bson.M{})
	return err
}
