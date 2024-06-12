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

	// Create a new right sizing rule.
	out, err := svc.CloudProviderAWS().CreateRightSizingRule(ctx, &aws.CreateRightSizingRuleInput{
		RightSizingRule: &aws.RightSizingRule{
			Name:        spotinst.String("tf-rule"),
			OceanId:     spotinst.String("o-1234abcd"),
			RestartPods: spotinst.Bool(true),
			RecommendationApplicationIntervals: []*aws.RecommendationApplicationInterval{
				&aws.RecommendationApplicationInterval{
					RepetitionBasis: spotinst.String("MONTHLY"),
					MonthlyRepetitionBasis: &aws.MonthlyRepetitionBasis{
						IntervalMonths: []int{1, 2, 3},
						WeekOfTheMonth: []string{"FIRST", "SECOND"},
						WeeklyRepetitionBasis: &aws.WeeklyRepetitionBasis{
							IntervalDays: []string{"TUESDAY"},
							IntervalHours: &aws.IntervalHours{
								StartTime: spotinst.String("12:00"),
								EndTime:   spotinst.String("16:00"),
							},
						},
					},
				},
			},
			RecommendationApplicationBoundaries: &aws.RecommendationApplicationBoundaries{
				Cpu: &aws.Cpu{
					Min: spotinst.Int(10),
					Max: spotinst.Int(100),
				},
				Memory: &aws.Memory{
					Min: spotinst.Int(10),
					Max: spotinst.Int(100),
				},
			},
			RecommendationApplicationMinThreshold: &aws.RecommendationApplicationMinThreshold{
				CpuPercentage:    spotinst.Float64(0.5),
				MemoryPercentage: spotinst.Float64(0.5),
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
