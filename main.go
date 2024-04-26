package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/robfig/cron/v3"
	"log"
	"net/http"
	"pair-programming/config"
	"pair-programming/controller"
	"pair-programming/repository"
	"pair-programming/service"
)

func main() {
	ctx := context.Background()
	client := config.ConnectMongoDB(ctx)
	collection := config.GetCollection(client, "transactions")
	transactionRepository := repository.NewTransactionRepository(ctx, collection)
	transactionService := service.NewTransactionService(transactionRepository)
	transactionController := controller.NewTransactionController(transactionService)
	e := echo.New()
	v1 := e.Group("/api/v1")
	{
		v1.POST("/transactions", transactionController.CreateTransaction)
		v1.GET("/transactions", transactionController.GetTransactions)
		v1.GET("/transactions/:id", transactionController.GetTransactionById)
		v1.PUT("/transactions/:id", transactionController.UpdateTransaction)
		v1.DELETE("/transactions/:id", transactionController.DeleteTransaction)

		v1.GET("/cronjob", transactionController.DeleteFirstTransaction)
	}
	c := cron.New()

	_, err := c.AddFunc("0 0 * * *", func() {
		resp, err := http.Get("http://localhost:8080/api/v1/cronjob")
		if err != nil {
			log.Println("Gagal menjalankan cron job:", err)
			return
		}
		defer resp.Body.Close()
		log.Println("Status response:", resp.Status)
	})
	if err != nil {
		log.Fatal("Gagal menambahkan cron job:", err)
	}
	c.Start()

	e.Logger.Fatal(e.Start(":8080"))
}
