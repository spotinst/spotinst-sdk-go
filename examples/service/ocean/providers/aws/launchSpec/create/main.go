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

	// Create a new launch spec.
	out, err := svc.CloudProviderAWS().CreateLaunchSpec(ctx, &aws.CreateLaunchSpecInput{
		LaunchSpec: &aws.LaunchSpec{
			Name:             spotinst.String("test-vng"),
			OceanID:          spotinst.String("o-abcd123456"),
			ImageID:          spotinst.String("ami-abcd123456"),
			SecurityGroupIDs: []string{"sg-123456"},
			SubnetIDs:        []string{"subnet-12345", "subnet-67890"},
			Strategy: &aws.LaunchSpecStrategy{
				DrainingTimeout:          spotinst.Int(500),
				UtilizeCommitments:       spotinst.Bool(true),
				UtilizeReservedInstances: spotinst.Bool(true),
			},
			InstanceTypesFilters: &aws.InstanceTypesFilters{
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
			BlockDeviceMappings: []*aws.BlockDeviceMapping{
				&aws.BlockDeviceMapping{
					DeviceName: spotinst.String("/dev/xvdb"),
					EBS: &aws.EBS{
						DeleteOnTermination: spotinst.Bool(true),
						Encrypted:           spotinst.Bool(false),
						IOPS:                spotinst.Int(100),
						Throughput:          spotinst.Int(125),
						VolumeType:          spotinst.String("gp3"),
						SnapshotID:          spotinst.String("snap-12345"),
						VolumeSize:          spotinst.Int(35),
					},
				}},
			EphemeralStorage: &aws.EphemeralStorage{
				DeviceName: spotinst.String("/dev/xvdb"),
			}},
	})
	if err != nil {
		log.Fatalf("spotinst: failed to create launch spec: %v", err)
	}

	// Output.
	if out.LaunchSpec != nil {
		log.Printf("LaunchSpec %q: %s",
			spotinst.StringValue(out.LaunchSpec.ID),
			stringutil.Stringify(out.LaunchSpec))
	}
}
