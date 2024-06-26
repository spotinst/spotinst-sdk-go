package main

import (
	"context"
	"github.com/spotinst/spotinst-sdk-go/service/ocean"
	"github.com/spotinst/spotinst-sdk-go/service/ocean/right_sizing"
	"log"

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
	out, err := svc.RightSizing().CreateRightsizingRule(ctx, &right_sizing.CreateRightsizingRuleInput{
		RightsizingRule: &right_sizing.RightsizingRule{
			RuleName:    spotinst.String("test-rule1"),
			OceanId:     spotinst.String("o-12ab34"),
			RestartPods: spotinst.Bool(true),
			RecommendationApplicationIntervals: []*right_sizing.RecommendationApplicationIntervals{
				{
					RepetitionBasis: spotinst.String("MONTHLY"),
					MonthlyRepetitionBasis: &right_sizing.MonthlyRepetitionBasis{
						IntervalMonths: []int{1, 2, 3},
						WeekOfTheMonth: []string{"FIRST", "SECOND"},
						WeeklyRepetitionBasis: &right_sizing.WeeklyRepetitionBasis{
							IntervalDays: []string{"TUESDAY"},
							IntervalHours: &right_sizing.IntervalHours{
								StartTime: spotinst.String("12:00"),
								EndTime:   spotinst.String("16:00"),
							},
						},
					},
				},
			},
			RecommendationApplicationBoundaries: &right_sizing.RecommendationApplicationBoundaries{
				Cpu: &right_sizing.Cpu{
					Min: spotinst.Float64(10),
					Max: spotinst.Float64(100),
				},
				Memory: &right_sizing.Memory{
					Min: spotinst.Int(10),
					Max: spotinst.Int(100),
				},
			},
			RecommendationApplicationMinThreshold: &right_sizing.RecommendationApplicationMinThreshold{
				CpuPercentage:    spotinst.Float64(0.5),
				MemoryPercentage: spotinst.Float64(0.5),
			},
			RecommendationApplicationOverheadValues: &right_sizing.RecommendationApplicationOverheadValues{
				CpuPercentage:    spotinst.Float64(0.25),
				MemoryPercentage: spotinst.Float64(0.25),
			},
		},
	})
	if err != nil {
		log.Fatalf("spotinst: failed to create right sizing rule: %v", err)
	}

	// Output.
	if out.RightsizingRule != nil {
		log.Printf("RightSizing  Rule name is %s:",
			spotinst.StringValue(out.RightsizingRule.RuleName))
	}
}
