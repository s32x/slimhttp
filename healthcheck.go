package slimhttp

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// HealthcheckService defines all functionality for a healthcheckService
type HealthcheckService interface {
	Healthcheck(r *http.Request) (interface{}, error)
}

// healthcheckService contains all dependencies for a HealthcheckService
type healthcheckService struct {
	log      *logrus.Entry
	hostname string
}

// NewHealthcheckService generates a new Healthchecker service
func NewHealthcheckService(log *logrus.Logger, hostname string) HealthcheckService {
	return &healthcheckService{
		log:      log.WithField("service", "healthcheck"),
		hostname: hostname,
	}
}

// HealthcheckResponse is the response to a healthcheck request
type HealthcheckResponse struct {
	Status   string `json:"status" xml:"status"`
	Hostname string `json:"hostname" xml:"hostname"`
}

// Healthcheck handles and returns a healthcheck request, returning
// a 200 as well as a fully populated HealthcheckResponse
func (s *healthcheckService) Healthcheck(r *http.Request) (interface{}, error) {
	l := s.log.WithField("handler", "Healthcheck")
	l.Debug("New Healthcheck request received")
	l.Debug("Returning newly generated Healthcheck")
	return &HealthcheckResponse{Status: "OK", Hostname: s.hostname}, nil
}
