package main

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/service/stateful/providers/azure"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/session"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/stringutil"
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
	svc := azure.New(sess)

	// Create a new context.
	ctx := context.Background()

	// Read stateful node configuration.
	out, err := svc.ImportVM(ctx, &azure.ImportVMStatefulNodeInput{
		StatefulNodeImport: &azure.StatefulNodeImport{
			ResourceGroupName: spotinst.String("foo"),
			OriginalVMName:    spotinst.String("foo"),
			DrainingTimeout:   spotinst.Int(0),
			StatefulNode: &azure.StatefulNode{
				Name:              spotinst.String("foo"),
				Region:            spotinst.String("eastus"),
				ResourceGroupName: spotinst.String("foo"),
				Description:       spotinst.String("foo"),
				Compute: &azure.Compute{
					VMSizes: &azure.VMSizes{
						OnDemandSizes: []string{
							"standard_ds1_v2"},
						SpotSizes: []string{
							"standard_ds1_v2"},
					},
					LaunchSpecification: &azure.LaunchSpecification{
						Tags: []*azure.Tag{
							{
								TagKey:   spotinst.String("foo"),
								TagValue: spotinst.String("bar"),
							},
						},
					},
				},
				Health: &azure.Health{
					HealthCheckTypes: []string{
						"vmState"},
					GracePeriod:       spotinst.Int(60),
					AutoHealing:       spotinst.Bool(true),
					UnhealthyDuration: spotinst.Int(120),
				},
				Persistence: &azure.Persistence{
					ShouldPersistOSDisk:      spotinst.Bool(true),
					ShouldPersistDataDisks:   spotinst.Bool(true),
					ShouldPersistNetwork:     spotinst.Bool(false),
					OSDiskPersistenceMode:    spotinst.String("reattach"),
					DataDisksPersistenceMode: spotinst.String("onLaunch"),
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("spotinst: failed to update stateful node state: %v", err)
	}

	// Output.
	if out.StatefulNodeImport != nil {
		log.Printf("%s",
			stringutil.Stringify(out.StatefulNodeImport),
		)
	}
}
