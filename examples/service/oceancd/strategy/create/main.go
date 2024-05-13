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

	// Create a new group.
	out, err := svc.CreateStrategy(ctx, &oceancd.CreateStrategyInput{
		Strategy: &oceancd.Strategy{
			Name: spotinst.String("TestStartegy"),
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
						SetCanaryScale: &oceancd.SetCanaryScale{
							MatchTrafficWeight: spotinst.Bool(true),
							Replicas:           spotinst.Int(2),
							Weight:             spotinst.Int(30),
						},
						SetHeaderRoute: &oceancd.SetHeaderRoute{
							Match: []*oceancd.Match{
								{
									HeaderName: spotinst.String("Test"),
									HeaderValue: &oceancd.HeaderValue{
										Exact:  spotinst.String("Test"),
										Prefix: spotinst.String("Test"),
										Regex:  spotinst.String("Test"),
									},
								},
							},
							Name: spotinst.String("Test"),
						},
						SetWeight: spotinst.Int(20),
						Verification: &oceancd.Verification{
							TemplateNames: []string{"Test", "Test"},
						},
					},
				},
			},
			Rolling: &oceancd.Rolling{
				Steps: []*oceancd.RollingSteps{
					{
						Name: spotinst.String("Test"),
						Pause: &oceancd.Pause{
							Duration: spotinst.String("10s"),
						},
						Verification: &oceancd.Verification{
							TemplateNames: []string{"Test", "Test"},
						},
					},
				},
			},
		},
	})

	if err != nil {
		log.Fatalf("spotinst: failed to create strategy: %v", err)
	}

	// Output.
	if out.Strategy != nil {
		log.Printf("Strategy %q: %s",
			spotinst.StringValue(out.Strategy.Name),
			stringutil.Stringify(out.Strategy))
	}
}
