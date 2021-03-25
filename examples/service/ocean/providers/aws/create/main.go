package main

import (
	"context"
	"log"

	"github.com/spotinst/spotinst-sdk-go/service/ocean"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/session"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/stringutil"
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
	svc := ocean.New(sess)

	// Create a new context.
	ctx := context.Background()

	// Create a new cluster.
	out, err := svc.CloudProviderAWS().CreateCluster(ctx, &aws.CreateClusterInput{
		Cluster: &aws.Cluster{
			Name:   spotinst.String("foo"),
			Region: spotinst.String("us-west-2"),
			Capacity: &aws.Capacity{
				Target: spotinst.Int(5),
			},
			Strategy: &aws.Strategy{
				SpotPercentage:     spotinst.Float64(100),
				FallbackToOnDemand: spotinst.Bool(true),
			},
			Compute: &aws.Compute{
				InstanceTypes: &aws.InstanceTypes{
					Whitelist: []string{
						"c3.large",
						"c4.large",
					},
				},
				SubnetIDs: []string{
					"subnet-12345",
					"subnet-67890",
				},
				LaunchSpecification: &aws.LaunchSpecification{
					ImageID:          spotinst.String("ami-12345"),
					Monitoring:       spotinst.Bool(false),
					SecurityGroupIDs: []string{"sg-foo"},
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("spotinst: failed to create cluster: %v", err)
	}

	// Output.
	if out.Cluster != nil {
		log.Printf("Cluster %q: %s",
			spotinst.StringValue(out.Cluster.ID),
			stringutil.Stringify(out.Cluster))
	}
}
