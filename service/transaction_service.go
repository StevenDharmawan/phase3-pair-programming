package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"pair-programming/helper"
	"pair-programming/model"
	"pair-programming/model/web"
	"pair-programming/repository"
)

type TransactionService interface {
	CreateTransaction()
	GetTransactionById()
	GetTransactions()
	UpdateTransaction(transaction web.TransactionRequest, transactionID primitive.ObjectID) *web.ErrorResponse
	DeleteTransaction(transactionID primitive.ObjectID) *web.ErrorResponse
	DeleteFirstTransaction() *web.ErrorResponse
}

type TransactionServiceImpl struct {
	repository.TransactionRepository
}

func NewTransactionService(transactionRepository repository.TransactionRepository) *TransactionServiceImpl {
	return &TransactionServiceImpl{TransactionRepository: transactionRepository}
}

func (service *TransactionServiceImpl) CreateTransaction() {
	//TODO implement me
	panic("implement me")
}

func (service *TransactionServiceImpl) GetTransactionById() {
	//TODO implement me
	panic("implement me")
}

func (service *TransactionServiceImpl) GetTransactions() {
	//TODO implement me
	panic("implement me")
}

func (service *TransactionServiceImpl) UpdateTransaction(request web.TransactionRequest, transactionID primitive.ObjectID) *web.ErrorResponse {
	transaction := model.Transaction{
		Description: request.Description,
		Amount:      request.Amount,
	}
	err := service.TransactionRepository.UpdateTransaction(transaction, transactionID)
	if err != nil {
		errResponse := helper.ErrInternalServer(err.Error())
		return &errResponse
	}
	return nil
}

func (service *TransactionServiceImpl) DeleteTransaction(transactionID primitive.ObjectID) *web.ErrorResponse {
	err := service.TransactionRepository.DeleteTransaction(transactionID)
	if err != nil {
		errResponse := helper.ErrInternalServer(err.Error())
		return &errResponse
	}
	return nil
}

func (service *TransactionServiceImpl) DeleteFirstTransaction() *web.ErrorResponse {
	err := service.TransactionRepository.DeleteFirstTransaction()
	if err != nil {
		errResponse := helper.ErrInternalServer(err.Error())
		return &errResponse
	}
	return nil
}
