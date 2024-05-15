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
	out, err := svc.CreateRolloutSpec(ctx, &oceancd.CreateRolloutSpecInput{
		RolloutSpec: &oceancd.RolloutSpec{
			Name: spotinst.String("TestRolloutSpec"),
			Strategy: &oceancd.RolloutSpecStrategy{
				Args: []*oceancd.RolloutSpecArgs{
					{
						Name:  spotinst.String("TestArg"),
						Value: spotinst.String("TestArgValue"),
						ValueFrom: &oceancd.RolloutSpecValueFrom{
							FieldRef: &oceancd.FieldRef{
								FieldPath: spotinst.String("metadata.labels[test]"),
							},
						},
					},
				},
				Name: spotinst.String("TestStrategy"),
			},
			FailurePolicy: &oceancd.FailurePolicy{
				Action: spotinst.String("Promote"),
			},
			SpotDeployment: &oceancd.SpotDeployment{
				ClusterId: spotinst.String("TestClusterID"),
				Name:      spotinst.String("SpotDeploymentName"),
				Namespace: spotinst.String("default"),
			},
			SpotDeployments: []*oceancd.SpotDeployment{
				{
					ClusterId: spotinst.String("TestClusterID"),
					Name:      spotinst.String("TestDeployment"),
					Namespace: spotinst.String("default"),
				},
			},
			Traffic: &oceancd.Traffic{
				Alb: &oceancd.Alb{
					AnnotationPrefix: spotinst.String("Test"),
					Ingress:          spotinst.String("Stable Ingress"),
					RootService:      spotinst.String("Root Service Test"),
					ServicePort:      spotinst.Int(8080),
					StickinessConfig: &oceancd.StickinessConfig{
						DurationSeconds: spotinst.Int(30),
						Enabled:         spotinst.Bool(true),
					},
				},
				Ambassador: &oceancd.Ambassador{
					Mappings: []string{"Test1", "Test2"},
				},
				StableService: spotinst.String("stable-service-test"),
				CanaryService: spotinst.String("canary-service-test"),
				Istio: &oceancd.Istio{
					DestinationRule: &oceancd.DestinationRule{
						CanarySubsetName: spotinst.String("testname"),
						Name:             spotinst.String("TestDesitinationRule"),
						StableSubsetName: spotinst.String("testname"),
					},
					VirtualServices: []*oceancd.VirtualServices{
						{
							Name:   spotinst.String("TestVSName"),
							Routes: []string{"Route1", "Route2"},
							TlsRoutes: []*oceancd.TlsRoutes{
								{
									Port:     spotinst.Int(80),
									SniHosts: []string{"test-host"},
								},
							},
						},
					},
				},
				Nginx: &oceancd.Nginx{
					AdditionalIngressAnnotations: &oceancd.AdditionalIngressAnnotations{
						CanaryByHeader: spotinst.String("TestHeader"),
						Key1:           spotinst.String("TestKey"),
					},
					AnnotationPrefix: spotinst.String("Test"),
					StableIngress:    spotinst.String("Hello-Ingress"),
				},
				PingPong: &oceancd.PingPong{
					PingService: spotinst.String("Test-Stable-Service"),
					PongService: spotinst.String("Test-Canary-Service"),
				},
				Smi: &oceancd.Smi{
					RootService:      spotinst.String("Test-Stable-Service"),
					TrafficSplitName: spotinst.String("Test-Name"),
				},
			},
		},
	})

	if err != nil {
		log.Fatalf("spotinst: failed to create rolloutSpec: %v", err)
	}

	// Output.
	if out.RolloutSpec != nil {
		log.Printf("RolloutSpec %q: %s",
			spotinst.StringValue(out.RolloutSpec.Name),
			stringutil.Stringify(out.RolloutSpec))
	}
}
