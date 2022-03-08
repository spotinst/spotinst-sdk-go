package main

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/service/ocean"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure"
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
	out, err := svc.CloudProviderAzure().CreateVirtualNodeGroup(ctx, &azure.CreateVirtualNodeGroupInput{
		VirtualNodeGroup: &azure.VirtualNodeGroup{
			OceanID: spotinst.String("o-12345"),
			Name:    spotinst.String("foo"),
			ResourceLimits: &azure.VirtualNodeGroupResourceLimits{
				MaxInstanceCount: spotinst.Int(10),
			},
			Labels: []*azure.Label{
				{
					Key:   spotinst.String("label-key"),
					Value: spotinst.String("label-value"),
				},
			},
			Taints: []*azure.Taint{
				{
					Key:    spotinst.String("taint-key"),
					Value:  spotinst.String("taint-value"),
					Effect: spotinst.String("NoSchedule"),
				},
			},
			AutoScale: &azure.VirtualNodeGroupAutoScale{
				Headrooms: []*azure.VirtualNodeGroupHeadroom{
					{
						CPUPerUnit: spotinst.Int(20),
						NumOfUnits: spotinst.Int(5),
					},
					{
						MemoryPerUnit: spotinst.Int(70),
						NumOfUnits:    spotinst.Int(3),
					},
				},
				AutoHeadroomPercentage: spotinst.Int(50),
			},
			LaunchSpecification: &azure.VirtualNodeGroupLaunchSpecification{
				OSDisk: &azure.OSDisk{
					SizeGB: spotinst.Int(40),
					Type:   spotinst.String("Standard_LRS"),
				},
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
