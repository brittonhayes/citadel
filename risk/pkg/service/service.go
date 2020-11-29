package service

import (
	"context"
	risk "github.com/brittonhayes/citadel/risk/pkg/models"
)

// RiskService describes the service.
type RiskService interface {
	// Add your methods here
	List(ctx context.Context) (response []risk.Risk, err error)
	Submit(ctx context.Context, risk risk.Risk) (response string, err error)
	Update(ctx context.Context, risk risk.Risk) (response string, err error)
}
