package main

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/gcp"
	"log"

	"github.com/spotinst/spotinst-sdk-go/service/elastigroup"
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
	out, err := svc.CloudProviderGCP().Create(ctx, &gcp.CreateGroupInput{
		Group: &gcp.Group{
			Name:        spotinst.String("TF_GCP_EG"),
			Description: spotinst.String("terraform"),
			Capacity: &gcp.Capacity{
				Target:  spotinst.Int(1),
				Maximum: spotinst.Int(2),
				Minimum: spotinst.Int(0),
			},
			Strategy: &gcp.Strategy{
				FallbackToOnDemand:    spotinst.Bool(true),
				PreemptiblePercentage: spotinst.Int(100),
				OnDemandCount:         spotinst.Int(1),
				RevertToPreemptible: &gcp.RevertToPreemptible{
					PerformAt: spotinst.String("timeWindow"),
				},
				OptimizationWindows: []string{"Mon:01:00-Mon:03:00"},
			},
			Compute: &gcp.Compute{
				InstanceTypes: &gcp.InstanceTypes{
					OnDemand: spotinst.String("n1-standard-1"),
					Preemptible: []string{
						"n1-standard-1",
						"n1-standard-2",
					},
				},
				AvailabilityZones: []string{"us-central1-a", "us-central1-b", "us-central1-c"},
				Subnets: []*gcp.Subnet{
					{
						Region: spotinst.String("us-central1"),
						SubnetNames: []string{
							"default",
						},
					},
				},
				LaunchSpecification: &gcp.LaunchSpecification{
					InstanceNamePrefix: spotinst.String("terraform"),
					NetworkInterfaces: []*gcp.NetworkInterface{
						{
							Network:   spotinst.String("default"),
							ProjectID: nil,
						},
					},
					Disks: []*gcp.Disk{
						{
							InitializeParams: &gcp.InitializeParams{
								DiskSizeGB:  spotinst.Int(10),
								DiskType:    spotinst.String("pd-standard"),
								SourceImage: spotinst.String("https://www.googleapis.com/compute/v1/projects/test/global/images/docker-image-1"),
							},
							Mode:       spotinst.String("READ_WRITE"),
							Type:       spotinst.String("PERSISTENT"),
							AutoDelete: spotinst.Bool(true),
							Boot:       spotinst.Bool(true),
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
