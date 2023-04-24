package tickets

import (
	"context"
	"github.com/bootcamp-go/desafio-go-web/internal/domain"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Ticket, error)
	GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error)
	GetTotalTickets(ctx context.Context, destination string) (int, error)
	AverageDestination(ctx context.Context, destination string) (float64, error)
}

type DefaultService struct{
	Storage Repository
}

func NewService(repository Repository) (service Service){
	return &DefaultService{
		Storage: repository,
	}
}

func (s *DefaultService)GetAll(ctx context.Context) (list []domain.Ticket, err error){
	list,err = s.Storage.GetAll(ctx)

	return
}
func (s *DefaultService)GetTicketByDestination(ctx context.Context, destination string) (list []domain.Ticket, err error){
	list,err = s.Storage.GetTicketByDestination(ctx,destination)

	return	
}

func (s *DefaultService) GetTotalTickets(ctx context.Context, destination string) (total int, err error){
	
	totalList,err := s.GetTicketByDestination(ctx,destination)

	if err != nil {
		return
	}
	
	total = len(totalList)

	return

}

func (s *DefaultService) AverageDestination(ctx context.Context, destination string) (average float64, err error){
	listByCountry,err := s.GetTicketByDestination(ctx,destination)

	if err != nil {
		return
	}
	
	totalList,err := s.GetAll(ctx)

	if err != nil{
		return
	}

	average = float64(len(listByCountry)) / float64(len(totalList))

	return	
}