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

	// Create a new cluster.
	out, err := svc.CloudProviderAWS().CreateCluster(ctx, &aws.CreateClusterInput{
		Cluster: &aws.Cluster{
			Name:                spotinst.String("foo"),
			Region:              spotinst.String("us-west-2"),
			ControllerClusterID: spotinst.String("foo"),
			Capacity: &aws.Capacity{
				Target: spotinst.Int(5),
			},
			Strategy: &aws.Strategy{
				SpotPercentage:     spotinst.Float64(100),
				FallbackToOnDemand: spotinst.Bool(true),
				ClusterOrientation: &aws.ClusterOrientation{
					AvailabilityVsCost: spotinst.String("cheapest"),
				},
				SpreadNodesBy: spotinst.String("vcpu"),
			},
			Compute: &aws.Compute{
				InstanceTypes: &aws.InstanceTypes{
					Filters: &aws.Filters{
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
				LaunchSpecification: &aws.LaunchSpecification{
					ImageID:          spotinst.String("ami-12345"),
					Monitoring:       spotinst.Bool(false),
					SecurityGroupIDs: []string{"sg-foo"},
					HealthCheckUnhealthyDurationBeforeReplacement: spotinst.Int(180),
					BlockDeviceMappings: []*aws.ClusterBlockDeviceMappings{
						&aws.ClusterBlockDeviceMappings{
							DeviceName: spotinst.String("/dev/sdf"),
							EBS: &aws.ClusterEBS{
								DeleteOnTermination: spotinst.Bool(true),
								Encrypted:           spotinst.Bool(false),
								IOPS:                spotinst.Int(100),
								Throughput:          spotinst.Int(125),
								VolumeType:          spotinst.String("gp3"),
								SnapshotID:          spotinst.String("snap-12345"),
								VolumeSize:          spotinst.Int(35),
								DynamicVolumeSize: &aws.ClusterDynamicVolumeSize{
									BaseSize:            spotinst.Int(20),
									Resource:            spotinst.String("CPU"),
									SizePerResourceUnit: spotinst.Int(2),
								},
							},
						},
					},
					AssociateIPv6Address: spotinst.Bool(true),
					ResourceTagSpecification: &aws.ResourceTagSpecification{
						Volumes: &aws.Volumes{
							ShouldTag: spotinst.Bool(true),
						},
					},
				},
			},
			Scheduling: &aws.Scheduling{
				Tasks: []*aws.Task{
					{
						IsEnabled: spotinst.Bool(true),
						Parameter: &aws.Parameter{
							AmiAutoUpdate: &aws.AmiAutoUpdate{
								ApplyRoll: spotinst.Bool(false),
								AmiAutoUpdateClusterRoll: &aws.AmiAutoUpdateClusterRoll{
									BatchMinHealthyPercentage: spotinst.Int(50),
									BatchSizePercentage:       spotinst.Int(1),
									Comment:                   spotinst.String("Test Description"),
									RespectPdb:                spotinst.Bool(false),
								},
								MinorVersion: spotinst.Bool(true),
								Patch:        spotinst.Bool(true),
							},
						},
						Type: spotinst.String("amiAutoUpdate"),
					},
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
