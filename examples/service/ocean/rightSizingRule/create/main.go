package main

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/stringutil"
	"log"

	"github.com/spotinst/spotinst-sdk-go/service/ocean"
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

	// Create a new right sizing rule.
	out, err := svc.CreateRightSizingRule(ctx, &ocean.CreateRightSizingRuleInput{
		RightSizingRule: &ocean.RightSizingRule{
			Name:        spotinst.String("tf-rule-3"),
			OceanId:     spotinst.String("o-9a8a856c"),
			RestartPods: spotinst.Bool(true),
			RecommendationApplicationIntervals: []*ocean.RecommendationApplicationInterval{
				&ocean.RecommendationApplicationInterval{
					RepetitionBasis: spotinst.String("MONTHLY"),
					MonthlyRepetitionBasis: &ocean.MonthlyRepetitionBasis{
						IntervalMonths: []int{1, 2, 3},
						WeekOfTheMonth: []string{"FIRST", "SECOND"},
						WeeklyRepetitionBasis: &ocean.WeeklyRepetitionBasis{
							IntervalDays: []string{"TUESDAY"},
							IntervalHours: &ocean.IntervalHours{
								StartTime: spotinst.String("12:00"),
								EndTime:   spotinst.String("16:00"),
							},
						},
					},
				},
			},
			RecommendationApplicationBoundaries: &ocean.RecommendationApplicationBoundaries{
				Cpu: &ocean.Cpu{
					Min: spotinst.Int(10),
					Max: spotinst.Int(100),
				},
				Memory: &ocean.Memory{
					Min: spotinst.Int(10),
					Max: spotinst.Int(100),
				},
			},
			RecommendationApplicationMinThreshold: &ocean.RecommendationApplicationMinThreshold{
				CpuPercentage:    spotinst.Float64(0.5),
				MemoryPercentage: spotinst.Float64(0.5),
			},
			RecommendationApplicationOverheadValues: &ocean.RecommendationApplicationOverheadValues{
				CpuPercentage:    spotinst.Float64(0.25),
				MemoryPercentage: spotinst.Float64(0.25),
			},
		},
	})
	if err != nil {
		log.Fatalf("spotinst: failed to create right sizing rule: %v", err)
	}

	// Output.
	/*	if out.RightSizingRule != nil {
		log.Printf("RightSizing  Rule %q: %s",
			spotinst.StringValue(out.RightSizingRule.Name),
			stringutil.Stringify(out.RightSizingRule))
	}*/
	if out != nil {
		log.Printf(stringutil.Stringify(out))
	}
}
