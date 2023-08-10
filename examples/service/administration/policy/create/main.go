package main

import (
	"context"
	"log"

	"github.com/spotinst/spotinst-sdk-go/service/administration"
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
	svc := administration.New(sess)

	// Create a new context.
	ctx := context.Background()

	// Create a new group.
	out, err := svc.CreatePolicy(ctx, &administration.CreatePolicyInput{
		Policy: &administration.Policy{
			Description: spotinst.String("TFPolicy"),
			Name:        spotinst.String("Pritish"),
			PolicyContent: &administration.PolicyContent{
				Statements: []*administration.Statement{
					{
						Actions: []string{
							"ocean:deleteCluster",
							"ocean:updateCluster",
						},
						Effect: spotinst.String("ALLOW"),
						Resources: []string{
							"*",
						},
					},
				},
			},
		},
	})

	if err != nil {
		log.Fatalf("spotinst: failed to create policy: %v", err)
	}

	// Output.
	if out.Policy != nil {
		log.Printf("Policy %q: %s",
			spotinst.StringValue(out.Policy.PolicyID),
			stringutil.Stringify(out.Policy))
	}
}
