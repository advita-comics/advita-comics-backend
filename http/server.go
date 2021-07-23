package http

import (
	"fmt"
	"net/http"

	"github.com/advita-comics/advita-comics-backend/db"
	"github.com/advita-comics/advita-comics-backend/http/handlers"

	"github.com/advita-comics/advita-comics-backend/config"
	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

// Server - сервер
type Server struct {
	HTTP     *echo.Echo
	Handlers *handlers.Handlers
}

// NewServer - конструктор
func NewServer(cfg *config.Config, db db.DB) *Server {
	e := echo.New()
	e.Server.Addr = fmt.Sprintf(":%d", cfg.HTTP.Port)

	return &Server{
		HTTP:     e,
		Handlers: handlers.NewHandlers(db),
	}
}

// Start - поднимает сервер
func (s *Server) Start() {
	log.Infof("Server starting %s ...", s.HTTP.Server.Addr)
	s.runRoutes()

	// Gracefully stop by signal
	if err := gracehttp.Serve(s.HTTP.Server); err != nil {
		log.Fatalf("Server filed: %s", err.Error())
	}
}

func (s *Server) runRoutes() {
	s.HTTP.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "ok")
	})

	s.HTTP.POST("/donation", s.Handlers.Donations.Donation)
	s.HTTP.GET("/company", s.Handlers.Company.CompanyInfo)
}
