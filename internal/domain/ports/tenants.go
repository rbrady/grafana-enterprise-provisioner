package ports

import (
	"github.com/rbrady/grafana-enterprise-provisioner/internal/domain/commands"
	"github.com/rbrady/grafana-enterprise-provisioner/internal/domain/models"
)

type TenantRepository interface {
	GetCurrentTenants() ([]string, error)
	CreateTenant(tenant *models.Tenant) error
}

type TenantHandler interface {
	Handle(cmd *commands.SyncTenantCommand) error
}
