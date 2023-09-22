package main

import (
	"context"
	"log"

	"github.com/spotinst/spotinst-sdk-go/service/organization"
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
	svc := organization.New(sess)

	// Create a new context.
	ctx := context.Background()

	// Create a new group.
	out, err := svc.CreateUser(ctx, &organization.User{
		Email:     spotinst.String("testautomation@netapp.com"),
		FirstName: spotinst.String("test"),
		LastName:  spotinst.String("user"),
		Password:  spotinst.String("testUser@123"),
		Role:      spotinst.String("viewer"),
	}, spotinst.Bool(true))

	if err != nil {
		log.Fatalf("spotinst: failed to create user: %v", err)
	}

	// Output.
	if out.User != nil {
		log.Printf("User %q: %s",
			spotinst.StringValue(out.User.UserID),
			stringutil.Stringify(out.User))
	}
}
