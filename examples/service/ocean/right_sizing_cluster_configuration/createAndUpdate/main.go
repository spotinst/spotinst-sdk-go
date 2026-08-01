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

	// Create a new cluster.

	out, err := svc.ClusterRightSizing().PostClusterConfiguration(ctx, &cluster_right_sizing.PostClusterConfigurationInput{
		OceanId:           spotinst.String("o-123456"),
		ClusterIdentifier: spotinst.String("my-cluster-identifier"),
		Config: &cluster_right_sizing.ClusterConfiguration{
			AdjustLimitOnDownsize:           spotinst.Bool(true),
			DownsideOnly:                    spotinst.Bool(true),
			RecommendationsCpuPercentile:    spotinst.Int(99),
			RecommendationsMemoryPercentile: spotinst.Int(100),
		},
	})
	if err != nil {
		log.Fatalf("spotinst: failed to post cluster right-sizing configuration: %v", err)
	}

	if out.ClusterConfiguration != nil {
		log.Printf("Posted cluster right-sizing config: %s",
			stringutil.Stringify(out.ClusterConfiguration))
	}
}
