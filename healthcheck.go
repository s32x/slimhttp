package slimhttp

// import (
// 	"net/http"

// 	"github.com/Sirupsen/logrus"
// 	"github.com/sdwolfe32/trumail/verifier"
// )

// // Lookuper defines all functionality for an email verification
// // lookup API service
// type Lookuper interface {
// 	Healthcheck(r *http.Request) (interface{}, error)
// }

// // lookuper contains all dependencies for a Lookuper
// type lookuper struct {
// 	log      *logrus.Entry
// 	hostname string
// }

// // NewLookuper generates a new email verification lookup API service
// func NewLookuper(log *logrus.Logger, hostname, sourceAddr string) Lookuper {
// 	return &lookuper{
// 		log:      log.WithField("service", "lookup"),
// 		hostname: hostname,
// 		ever:     verifier.NewVerifier(maxWorkerCount, hostname, sourceAddr),
// 	}
// }

// // healthcheck represents the response to a healthcheck request
// type healthcheck struct {
// 	Status   string `json:"status" xml:"status"`
// 	Hostname string `json:"hostname" xml:"hostname"`
// }

// // GetHealthcheck handles and returns a 200 and our hostname
// func (s *lookuper) Healthcheck(r *http.Request) (interface{}, error) {
// 	l := s.log.WithField("handler", "Healthcheck")
// 	l.Info("New Healthcheck request received")
// 	l.Info("Returning newly generated Healthcheck")
// 	return &healthcheck{Status: "OK", Hostname: s.hostname}, nil
// }
