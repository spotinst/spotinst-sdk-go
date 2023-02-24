package main

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure_np"
	"log"

	"github.com/spotinst/spotinst-sdk-go/service/ocean"
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
	out, err := svc.CloudProviderAzureNP().CreateCluster(ctx, &azure_np.CreateClusterInput{
		Cluster: &azure_np.Cluster{
			Name:                spotinst.String("AKS2.0-Cluster"),
			ControllerClusterID: spotinst.String("AKS2.0-Cluster"),
			AKS: &azure_np.AKS{
				ClusterName:                     spotinst.String("node-pool-test"),
				ResourceGroupName:               spotinst.String("AutomationResourceGroup"),
				Region:                          spotinst.String("eastus"),
				InfrastructureResourceGroupName: spotinst.String("MC_AutomationResourceGroup_node-pool-test_eastus"),
			},
			AutoScaler: &azure_np.AutoScaler{
				IsEnabled: spotinst.Bool(true),
				ResourceLimits: &azure_np.ResourceLimits{
					MaxVCPU:      spotinst.Int(120),
					MaxMemoryGib: spotinst.Int(120),
				},
				Down: &azure_np.Down{
					MaxScaleDownPercentage: spotinst.Float64(30),
				},
				Headroom: &azure_np.Headroom{
					Automatic: &azure_np.Automatic{
						IsEnabled:  spotinst.Bool(true),
						Percentage: spotinst.Int(10),
					},
				},
			},
			Health: &azure_np.Health{
				GracePeriod: spotinst.Int(300),
			},
			VirtualNodeGroupTemplate: &azure_np.VirtualNodeGroupTemplate{
				NodePoolProperties: &azure_np.NodePoolProperties{
					MaxPodsPerNode:     spotinst.Int(100),
					EnableNodePublicIP: spotinst.Bool(false),
					OsDiskSizeGB:       spotinst.Int(128),
					OsDiskType:         spotinst.String("Managed"),
					OsType:             spotinst.String("Linux"),
				},
				NodeCountLimits: &azure_np.NodeCountLimits{
					MinCount: spotinst.Int(0),
					MaxCount: spotinst.Int(100),
				},
				Strategy: &azure_np.Strategy{
					SpotPercentage: spotinst.Int(100),
					FallbackToOD:   spotinst.Bool(true),
				},
				/*Taints:[]&azure.Taint{
					Key: spotinst.String("test"),
					Value: spotinst.String("new"),
					Effect: spotinst.String("Running"),
				},*/

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
