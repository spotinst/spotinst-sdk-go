package main

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/service/ocean"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/right_sizing"
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

	// Attach Workloads to existing Right Sizing Rule
	_, err := svc.RightSizing().AttachRightSizingRule(ctx, &right_sizing.RightSizingAttachDetachInput{
		RuleName: spotinst.String("test-rule1"),
		OceanId:  spotinst.String("o-123ab54"),
		Namespaces: []*right_sizing.Namespace{
			{
				NamespaceName: spotinst.String("nameSpace"),
				Workloads: []*right_sizing.Workload{
					{
						Name:         spotinst.String("workloadName"),
						WorkloadType: spotinst.String("Deployment"),
						// RegexName: spotinst.String("test*"), Regex and Name are mutually exclusive
					},
				},
			},
			{
				NamespaceName: spotinst.String("nameSpace"),
				Labels: []*right_sizing.Label{
					{
						Key:   spotinst.String("test-key"),
						Value: spotinst.String("test-value"),
					},
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("spotinst: failed to attach workloads to right sizing rule: %v", err)
	}

}
