package main

import (
	"context"
	"log"

	"github.com/spotinst/spotinst-sdk-go/service/ocean"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/spark"
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
	out, err := svc.Spark().CreateCluster(ctx, &spark.CreateClusterInput{
		Cluster: &spark.CreateClusterRequest{
			OceanClusterID: spotinst.String("ofas-cluster-01"),
			Config: &spark.Config{
				LogCollection: &spark.LogCollectionConfig{
					CollectDriverLogs: spotinst.Bool(true),
				},
				Compute: &spark.ComputeConfig{
					UseTaints:  spotinst.Bool(true),
					CreateVngs: spotinst.Bool(true),
				},
				Ingress: &spark.IngressConfig{
					ServiceAnnotations: map[string]string{
						"my-custom-annotation": "custom_value",
					},
				},
				Webhook: &spark.WebhookConfig{
					UseHostNetwork: spotinst.Bool(true),
					HostNetworkPorts: spotinst.IntSlice([]int{
						2000,
					}),
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
