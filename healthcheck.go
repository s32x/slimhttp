package slimhttp

import (
	"net/http"

	"github.com/Sirupsen/logrus"
)

// Healthchecker defines all functionality for a healtcheck
// service
type Healthchecker interface {
	Healthcheck(r *http.Request) (interface{}, error)
}

// healthchecker contains all dependencies for a Healthchecker
type healthchecker struct {
	log      *logrus.Entry
	hostname string
}

// NewHealthchecker generates a new Healthchecker service
func NewHealthchecker(log *logrus.Logger, hostname string) Healthchecker {
	return &healthchecker{
		log:      log.WithField("service", "lookup"),
		hostname: hostname,
	}
}

// healthcheck represents the response to a healthcheck request
type healthcheck struct {
	Status   string `json:"status" xml:"status"`
	Hostname string `json:"hostname" xml:"hostname"`
}

// Healthcheck handles and returns a 200 as well as a fully populated
// and encoded healthcheck body
func (s *healthchecker) Healthcheck(r *http.Request) (interface{}, error) {
	l := s.log.WithField("handler", "Healthcheck")
	l.Debug("New Healthcheck request received")
	l.Debug("Returning newly generated Healthcheck")
	return &healthcheck{Status: "OK", Hostname: s.hostname}, nil
}
