package server

import (
	"cassandratest/internal/database"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Router *gin.Engine
	DB     database.Database
}

func NewServer() *Server {

	r := gin.Default()

	Server := &Server{
		Router: r,
	}

	r.GET("/hotels", Server.GetHotels)

	return Server
}

func (s *Server) GetHotels(c *gin.Context) {
	hotels, err := s.DB.GetHotels()
	if err != nil {
		log.Println("Error fetching hotels:", err)
		c.JSON(500, gin.H{"error": "Failed to fetch hotels"})
		return
	}
	c.JSON(200, hotels)
}