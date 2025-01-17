package commands

type SyncTenantCommand struct {
	ServiceURL   string
	Name         string
	AccessPolicy []string
	Token        string
}
