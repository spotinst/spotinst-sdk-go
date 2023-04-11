package main

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/service/ocean"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
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
	svc := ocean.New(sess)

	// Create a new context.
	ctx := context.Background()

	// Create a new ECS cluster.
	out, err := svc.CloudProviderAWS().CreateECSCluster(ctx, &aws.CreateECSClusterInput{
		Cluster: &aws.ECSCluster{
			Name:        spotinst.String("foo"),
			Region:      spotinst.String("us-west-2"),
			ClusterName: spotinst.String("foo"),
			Capacity: &aws.ECSCapacity{
				Target: spotinst.Int(5),
			},
			Strategy: &aws.ECSStrategy{
				SpotPercentage: spotinst.Int(100),
				ClusterOrientation: &aws.ECSClusterOrientation{
					AvailabilityVsCost: spotinst.String("balanced"),
				},
			},
			AutoScaler: &aws.ECSAutoScaler{
				IsEnabled:    spotinst.Bool(false),
				IsAutoConfig: spotinst.Bool(false),
				Cooldown:     spotinst.Int(300),

				Down: &aws.ECSAutoScalerDown{
					MaxScaleDownPercentage: spotinst.Float64(20),
				},
				ShouldScaleDownNonServiceTasks:   spotinst.Bool(true),
				EnableAutomaticAndManualHeadroom: spotinst.Bool(true),

				ResourceLimits: &aws.ECSAutoScalerResourceLimits{
					MaxVCPU:      spotinst.Int(1024),
					MaxMemoryGiB: spotinst.Int(20),
				},
			},
			Compute: &aws.ECSCompute{
				InstanceTypes: &aws.ECSInstanceTypes{
					Filters: &aws.ECSFilters{
						Architectures:         []string{"x86_64", "i386"},
						DiskTypes:             []string{"EBS", "SSD"},
						MinVcpu:               spotinst.Int(2),
						MaxVcpu:               spotinst.Int(80),
						MinGpu:                spotinst.Int(0),
						MaxGpu:                spotinst.Int(5),
						IncludeFamilies:       []string{"c*", "t*"},
						ExcludeFamilies:       []string{"m*"},
						ExcludeMetal:          spotinst.Bool(true),
						Categories:            []string{"General_purpose", "Compute_optimized"},
						Hypervisor:            []string{"nitro", "xen"},
						IsEnaSupported:        spotinst.Bool(true),
						MaxMemoryGiB:          spotinst.Float64(16),
						MinMemoryGiB:          spotinst.Float64(8),
						MinEnis:               spotinst.Int(2),
						VirtualizationTypes:   []string{"hvm"},
						RootDeviceTypes:       []string{"ebs"},
						MaxNetworkPerformance: spotinst.Int(20),
						MinNetworkPerformance: spotinst.Int(2),
					},
				},
				SubnetIDs: []string{
					"subnet-12345",
					"subnet-67890",
				},
				LaunchSpecification: &aws.ECSLaunchSpecification{
					ImageID:          spotinst.String("ami-12345"),
					Monitoring:       spotinst.Bool(false),
					SecurityGroupIDs: []string{"sg-12345"},
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
