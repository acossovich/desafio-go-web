package routes

import (
	"github.com/bootcamp-go/desafio-go-web/cmd/handler"
	"github.com/bootcamp-go/desafio-go-web/internal/domain"

	"github.com/bootcamp-go/desafio-go-web/internal/tickets"	
	"github.com/gin-gonic/gin"
)

type Router struct{
	Engine *gin.Engine
	Handler *handler.Service
}

func NewRouter(engine *gin.Engine,db []domain.Ticket) (router *Router){
	// Setup components.

	repository := tickets.NewRepository(db)

	service := tickets.NewService(repository)

	handler := handler.NewService(service)

	router = &Router{
		Engine: engine,
		Handler: handler,
	}

	return
}

func (router *Router) MapRouter(){

	group := router.Engine.Group("/ticket")
	{
		group.GET("/getByCountry/:dest", router.Handler.GetTicketsByCountry())
		group.GET("/getAverage/:dest", router.Handler.AverageDestination())
	}

}

