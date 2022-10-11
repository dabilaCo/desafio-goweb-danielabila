package handler

import (
	"desafio-goweb-danielabila/internal/tickets"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Service struct {
	service tickets.Service
}

func NewService(s tickets.Service) *Service {
	return &Service{
		service: s,
	}

}

func (s *Service) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		tick, err := s.service.GetAll(c)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}
		c.JSON(200, tick)
	}
}

func (s *Service) GetTicketsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		tick, err := s.service.GetTicketByDestination(c, destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, tick)
	}
}

func (s *Service) AverageDestination() gin.HandlerFunc {
	return func(c *gin.Context) {

		destination := c.Param("dest")

		avg, err := s.service.GetAveragePerCountry(c, destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}

		c.JSON(200, avg)
	}
}
