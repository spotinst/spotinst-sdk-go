package main

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/service/oceancd"
	"log"

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
	svc := oceancd.New(sess)

	// Create a new context.
	ctx := context.Background()

	// Update RolloutSpec configuration.
	out, err := svc.UpdateRolloutSpec(ctx, &oceancd.UpdateRolloutSpecInput{
		RolloutSpec: &oceancd.RolloutSpec{
			Name: spotinst.String("name"),
			Strategy: &oceancd.RolloutSpecStrategy{
				Name: spotinst.String("StrategyName"),
			},
			Traffic: &oceancd.Traffic{
				Nginx: &oceancd.Nginx{
					AdditionalIngressAnnotations: &oceancd.AdditionalIngressAnnotations{
						CanaryByHeader: spotinst.String("TestHeader"),
						Key1:           spotinst.String("TestKey"),
					},
					AnnotationPrefix: spotinst.String("Test"),
					StableIngress:    spotinst.String("Hello-Ingress"),
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("spotinst: failed to update RolloutSpec: %v", err)
	}

	// Output.
	if out.RolloutSpec != nil {
		log.Printf("RolloutSpec %q: %s",
			spotinst.StringValue(out.RolloutSpec.Name),
			stringutil.Stringify(out.RolloutSpec))
	}
}
