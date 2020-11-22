package main

import (
	"context"
	"log"

	"github.com/spotinst/spotinst-sdk-go/service/elastigroup"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
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
	svc := elastigroup.New(sess)

	// Create a new context.
	ctx := context.Background()

	// Create a new group.
	out, err := svc.CloudProviderAWS().Create(ctx, &aws.CreateGroupInput{
		Group: &aws.Group{
			Name:        spotinst.String("foo"),
			Description: spotinst.String("bar"),
			Region:      spotinst.String("us-west-2"),
			Capacity: &aws.Capacity{
				Target: spotinst.Int(5),
			},
			Strategy: &aws.Strategy{
				Risk:               spotinst.Float64(100),
				FallbackToOnDemand: spotinst.Bool(true),
			},
			Compute: &aws.Compute{
				Product: spotinst.String(aws.ProductName[aws.ProductLinuxUnix]),
				InstanceTypes: &aws.InstanceTypes{
					OnDemand: spotinst.String("c3.large"),
					Spot: []string{
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
		log.Fatalf("spotinst: failed to create group: %v", err)
	}

	// Output.
	if out.Group != nil {
		log.Printf("Group %q: %s",
			spotinst.StringValue(out.Group.ID),
			stringutil.Stringify(out.Group))
	}
}
