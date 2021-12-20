package main

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup"
	azurev3 "github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/azure/v3"
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
	svc := elastigroup.New(sess)

	// Create a new context.
	ctx := context.Background()

	// Create a new group.
	out, err := svc.CloudProviderAzureV3().Create(ctx, &azurev3.CreateGroupInput{
		Group: &azurev3.Group{
			Name:              spotinst.String("foo"),
			Region:            spotinst.String("foo"),
			ResourceGroupName: spotinst.String("foo"),
			Capacity: &azurev3.Capacity{
				Minimum: spotinst.Int(0),
				Target:  spotinst.Int(1),
				Maximum: spotinst.Int(2),
			},
			Strategy: &azurev3.Strategy{
				FallbackToOnDemand: spotinst.Bool(true),
				SpotPercentage:     spotinst.Int(90),
			},
			Compute: &azurev3.Compute{
				OS: spotinst.String("Linux"),
				VMSizes: &azurev3.VMSizes{
					SpotSizes:     []string{"standard_a1_v2"},
					OnDemandSizes: []string{"standard_a1_v2"},
				},
				LaunchSpecification: &azurev3.LaunchSpecification{
					Login: &azurev3.Login{
						UserName:     spotinst.String("foo"),
						SSHPublicKey: spotinst.String("foo"),
					},
					Image: &azurev3.Image{
						MarketPlace: &azurev3.MarketPlaceImage{
							Publisher: spotinst.String("Canonical"),
							SKU:       spotinst.String("18.04-LTS"),
							Version:   spotinst.String("latest"),
							Offer:     spotinst.String("UbuntuServer"),
						},
					},
					ShutdownScript: spotinst.String("foo"),
					Network: &azurev3.Network{
						ResourceGroupName:  spotinst.String("foo"),
						VirtualNetworkName: spotinst.String("foo"),
						NetworkInterfaces: []*azurev3.NetworkInterface{
							{
								SubnetName:     spotinst.String("foo"),
								AssignPublicIP: spotinst.Bool(false),
								IsPrimary:      spotinst.Bool(true),
							},
						},
					},
					LoadBalancersConfig: &azurev3.LoadBalancersConfig{
						LoadBalancers: []*azurev3.LoadBalancer{
							{
								Name:              spotinst.String("foo"),
								Type:              spotinst.String("loadBalancer"),
								ResourceGroupName: spotinst.String("foo"),
								SKU:               spotinst.String("Basic"),
								BackendPoolNames:  []string{"foo"},
							},
						},
					},
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("spotinst: failed to create group: %v", err)
	}

	// Output.
	if out.Group != nil {
		log.Printf("Group %q: %s",
			spotinst.StringValue(out.Group.ID),
			stringutil.Stringify(out.Group))
	}
}
