package models

type Tenant struct {
	Name         string
	AccessPolicy []string
	Token        string
}

type TenantRepository interface {
	GetCurrentTenants() ([]string, error)
	CreateTenant(tenant *Tenant) error
}

type TenantService interface {
	SyncTenant(tenant *Tenant) error
}
