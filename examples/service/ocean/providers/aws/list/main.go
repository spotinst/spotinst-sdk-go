package main

import (
	"context"
	"log"

	"github.com/spotinst/spotinst-sdk-go/service/ocean"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
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

	// List eks clusters.
	eks, err := svc.CloudProviderAWS().ListClusters(ctx, &aws.ListClustersInput{})
	if err != nil {
		log.Fatalf("spotinst: failed to list clusters: %v", err)
	}

	// Output all clusters, if any.
	if len(eks.Clusters) > 0 {
		for _, cluster := range eks.Clusters {
			log.Printf("Cluster %q: %s",
				spotinst.StringValue(cluster.ID),
				stringutil.Stringify(cluster))
		}
	}

	// List ecs clusters.
	ecs, err := svc.CloudProviderAWS().ListECSClusters(ctx, &aws.ListECSClustersInput{})
	if err != nil {
		log.Fatalf("spotinst: failed to list clusters: %v", err)
	}

	// Output all clusters, if any.
	if len(ecs.Clusters) > 0 {
		for _, cluster := range ecs.Clusters {
			log.Printf("Cluster %q: %s",
				spotinst.StringValue(cluster.ID),
				stringutil.Stringify(cluster))
		}
	}

}
