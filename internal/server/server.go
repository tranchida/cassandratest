package server

import (
	"cassandratest/internal/database"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	Router *gin.Engine
	DB     database.Database
}

// ErrorResponse represents an error response
// @Description Error response
// @name ErrorResponse
// @property error string
//
type ErrorResponse struct {
	Error string `json:"error"`
}

func NewServer() *Server {

	r := gin.Default()

	Server := &Server{
		Router: r,
	}

	r.GET("/hotels", Server.GetHotels)
	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return Server
}

// GetHotels godoc
// @Summary List hotels
// @Description get hotels
// @Tags hotels
// @Produce json
// @Success 200 {array} database.Hotel
// @Failure 500 {object} server.ErrorResponse
// @Router /hotels [get]
func (s *Server) GetHotels(c *gin.Context) {
	hotels, err := s.DB.GetHotels()
	if err != nil {
		log.Println("Error fetching hotels:", err)
		c.JSON(500, ErrorResponse{Error: "Failed to fetch hotels"})
		return
	}
	c.JSON(200, hotels)
}