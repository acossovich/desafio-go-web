package routes

import (
	"github.com/bootcamp-go/desafio-go-web/cmd/handler"
	"github.com/bootcamp-go/desafio-go-web/internal/domain"

	"github.com/bootcamp-go/desafio-go-web/internal/tickets"	
	"github.com/gin-gonic/gin"
)

func NewRouter(engine *gin.Engine,db []domain.Ticket) (*handler.Service){
	// Setup components.

	repository := tickets.NewRepository(db)

	service := tickets.NewService(repository)

	handler := handler.NewService(service)

	// Set routes.
	group := engine.Group("/ticket")
	{
		group.GET("/getByCountry/:dest", handler.GetTicketsByCountry())
		group.GET("/getAverage/:dest", handler.AverageDestination())
	}

	return handler
}