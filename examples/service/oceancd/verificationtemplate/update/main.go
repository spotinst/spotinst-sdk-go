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

	// Update Verification Template configuration.
	out, err := svc.UpdateVerificationTemplate(ctx, &oceancd.UpdateVerificationTemplateInput{
		VerificationTemplate: &oceancd.VerificationTemplate{
			Name: spotinst.String("name"),
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
		},
	})
	if err != nil {
		log.Fatalf("spotinst: failed to update Verification Template: %v", err)
	}

	// Output.
	if out.VerificationTemplate != nil {
		log.Printf("Verification Template %q: %s",
			spotinst.StringValue(out.VerificationTemplate.Name),
			stringutil.Stringify(out.VerificationTemplate))
	}
}
