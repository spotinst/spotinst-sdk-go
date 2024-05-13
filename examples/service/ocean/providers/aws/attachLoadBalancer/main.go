package main

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/service/ocean"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/session"
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

	// Attach LoadBalancer configuration.
	_, err := svc.CloudProviderAWS().AttachLoadBalancer(ctx, &aws.AttachLoadbalancerInput{
		LoadBalancers: []*aws.LoadBalancers{
			&aws.LoadBalancers{
				Name: spotinst.String("Test-LoadBalancer-Attach"),
				Arn:  spotinst.String("arn:aws:"),
				Type: spotinst.String("TARGET_GROUP"),
			},
		},
		ID: spotinst.String("o-1234567"),
	})
	if err != nil {
		log.Fatalf("spotinst: failed to attach Load Balancer to Ocean Cluster: %v", err)
	}
}
