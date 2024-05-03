package main

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure_np"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/session"
	"log"
)

func main() {
	// All clients require a Session. The Session provides the client with
	// shared configuration such as account and credentials.
	// A Session should be shared where possible to take advantage of
	// configuration and credential caching. See the session package for
	// more information.
	sess := session.New()

	// Create a new instance of the service's client with a Session.
	// Optional spotinst.Config values can also be provided as variadic
	// arguments to the New function. This option allows you to provide
	// service specific configuration.
	svc := azure_np.New(sess)

	// Create a new context.
	ctx := context.Background()

	// Read stateful node configuration.
	_, err := svc.LaunchNewNodes(ctx, &azure_np.LaunchNewNodesInput{
		Adjustment:          spotinst.Int(2),
		OceanId:             spotinst.String("o-123456"),
		ApplicableVmSizes:   []string{"standard_d2a_v4"},
		AvailabilityZones:   []string{"1", "2"},
		MinCoresPerNode:     spotinst.Int(2),
		MinMemoryGiBPerNode: spotinst.Float64(2),
		PreferredLifecycle:  spotinst.String("Spot"),
		VngIds:              []string{"vng-12345", "vng-78901"},
	})
	if err != nil {
		log.Fatalf("spotinst: failed to launch the nodes: %v", err)
	}
}
