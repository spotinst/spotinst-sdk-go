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

	// Update Strategy configuration.
	out, err := svc.UpdateStrategy(ctx, &oceancd.UpdateStrategyInput{
		Strategy: &oceancd.Strategy{
			Name: spotinst.String("name"),
			Canary: &oceancd.Canary{
				BackgroundVerification: &oceancd.BackgroundVerification{
					TemplateNames: []string{"Test", "Test"},
				},
				Steps: []*oceancd.CanarySteps{
					{
						Name: spotinst.String("Test"),
						Pause: &oceancd.Pause{
							Duration: spotinst.String("10s"),
						},
						SetWeight: spotinst.Int(20),
						Verification: &oceancd.Verification{
							TemplateNames: []string{"Test", "Test"},
						},
					},
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("spotinst: failed to update Strategy: %v", err)
	}

	// Output.
	if out.Strategy != nil {
		log.Printf("Strategy %q: %s",
			spotinst.StringValue(out.Strategy.Name),
			stringutil.Stringify(out.Strategy))
	}
}
