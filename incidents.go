package citadel

import (
	"context"
	"log"
	"time"

	"github.com/brittonhayes/citadel/gen/incidents"
)

// incidents service example implementation.
// The example methods log the requests and return zero values.
type incidentssrvc struct {
	logger *log.Logger
}

// NewIncidents returns the incidents service implementation.
func NewIncidents(logger *log.Logger) incidents.Service {
	return &incidentssrvc{logger}
}

// Find implements find.
func (s *incidentssrvc) Find(ctx context.Context, p *incidents.FindPayload) (res *incidents.Incident, err error) {
	var (
		title = "hello"
		date  = time.Now().String()
	)
	res = &incidents.Incident{
		Title: &title,
		Date:  &date,
	}
	s.logger.Print("incidents.find")
	return
}

// ListAll implements list all.
func (s *incidentssrvc) ListAll(ctx context.Context, p *incidents.LimitPayload) (res []*incidents.Incident, err error) {
	s.logger.Print("incidents.list all")
	return
}
