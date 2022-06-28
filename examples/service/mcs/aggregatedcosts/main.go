package main

import (
	"context"
	"encoding/json"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/credentials"
	"github.com/spotinst/spotinst-sdk-go/spotinst/session"
	"log"
	"os"
)

func main() {
	// All clients require a Session. The Session provides the client with
	// shared configuration such as account and credentials.
	// A Session should be shared where possible to take advantage of
	// configuration and credential caching. See the session package for
	// more information.
	sess := session.New()

	os.Setenv(credentials.EnvCredentialsVarToken, "secret token")
	os.Setenv(credentials.EnvCredentialsVarAccount, "some account")

	cred := credentials.NewChainCredentials(
		new(credentials.EnvProvider),
	)

	// Create a new instance of the service's client with a Session.
	// Optional spotinst.Config values can also be provided as variadic
	// arguments to the New function. This option allows you to provide
	// service specific configuration.
	svc := aws.New(sess, &spotinst.Config{Credentials: cred})

	// Create a new context.
	ctx := context.Background()

	// get aggregated costs of a new ocean cluster.

	AllMatch := []*aws.AllMatchInner{{
		Type:     spotinst.String("label"),
		Key:      spotinst.String("k8s-app"),
		Operator: spotinst.String("notEquals"),
		Value:    spotinst.String("coredns-autoscaler"),
	}}
	AllMatchArray := []*aws.AllMatch{{AllMatches: AllMatch}}

	out, err := svc.GetClusterAggregatedCosts(ctx, &aws.ClusterAggregatedCostInput{
		OceanId:   spotinst.String("o-56d4124b"),
		StartTime: spotinst.String("1655769600000"),
		EndTime:   spotinst.String("1655856000000"),
		GroupBy:   spotinst.String("resource.label.K8s-App"),
		Filter: &aws.AggregatedFilter{
			Scope:      spotinst.String("resource"),
			Conditions: &aws.Conditions{AnyMatch: AllMatchArray},
		},
	})
	if err != nil {
		log.Fatalf("spotinst: Failed to revieve the aggregated costs: %v", err)
	}

	output, _ := json.Marshal(out.Items)
	// Do something with output.
	if out.Items != nil {
		log.Printf("Aggregated Costs:\n %s",
			output)
	}
}
