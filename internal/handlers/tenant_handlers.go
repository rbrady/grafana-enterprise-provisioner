package handlers

import (
	"github.com/rbrady/grafana-enterprise-provisioner/internal/domain/commands"
	"github.com/rbrady/grafana-enterprise-provisioner/internal/domain/models"
	"github.com/rbrady/grafana-enterprise-provisioner/internal/domain/ports"
)

type TenantHandler struct {
	repo ports.TenantRepository
}

func NewTenantHandler(repo ports.TenantRepository) *TenantHandler {
	return &TenantHandler{
		repo: repo,
	}
}

func (h *TenantHandler) Handle(cmd commands.SyncTenantCommand) error {
	currentTenants, err := h.repo.GetCurrentTenants()
	if err != nil {
		return err
	}

	exists := false
	for _, t := range currentTenants {
		if t == cmd.Name {
			exists = true
			break
		}
	}

	if !exists {
		tenant := &models.Tenant{
			Name:         cmd.Name,
			AccessPolicy: cmd.AccessPolicy,
			Token:        cmd.Token,
		}
		return h.repo.CreateTenant(tenant)
	}

	return nil
}
