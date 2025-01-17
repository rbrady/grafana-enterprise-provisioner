package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/rbrady/grafana-enterprise-provisioner/internal/domain/models"
	"io/ioutil"
	"net/http"
)

type tenantRepository struct {
	serviceURL string
}

func NewTenantRepository(serviceURL string) models.TenantRepository {
	return &tenantRepository{serviceURL: serviceURL}
}

func (r *tenantRepository) GetCurrentTenants() ([]string, error) {
	resp, err := http.Get(r.serviceURL)
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %v", err)
	}
	defer resp.Body.Close()

	var response struct {
		Tenants []string `json:"tenants"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return response.Tenants, nil
}

func (r *tenantRepository) CreateTenant(tenant *models.Tenant) error {
	payload := map[string]interface{}{
		"tenant":        tenant.Name,
		"access_policy": tenant.AccessPolicy,
		"token":         tenant.Token,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error creating JSON payload: %v", err)
	}

	resp, err := http.Post(r.serviceURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error making POST request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("error response from server: %s (status: %d)", string(body), resp.StatusCode)
	}

	return nil
}
