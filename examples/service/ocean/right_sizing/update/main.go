package main

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/service/ocean"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/right_sizing"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/session"
	"github.com/spotinst/spotinst-sdk-go/spotinst/util/stringutil"
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

	// Update an existing right sizing rule
	out, err := svc.RightSizing().UpdateRightsizingRule(ctx, &right_sizing.UpdateRightsizingRuleInput{
		RuleName: spotinst.String("test-rule1"),
		RightsizingRule: &right_sizing.RightsizingRule{
			OceanId: spotinst.String("o-12ab34"),
			RecommendationApplicationIntervals: []*right_sizing.RecommendationApplicationIntervals{
				{
					RepetitionBasis: spotinst.String("MONTHLY"),
					MonthlyRepetitionBasis: &right_sizing.MonthlyRepetitionBasis{
						IntervalMonths: []int{4, 5, 6},
						WeekOfTheMonth: []string{"THIRD", "SECOND"},
						WeeklyRepetitionBasis: &right_sizing.WeeklyRepetitionBasis{
							IntervalDays: []string{"MONDAY"},
							IntervalHours: &right_sizing.IntervalHours{
								StartTime: spotinst.String("10:00"),
								EndTime:   spotinst.String("18:00"),
							},
						},
					},
				},
			},
			RecommendationApplicationBoundaries: &right_sizing.RecommendationApplicationBoundaries{
				Cpu: &right_sizing.Cpu{
					Min: spotinst.Float64(30),
					Max: spotinst.Float64(80),
				},
				Memory: &right_sizing.Memory{
					Min: spotinst.Int(20),
					Max: spotinst.Int(70),
				},
			},
			RecommendationApplicationMinThreshold: &right_sizing.RecommendationApplicationMinThreshold{
				CpuPercentage:    spotinst.Float64(0.50),
				MemoryPercentage: spotinst.Float64(0.75),
			},
			RecommendationApplicationOverheadValues: &right_sizing.RecommendationApplicationOverheadValues{
				CpuPercentage:    spotinst.Float64(0.75),
				MemoryPercentage: spotinst.Float64(0.50),
			},
		},
	})
	if err != nil {
		log.Fatalf("spotinst: failed to update the right sizing rule: %v", err)
	}

	// Output.
	if out.RightsizingRule != nil {
		log.Printf("RightSizing  Rule %q:",
			stringutil.Stringify(out.RightsizingRule))
	}
}
