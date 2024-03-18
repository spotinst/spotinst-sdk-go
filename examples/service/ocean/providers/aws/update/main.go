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

	// Update cluster configuration.
	out, err := svc.CloudProviderAWS().UpdateCluster(ctx, &aws.UpdateClusterInput{
		Cluster: &aws.Cluster{
			ID: spotinst.String("o-12345"),
			Compute: &aws.Compute{
				InstanceTypes: &aws.InstanceTypes{
					Filters: &aws.Filters{
						IncludeFamilies: []string{
							"m*", "t*",
						},
						Categories: []string{"General_purpose", "Compute_optimized"},
					},
				},
				LaunchSpecification: &aws.LaunchSpecification{
					HealthCheckUnhealthyDurationBeforeReplacement: spotinst.Int(60),
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("spotinst: failed to update cluster: %v", err)
	}

	// Output.
	if out.Cluster != nil {
		log.Printf("Cluster %q: %s",
			spotinst.StringValue(out.Cluster.ID),
			stringutil.Stringify(out.Cluster))
	}
}
