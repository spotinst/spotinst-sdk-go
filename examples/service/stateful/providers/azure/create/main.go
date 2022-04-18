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
	out, err := svc.Create(ctx, &azure.CreateStatefulNodeInput{
		StatefulNode: &azure.StatefulNode{
			Name:              spotinst.String("foo"),
			Region:            spotinst.String("eastus"),
			ResourceGroupName: spotinst.String("foo"),
			Description:       spotinst.String("foo"),
			Strategy: &azure.Strategy{
				Signals: []*azure.Signal{
					{
						Type:    spotinst.String("vmReady"),
						Timeout: spotinst.Int(20),
					},
				},
				FallbackToOnDemand: spotinst.Bool(true),
				DrainingTimeout:    spotinst.Int(12),
				RevertToSpot: &azure.RevertToSpot{
					PerformAt: spotinst.String("timeWindow"),
				},
				OptimizationWindows: []string{
					"Tue:19:46-Tue:20:46",
				},
			},
			Compute: &azure.Compute{
				OS: spotinst.String("Linux"),
				VMSizes: &azure.VMSizes{
					SpotSizes: []string{
						"standard_ds1_v2",
						"standard_ds2_v2",
						"standard_ds3_v2",
						"standard_ds4_v2",
					},
					OnDemandSizes: []string{
						"standard_ds1_v2",
						"standard_ds2_v2",
					},
				},
				Zones: []string{
					"1",
					"2",
				},
				PreferredZone: spotinst.String("2"),
				LaunchSpecification: &azure.LaunchSpecification{
					Image: &azure.Image{
						MarketPlace: &azure.MarketPlaceImage{
							Publisher: spotinst.String("Canonical"),
							Offer:     spotinst.String("UbuntuServer"),
							SKU:       spotinst.String("18.04-LTS"),
							Version:   spotinst.String("latest"),
						},
					},
					Network: &azure.Network{
						ResourceGroupName:  spotinst.String("foo"),
						VirtualNetworkName: spotinst.String("foo"),
						NetworkInterfaces: []*azure.NetworkInterface{
							{
								IsPrimary:      spotinst.Bool(true),
								SubnetName:     spotinst.String("default"),
								AssignPublicIP: spotinst.Bool(true),
								PublicIPSku:    spotinst.String("Standard"),
								NetworkSecurityGroup: &azure.NetworkSecurityGroup{
									ResourceGroupName: spotinst.String("foo"),
									Name:              spotinst.String("foo"),
								},
							},
						},
					},
					DataDisks: []*azure.DataDisk{
						{
							SizeGB: spotinst.Int(1),
							LUN:    spotinst.Int(1),
							Type:   spotinst.String("Standard_LRS"),
						},
					},
					Extensions: []*azure.Extension{
						{
							Name:                    spotinst.String("foo"),
							Type:                    spotinst.String("customScript"),
							Publisher:               spotinst.String("Microsoft.Azure.Extensions"),
							APIVersion:              spotinst.String("2.0"),
							MinorVersionAutoUpgrade: spotinst.Bool(true),
							ProtectedSettings: map[string]interface{}{
								"script": "foo"},
						},
					},
					Login: &azure.Login{
						UserName: spotinst.String("foo"),
						Password: spotinst.String("bar"),
					},
					Tags: []*azure.Tag{
						{
							TagKey:   spotinst.String("Creator"),
							TagValue: spotinst.String("Tamiry@netapp.com"),
						},
					},
				},
			},
			Scheduling: &azure.Scheduling{
				Tasks: []*azure.Task{
					{
						IsEnabled:      spotinst.Bool(true),
						Type:           spotinst.String("pause"),
						CronExpression: spotinst.String("44 10 * * *"),
					},
					{
						IsEnabled:      spotinst.Bool(true),
						Type:           spotinst.String("resume"),
						CronExpression: spotinst.String("48 10 * * *"),
					},
					{
						IsEnabled:      spotinst.Bool(true),
						Type:           spotinst.String("recycle"),
						CronExpression: spotinst.String("52 10 * * *"),
					},
				},
			},
			Persistence: &azure.Persistence{
				ShouldPersistDataDisks:   spotinst.Bool(true),
				OSDiskPersistenceMode:    spotinst.String("reattach"),
				ShouldPersistNetwork:     spotinst.Bool(true),
				DataDisksPersistenceMode: spotinst.String("reattach"),
				ShouldPersistOSDisk:      spotinst.Bool(true),
			},
			Health: &azure.Health{
				HealthCheckTypes: []string{
					"vmState",
				},
				UnhealthyDuration: spotinst.Int(300),
				GracePeriod:       spotinst.Int(180),
				AutoHealing:       spotinst.Bool(true),
			},
		},
	})
	if err != nil {
		log.Fatalf("spotinst: failed to update stateful node: %v", err)
	}

	// Output.
	if out.StatefulNode != nil {
		log.Printf("StatefulNode %q: %s",
			spotinst.StringValue(out.StatefulNode.ID),
			stringutil.Stringify(out.StatefulNode))
	}
}
