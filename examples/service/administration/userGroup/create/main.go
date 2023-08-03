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
	out, err := svc.CreateUserGroup(ctx, &administration.UserGroup{
		Description: spotinst.String("TFUserGroup"),
		Name:        spotinst.String("Pritish"),
		Policies: []*administration.UserGroupPolicy{
			{
				AccountIds: spotinst.StringSlice([]string{
					"act-7c46c6df",
				}),
				PolicyId: spotinst.String("pol-c063c6c9"),
			},
		},
	})

	if err != nil {
		log.Fatalf("spotinst: failed to create user group: %v", err)
	}

	// Output.
	if out.UserGroup != nil {
		log.Printf("UserGroup %q: %s",
			spotinst.StringValue(out.UserGroup.UserGroupId),
			stringutil.Stringify(out.UserGroup))
	}
}
