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

	// Create a new Verification Template.
	out, err := svc.CreateVerificationTemplate(ctx, &oceancd.CreateVerificationTemplateInput{
		VerificationTemplate: &oceancd.VerificationTemplate{
			Name: spotinst.String("Test"),
			Args: []*oceancd.Args{
				{
					Name:  spotinst.String("test"),
					Value: spotinst.String("test"),
					ValueFrom: &oceancd.ValueFrom{
						SecretKeyRef: &oceancd.SecretKeyRef{
							Key: spotinst.String("TestKey"),
						},
					},
				},
			},
			Metrics: []*oceancd.Metrics{
				{
					Name:                  spotinst.String("TestMetric"),
					DryRun:                spotinst.Bool(true),
					Interval:              spotinst.String("5m"),
					InitialDelay:          spotinst.String("2s"),
					Count:                 spotinst.Int(1),
					FailureCondition:      spotinst.String("test[0]>1"),
					FailureLimit:          spotinst.Int(1),
					SuccessCondition:      spotinst.String("test[1]>1"),
					ConsecutiveErrorLimit: spotinst.Int(4),
					Provider: &oceancd.Provider{
						CloudWatch: &oceancd.CloudWatchProvider{
							Duration: spotinst.String("1s"),
							MetricDataQueries: []*oceancd.MetricDataQueries{
								{
									ID: spotinst.String("TestID"),
									MetricStat: &oceancd.MetricStat{
										Metric: &oceancd.Metric{
											MetricName: spotinst.String("TestName"),
											Namespace:  spotinst.String("Test"),
											Dimensions: []*oceancd.Dimensions{
												{
													Name:  spotinst.String("Test"),
													Value: spotinst.String("Test"),
												},
											},
										},
										Period: spotinst.Int(30),
										Stat:   spotinst.String("Test"),
										Unit:   spotinst.String("none"),
									},
									Expression: spotinst.String("Test"),
									Label:      spotinst.String("TestLabel"),
									ReturnData: spotinst.Bool(false),
									Period:     spotinst.Int(30),
								},
							},
						},
						Datadog: &oceancd.DataDogProvider{
							Duration: spotinst.String("3s"),
							Query:    spotinst.String("avg.cpu.utilization"),
						},
						Prometheus: &oceancd.PrometheusProvider{
							Query: spotinst.String("Total_CPU_Test"),
						},
						Job: &oceancd.Job{
							Spec: &oceancd.Spec{
								BackoffLimit: spotinst.Int(1),
								Template: &oceancd.Template{
									Spec: &oceancd.TemplateSpec{
										RestartPolicy: spotinst.String("never"),
										Containers: []*oceancd.Containers{
											{
												Name:    spotinst.String("Test"),
												Command: []string{"Test"},
												Image:   spotinst.String("Test"),
											},
										},
									},
								},
							},
						},
						NewRelic: &oceancd.NewRelicProvider{
							Profile: spotinst.String("Test"),
							Query:   spotinst.String("Test"),
						},
						Jenkins: &oceancd.JenkinsProvider{
							PipelineName:    spotinst.String("Test"),
							TLSVerification: spotinst.Bool(true),
							Timeout:         spotinst.String("test"),
							Interval:        spotinst.String("2s"),
							Parameters: []*oceancd.Parameters{
								{
									Key:   spotinst.String("test"),
									Value: spotinst.String("Test"),
								},
							},
						},
						Web: &oceancd.Web{
							Method: spotinst.String("Test"),
							Url:    spotinst.String("TestUrl"),
							Headers: []*oceancd.Headers{
								{
									Key:   spotinst.String("Test"),
									Value: spotinst.String("Test"),
								},
							},
							Body:           spotinst.String("Test"),
							TimeoutSeconds: spotinst.Int(30),
							JsonPath:       spotinst.String("Test"),
							Insecure:       spotinst.Bool(true),
						},
					},
					Baseline: &oceancd.Baseline{
						Threshold: spotinst.String(">"),
						MaxRange:  spotinst.Int(30),
						MinRange:  spotinst.Int(40),
						Provider: &oceancd.Provider{
							Datadog: &oceancd.DataDogProvider{
								Duration: spotinst.String("3s"),
								Query:    spotinst.String("avg.cpu.utilization"),
							},
							Prometheus: &oceancd.PrometheusProvider{
								Query: spotinst.String("Total_CPU_Test"),
							},
							NewRelic: &oceancd.NewRelicProvider{
								Profile: spotinst.String("Test"),
								Query:   spotinst.String("Test"),
							},
						},
					},
				},
			},
		},
	})

	if err != nil {
		log.Fatalf("spotinst: failed to create verification provider: %v", err)
	}

	// Output.
	if out.VerificationTemplate != nil {
		log.Printf("VerificationTemplate %q: %s",
			spotinst.StringValue(out.VerificationTemplate.Name),
			stringutil.Stringify(out.VerificationTemplate))
	}
}
