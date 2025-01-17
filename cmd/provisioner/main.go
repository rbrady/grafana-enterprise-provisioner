package main

import (
	"flag"
	"log"
	"strings"

	"github.com/rbrady/grafana-enterprise-provisioner/internal/config"
	"github.com/rbrady/grafana-enterprise-provisioner/internal/services"
)

func main() {
	configFile := flag.String("config", "config.yaml", "path to config file")
	verbose := flag.Bool("verbose", false, "enable verbose logging")
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Create service with configuration
	tenantService := service.NewTenantService(service.ServiceConfig{
		BaseURL: "", // Base URL will be provided per tenant
	})

	for _, tenantCfg := range cfg.Tenants {
		err := tenantService.SyncTenant(
			tenantCfg.ServiceURL,
			tenantCfg.Tenant,
			strings.Split(tenantCfg.AccessPolicy, ","),
			tenantCfg.Token,
		)

		if err != nil {
			log.Printf("Error syncing tenant %s: %v", tenantCfg.Tenant, err)
			continue
		}

		if *verbose {
			log.Printf("Successfully synced tenant: %s", tenantCfg.Tenant)
		}
	}
}
