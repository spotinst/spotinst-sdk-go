package main

import (
	"context"
	"log"

	"github.com/spotinst/spotinst-sdk-go/service/ocean"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/azure"
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
	out, err := svc.CloudProviderAzure().CreateCluster(ctx, &azure.CreateClusterInput{
		Cluster: &azure.Cluster{
			Name:                spotinst.String("foo"),
			ControllerClusterID: spotinst.String("foo"),
			AKS: &azure.AKS{
				Name:              spotinst.String("foo"),
				ResourceGroupName: spotinst.String("foo"),
			},
			Strategy: &azure.Strategy{
				SpotPercentage: spotinst.Int(100),
				FallbackToOD:   spotinst.Bool(true),
			},
			VirtualNodeGroupTemplate: &azure.VirtualNodeGroupTemplate{
				VMSizes: &azure.VMSizes{
					Whitelist: []string{"Standard_D2s_v3"},
				},
				LaunchSpecification: &azure.LaunchSpecification{
					ResourceGroupName: spotinst.String("foo"),
					Image: &azure.Image{
						MarketplaceImage: &azure.MarketplaceImage{
							Publisher: spotinst.String("Canonical"),
							Offer:     spotinst.String("UbuntuServer"),
							SKU:       spotinst.String("18.04-LTS"),
							Version:   spotinst.String("latest"),
						},
					},
					Network: &azure.Network{
						ResourceGroupName:  spotinst.String("foo"),
						VirtualNetworkName: spotinst.String("foo"),
						NetworkInterfaces: []*azure.NetworkInterface{
							{
								SubnetName:     spotinst.String("foo"),
								AssignPublicIP: spotinst.Bool(true),
								IsPrimary:      spotinst.Bool(true),
							},
						},
					},
					OSDisk: &azure.OSDisk{
						SizeGB: spotinst.Int(50),
						Type:   spotinst.String("Standard_LRS"),
					},
					Extensions: []*azure.Extension{
						{
							APIVersion:              spotinst.String("2019-03-01"),
							Publisher:               spotinst.String("Microsoft.Azure.Extensions"),
							Type:                    spotinst.String("CustomScript"),
							Name:                    spotinst.String("config-app"),
							MinorVersionAutoUpgrade: spotinst.Bool(true),
							ProtectedSettings: &map[string]interface{}{
								"commandToExecute":   "<command-to-execute>",
								"script":             "<base64-script-to-execute>",
								"storageAccountName": "<storage-account-name>",
								"storageAccountKey":  "<storage-account-key>",
								"managedIdentity":    "<managed-identity-identifier>",
								"fileUris": []string{
									"https://...",
								},
							},
						},
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
