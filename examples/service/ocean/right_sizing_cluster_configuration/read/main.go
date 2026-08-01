package main

import (
	"context"
	"log"

	"github.com/spotinst/spotinst-sdk-go/service/ocean"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/cluster_right_sizing"
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

	// Read cluster configuration.
	out, err := svc.ClusterRightSizing().ReadClusterConfiguration(ctx, &cluster_right_sizing.ReadClusterConfigurationInput{
		OceanId:           spotinst.String("o-123456"),
		ClusterIdentifier: spotinst.String("my-cluster-identifier"),
	})
	if err != nil {
		log.Fatalf("spotinst: failed to read cluster right-sizing configuration: %v", err)
	}

	if out.ClusterConfiguration != nil {
		log.Printf("Current cluster right-sizing config: %s",
			stringutil.Stringify(out.ClusterConfiguration))
	}
}
