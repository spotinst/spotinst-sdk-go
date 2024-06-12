package main

import (
	"context"
	"log"

	"github.com/spotinst/spotinst-sdk-go/service/ocean"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
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

	// Create a new context.
	ctx := context.Background()

	// Detach Workloads from existing Right Sizing Rule
	_, err := svc.CloudProviderAWS().DetachWorkloadsFromRule(ctx, &ocean.RightSizingAttachDetachInput{
		RuleName: spotinst.String("tf-rule"),
		OceanId:  spotinst.String("o-1234abcd"),
		Namespaces: []*ocean.Namespace{
			&ocean.Namespace{
				NamespaceName: spotinst.String("test-namespace"),
				Workloads: []*ocean.Workload{
					&ocean.Workload{
						Name:         spotinst.String("testdeploy"),
						WorkloadType: spotinst.String("Deployment"),
						// RegexName: spotinst.String("test*"), Regex and Name are mutually exclusive
					},
				},
			},
			&ocean.Namespace{
				NamespaceName: spotinst.String("test-namespace"),
				Labels: []*ocean.Label{
					&ocean.Label{
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
