package main

import (
	"context"
	"log"

	"github.com/spotinst/spotinst-sdk-go/service/ocean"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/spark"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/session"
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

	// Detach VNG.
	ctx := context.Background()

	// Delete an existing cluster.
	_, err := svc.Spark().DetachVirtualNodeGroup(ctx, &spark.DetachVngInput{
		ClusterID: spotinst.String("osc-12345"),
		VngID:     spotinst.String("ols-12345"),
	})
	if err != nil {
		log.Fatalf("spotinst: failed to detach VNG: %v", err)
	}
}
