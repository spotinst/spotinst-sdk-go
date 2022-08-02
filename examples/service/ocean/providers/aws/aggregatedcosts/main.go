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
	os.Setenv(credentials.EnvCredentialsVarAccount, "account")

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

	// Get aggregated costs of an ocean cluster.
	AllMatch := []*aws.AllMatchInner{{
		Type:     spotinst.String("label"),
		Key:      spotinst.String("k8s-app"),
		Operator: spotinst.String("notEquals"),
		Value:    spotinst.String("dns-controller"),
	}}
	AllMatchArray := []*aws.AllMatch{{AllMatches: AllMatch}}

	out, err := svc.GetClusterAggregatedCosts(ctx, &aws.ClusterAggregatedCostInput{
		OceanId:   spotinst.String("o-12345"),
		StartTime: spotinst.String("1655812800000"),
		EndTime:   spotinst.String("1655985600000"),
		GroupBy:   spotinst.String("namespace"),
		Filter: &aws.AggregatedFilter{
			Scope:      spotinst.String("resource"),
			Conditions: &aws.Conditions{AnyMatch: AllMatchArray},
		},
	})
	if err != nil {
		log.Fatalf("spotinst: Failed to revieve the aggregated costs: %v", err)
	}

	output, errJson := json.Marshal(out.AggregatedClusterCosts)
	if errJson != nil {
		log.Fatalf("spotinst: Failed to marshal output into Json: %v", err)
	}
	// Do something with output.
	if out.AggregatedClusterCosts != nil {
		log.Printf("Aggregated Costs:\n %s",
			output)
	}

}
