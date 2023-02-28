package main

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/service/ocean"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure_np"
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
	out, err := svc.CloudProviderAzureNP().CreateVirtualNodeGroup(ctx, &azure_np.CreateVirtualNodeGroupInput{
		VirtualNodeGroup: &azure_np.VirtualNodeGroup{
			OceanID: spotinst.String("foo"),
			Name:    spotinst.String("foo"),
			AvailabilityZones: []string{"1",
				"2"},
			Labels: &map[string]string{
				"Key":   "label-key",
				"Value": "label-value",
			},
			Taints: []*azure_np.Taint{
				{
					Key:    spotinst.String("taint-key"),
					Value:  spotinst.String("taint-value"),
					Effect: spotinst.String("NoSchedule"),
				},
			},
			NodePoolProperties: &azure_np.NodePoolProperties{
				MaxPodsPerNode:     spotinst.Int(100),
				EnableNodePublicIP: spotinst.Bool(false),
				OsDiskSizeGB:       spotinst.Int(128),
				OsDiskType:         spotinst.String("Managed"),
				OsType:             spotinst.String("Linux"),
			},
			NodeCountLimits: &azure_np.NodeCountLimits{
				MinCount: spotinst.Int(0),
				MaxCount: spotinst.Int(20),
			},
			Strategy: &azure_np.Strategy{
				SpotPercentage: spotinst.Int(100),
				FallbackToOD:   spotinst.Bool(true),
			},
		},
	})
	if err != nil {
		log.Fatalf("spotinst: failed to create virtual node group: %v", err)
	}

	// Output.
	if out.VirtualNodeGroup != nil {
		log.Printf("Virtual Node Group %q: %s",
			spotinst.StringValue(out.VirtualNodeGroup.ID),
			stringutil.Stringify(out.VirtualNodeGroup))
	}
}
