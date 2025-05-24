// common/consul.go
package common

import (
	"github.com/hashicorp/consul/api"
)

// CreateConsulClient creates and returns a new Consul client
func CreateConsulClient() (*api.Client, error) {
	return api.NewClient(api.DefaultConfig())
}

// ServiceName is the common service name used by both components
const ServiceName = "my-service-1"
