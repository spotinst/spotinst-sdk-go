package main

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/service/ocean"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/gcp"
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
	out, err := svc.CloudProviderGCP().CreateCluster(ctx, &gcp.CreateClusterInput{
		Cluster: &gcp.Cluster{
			Name:                spotinst.String("terraform-cluster-3"),
			ControllerClusterID: spotinst.String("terraform-cluster-3-a1de28dd"),
			Scheduling: &gcp.Scheduling{
				Tasks: []*gcp.Task{
					{
						IsEnabled:      spotinst.Bool(true),
						Type:           spotinst.String("clusterRoll"),
						CronExpression: spotinst.String("0 1 * * *"),
						Parameters: &gcp.Parameters{
							ClusterRoll: &gcp.ClusterRoll{
								BatchSizePercentage:       spotinst.Int(20),
								Comment:                   spotinst.String("reason for cluster roll"),
								BatchMinHealthyPercentage: spotinst.Int(100),
								RespectPdb:                spotinst.Bool(true),
							},
						},
					},
				},
			},
			Capacity: &gcp.Capacity{
				Minimum: spotinst.Int(0),
				Maximum: spotinst.Int(1000),
				Target:  spotinst.Int(1),
			},
			Compute: &gcp.Compute{
				SubnetName: spotinst.String("default"),
				InstanceTypes: &gcp.InstanceTypes{
					Whitelist: []string{"e2-micro"},
				},
				LaunchSpecification: &gcp.LaunchSpecification{
					ServiceAccount: spotinst.String("493916419393-compute@developer.gserviceaccount.com"),
					SourceImage:    spotinst.String("https://www.googleapis.com/compute/v1/projects/gke-node-images/global/images/gke-1249-gke3200-cos-97-16919-235-1-v230120-c-pre"),
					Tags:           []string{"gke-terraform-cluster-3-dab431c8-node"},
					Metadata: []*gcp.Metadata{
						{
							Key:   spotinst.String("google-compute-enable-pcid"),
							Value: spotinst.String("true")},
					},
				},
				NetworkInterfaces: []*gcp.NetworkInterface{
					{
						AccessConfigs: []*gcp.AccessConfig{
							{
								Name: spotinst.String("external-nat"),
								Type: spotinst.String("ONE_TO_ONE_NAT"),
							},
						},
						Network: spotinst.String("default"),
					},
				},
				AvailabilityZones: []string{"us-east1-c"},
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
