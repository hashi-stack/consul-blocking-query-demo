// main.go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/hashi-stack/consul-blocking-query-demo/common"
	"github.com/hashicorp/consul/api"
)

func main() {
	client, err := common.CreateConsulClient()
	if err != nil {
		log.Fatalf("Failed to create Consul client: %v", err)
	}

	// get the service name from args
	serviceName := common.ServiceName
	serviceID := fmt.Sprintf("%s-1", serviceName)
	port := 8080

	log.Println("Starting service updater...")
	for {
		// Generate a unique nonce for the service registration
		nonce := time.Now().UnixNano()

		registration := &api.AgentServiceRegistration{
			ID:      serviceID,
			Name:    serviceName,
			Address: "fake-service",
			Port:    port,
			Check: &api.AgentServiceCheck{
				HTTP:     fmt.Sprintf("http://fake-service:%d/health", port),
				Interval: "5s",
				Timeout:  "1s",
			},
			Meta: map[string]string{
				"nonce":   fmt.Sprintf("%d", nonce),
				"version": "1.0.0",
			},
		}

		err := client.Agent().ServiceRegister(registration)
		if err != nil {
			log.Printf("[Register] Failed to register service: %v", err)
		} else {
			fmt.Println("[Register] Service registered/updated")
		}

		time.Sleep(10 * time.Second)
	}
}
