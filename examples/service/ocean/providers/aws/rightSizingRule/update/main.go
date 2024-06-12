package main

import (
	"context"
	"log"

	"github.com/spotinst/spotinst-sdk-go/service/ocean"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/providers/aws"
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

	// Update an existing right sizing rule
	out, err := svc.CloudProviderAWS().UpdateRightSizingRule(ctx, &ocean.UpdateRightSizingRuleInput{
		RuleName: spotinst.String("tf-rule"),
		RightSizingRule: &ocean.RightSizingRule{
			Name:        spotinst.String("tf-rule-updated"),
			OceanId:     spotinst.String("o-1234abcd"),
			RestartPods: spotinst.Bool(true),
			RecommendationApplicationIntervals: []*ocean.RecommendationApplicationInterval{
				&ocean.RecommendationApplicationInterval{
					RepetitionBasis: spotinst.String("WEEKLY"),
					WeeklyRepetitionBasis: &ocean.WeeklyRepetitionBasis{
						IntervalDays: []string{"MONDAY"},
						IntervalHours: &ocean.IntervalHours{
							StartTime: spotinst.String("13:00"),
							EndTime:   spotinst.String("15:00"),
						},
					},
				},
			},
			RecommendationApplicationBoundaries: &ocean.RecommendationApplicationBoundaries{
				Cpu: &ocean.Cpu{
					Min: spotinst.Int(9),
					Max: spotinst.Int(99),
				},
				Memory: &ocean.Memory{
					Min: spotinst.Int(9),
					Max: spotinst.Int(99),
				},
			},
			RecommendationApplicationMinThreshold: &ocean.RecommendationApplicationMinThreshold{
				CpuPercentage:    spotinst.Float64(0.75),
				MemoryPercentage: spotinst.Float64(0.75),
			},
		},
	})
	if err != nil {
		log.Fatalf("spotinst: failed to create right sizing rule: %v", err)
	}

	// Output.
	if out.RightSizingRule != nil {
		log.Printf("RightSizing  Rule %q: %s",
			spotinst.StringValue(out.RightSizingRule.Name),
			stringutil.Stringify(out.RightSizingRule))
	}
}
