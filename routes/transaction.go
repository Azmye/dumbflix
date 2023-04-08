package routes

import (
	"dumbflix/handlers"
	"dumbflix/pkg/mysql"
	"dumbflix/repositories"

	"github.com/labstack/echo/v4"
)

func TransactionRoutes(e *echo.Group) {
	TransactionRepository := repositories.RepositoryTransaction(mysql.DB)

	h := handlers.HandlerTransaction(TransactionRepository)

	e.GET("/transactions", h.FindTransactions)
	e.GET("/transaction/:id", h.GetTransaction)
	e.POST("/transaction", h.CreateTransaction)
	e.PATCH("/transaction", h.UpdateTransaction)
	e.DELETE("/transaction", h.DeleteTransaction)
}
