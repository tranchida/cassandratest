package server

import (
	"cassandratest/internal/model"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lucavallin/gotel"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	Router    *gin.Engine
	DB        model.Repository
	Telemetry gotel.TelemetryProvider
}

// ErrorResponse represents an error response
// @Description Error response
// @name ErrorResponse
// @property error string
type ErrorResponse struct {
	Error string `json:"error"`
}

func NewServer(db model.Repository) *Server {

	r := gin.Default()

	Server := &Server{
		Router: r,
		DB:     db,
	}

	r.GET("/hotels", Server.GetHotels)
	r.GET("/hotel/:id", Server.GetHotel)
	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return Server
}

// GetHotels godoc
// @Summary List hotels
// @Description get hotels
// @Tags hotels
// @Produce json
// @Success 200 {array} model.Hotel
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

// GetHotel godoc
// @Summary Get hotel by id
// @Description get hotel by id
// @Tags hotels
// @Produce json
// @Param id path string true "Hotel ID"
// @Success 200 {object} model.Hotel
// @Failure 500 {object} server.ErrorResponse
// @Router /hotel/{id} [get]
func (s *Server) GetHotel(c *gin.Context) {
	hotelId := c.Param("id")
	hotel, err := s.DB.GetHotel(hotelId)
	if err != nil {
		log.Println("Error fetching hotel:", err)
		c.JSON(500, ErrorResponse{Error: "Failed to fetch hotel"})
		return
	}
	c.JSON(200, hotel)
}