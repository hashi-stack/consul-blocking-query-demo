// blocking_query.go
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

	var lastIndex uint64
	log.Println("Starting blocking query watcher...")

	for {
		log.Printf("[Query] Waiting for service updates, index %d...", lastIndex)
		services, meta, err := client.Health().Service(common.ServiceName, "", false, &api.QueryOptions{
			WaitIndex: lastIndex,
			WaitTime:  10 * time.Minute,
		})
		if err != nil {
			log.Printf("[Query] Error querying service: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		if meta.LastIndex != lastIndex {
			fmt.Printf("[Query] Got %d instances at index %d\n", len(services), meta.LastIndex)
			lastIndex = meta.LastIndex
		}
	}
}
