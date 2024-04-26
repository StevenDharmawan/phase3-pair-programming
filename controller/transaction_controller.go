package controller

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"pair-programming/helper"
	"pair-programming/model/web"
	"pair-programming/service"
)

type TransactionController interface {
	CreateTransaction(c echo.Context) error
	GetTransactionById(c echo.Context) error
	GetTransactions(c echo.Context) error
	UpdateTransaction(c echo.Context) error
	DeleteTransaction(c echo.Context) error
	DeleteFirstTransaction(c echo.Context) error
}

type TransactionControllerImpl struct {
	service.TransactionService
}

func NewTransactionController(transactionService service.TransactionService) *TransactionControllerImpl {
	return &TransactionControllerImpl{TransactionService: transactionService}
}

func (controller *TransactionControllerImpl) CreateTransaction(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (controller *TransactionControllerImpl) GetTransactionById(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (controller *TransactionControllerImpl) GetTransactions(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (controller *TransactionControllerImpl) UpdateTransaction(c echo.Context) error {
	transactionID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		errResponse := helper.ErrBadRequest(err)
		return c.JSON(errResponse.Code, errResponse)
	}
	var request web.TransactionRequest
	err = c.Bind(&request)
	if err != nil {
		errResponse := helper.ErrBadRequest(err.Error())
		return c.JSON(errResponse.Code, errResponse)
	}
	errResponse := controller.TransactionService.UpdateTransaction(request, transactionID)
	if errResponse != nil {
		return c.JSON(errResponse.Code, *errResponse)
	}
	return c.JSON(http.StatusCreated, web.Response{
		Message: "success update transaction",
	})
}

func (controller *TransactionControllerImpl) DeleteTransaction(c echo.Context) error {
	transactionID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		errResponse := helper.ErrBadRequest(err)
		return c.JSON(errResponse.Code, errResponse)
	}
	errResponse := controller.TransactionService.DeleteTransaction(transactionID)
	if errResponse != nil {
		return c.JSON(errResponse.Code, *errResponse)
	}
	return c.JSON(http.StatusCreated, web.Response{
		Message: "success delete transaction",
	})
}

func (controller *TransactionControllerImpl) DeleteFirstTransaction(c echo.Context) error {
	errResponse := controller.TransactionService.DeleteFirstTransaction()
	if errResponse != nil {
		return c.JSON(errResponse.Code, *errResponse)
	}
	return c.JSON(http.StatusCreated, web.Response{
		Message: "success delete first transaction",
	})
}
