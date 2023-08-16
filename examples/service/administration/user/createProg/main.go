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
	out, err := svc.CreateProgUser(ctx, &administration.ProgrammaticUser{
		Name:        spotinst.String("test-programmatic-user"),
		Description: spotinst.String("description"),
		/*Accounts: []*administration.Account{
			{
				Id:   spotinst.String("act-7c46c6df"),
				Role: spotinst.String("viewer"),
			},
		},*/ //Accounts and Policies are exclusive
		Policies: []*administration.ProgPolicy{
			{
				PolicyId: spotinst.String("pol-c75d8c06"),
				AccountIds: []string{
					"act-7c46c6df",
				},
			},
		},
	})

	if err != nil {
		log.Fatalf("spotinst: failed to create user: %v", err)
	}

	// Output.
	if out.ProgrammaticUser != nil {
		log.Printf("User %q: %s",
			spotinst.StringValue(out.ProgrammaticUser.ProgUserId),
			stringutil.Stringify(out.ProgrammaticUser))
	}
}
