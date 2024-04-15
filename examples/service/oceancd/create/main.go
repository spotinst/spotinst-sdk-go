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
	out, err := svc.CreateVerificationProvider(ctx, &oceancd.CreateVerificationProviderInput{
		VerificationProvider: &oceancd.VerificationProvider{
			CloudWatch: &oceancd.CloudWatch{
				IAmArn: spotinst.String("Test"),
			},
			ClusterIDs: []string{"Foo", "Foo"},
			DataDog: &oceancd.DataDog{
				Address: spotinst.String("https://api.datadoghq.eu"),
				ApiKey:  spotinst.String("Api-Key-Test"),
				AppKey:  spotinst.String("App-Key-Test"),
			},
			Jenkins: &oceancd.Jenkins{
				ApiToken: spotinst.String("test-api-token"),
				BaseUrl:  spotinst.String("test-baseurl"),
				UserName: spotinst.String("test-username"),
			},
			Name: spotinst.String("test"),
			NewRelic: &oceancd.NewRelic{
				AccountId:        spotinst.String("test-acc-id"),
				BaseUrlNerdGraph: spotinst.String("foo"),
				BaseUrlRest:      spotinst.String("foo"),
				PersonalApiKey:   spotinst.String("test-personal-api-key"),
				Region:           spotinst.String("eastus"),
			},
			Prometheus: &oceancd.Prometheus{
				Address: spotinst.String("foo"),
			},
		},
	})

	if err != nil {
		log.Fatalf("spotinst: failed to create policy: %v", err)
	}

	// Output.
	if out.VerificationProvider != nil {
		log.Printf("VerificationProvider %q: %s",
			spotinst.StringValue(out.VerificationProvider.Name),
			stringutil.Stringify(out.VerificationProvider))
	}
}
