package tickets

import (
	"context"
	"desafio-goweb-danielabila/internal/domain"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Ticket, error)
	GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error)
	GetAveragePerCountry(ctx context.Context, country string) (float64, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
			repository: r,
	}
}

func (s *service) GetAll(ctx context.Context) ([]domain.Ticket, error) {
	all, err := s.repository.GetAll(ctx)
	if err != nil{
		return nil, err
	}
	return all, nil
}

func (s *service) GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error){
	ticketDestine, err := s.repository.GetTicketByDestination(ctx, destination)
	if err != nil {
		return nil, err
	}
	return ticketDestine, nil
}

func (s *service) GetAveragePerCountry(ctx context.Context, country string) (float64, error){
	
	totalTickets, errTotal := s.repository.GetAll(ctx)
	if errTotal != nil {
		return 0, errTotal
	}

	tickets, errTicket := s.repository.GetTicketByDestination(ctx, country)
	if errTicket != nil {
		return 0, errTicket
	}

	return float64(len(tickets)) / float64(len(totalTickets)), nil
}