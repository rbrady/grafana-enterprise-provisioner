package service

import (
	"github.com/rbrady/grafana-enterprise-provisioner/internal/adapters/http"
	"github.com/rbrady/grafana-enterprise-provisioner/internal/domain/commands"
	"github.com/rbrady/grafana-enterprise-provisioner/internal/handlers"
)

type TenantService struct {
	handler *handlers.TenantHandler
}

type ServiceConfig struct {
	BaseURL string
}

func NewTenantService(config ServiceConfig) *TenantService {
	// Create dependencies inside the service
	repo := http.NewTenantRepository(config.BaseURL)
	handler := handlers.NewTenantHandler(repo)

	return &TenantService{
		handler: handler,
	}
}

func (s *TenantService) SyncTenant(serviceURL, name string, accessPolicy []string, token string) error {
	cmd := commands.SyncTenantCommand{
		ServiceURL:   serviceURL,
		Name:         name,
		AccessPolicy: accessPolicy,
		Token:        token,
	}

	return s.handler.Handle(cmd)
}
