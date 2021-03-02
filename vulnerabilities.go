package citadel

import (
	"context"
	"log"

	vulnerabilities "github.com/brittonhayes/citadel/gen/vulnerabilities"
)

// vulnerabilities service example implementation.
// The example methods log the requests and return zero values.
type vulnerabilitiessrvc struct {
	logger *log.Logger
}

// NewVulnerabilities returns the vulnerabilities service implementation.
func NewVulnerabilities(logger *log.Logger) vulnerabilities.Service {
	return &vulnerabilitiessrvc{logger}
}

// Find implements find.
func (s *vulnerabilitiessrvc) Find(ctx context.Context, p *vulnerabilities.FindPayload) (res *vulnerabilities.Vulnerability, err error) {
	res = &vulnerabilities.Vulnerability{}
	s.logger.Print("vulnerabilities.find")
	return
}

// List all of the vulnerabilities
func (s *vulnerabilitiessrvc) List(ctx context.Context, p *vulnerabilities.LimitPayload) (res []*vulnerabilities.Vulnerability, err error) {
	s.logger.Print("vulnerabilities.list")
	return
}

// Submit implements submit.
func (s *vulnerabilitiessrvc) Submit(ctx context.Context, p *vulnerabilities.SubmitPayload) (err error) {
	s.logger.Print("vulnerabilities.submit")
	return
}
