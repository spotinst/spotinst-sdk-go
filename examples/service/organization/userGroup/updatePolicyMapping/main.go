package main

import (
	"context"
	"log"

	"github.com/spotinst/spotinst-sdk-go/service/organization"
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
	svc := organization.New(sess)

	// Create a new context.
	ctx := context.Background()

	// Create a new group.
	err := svc.UpdatePolicyMappingOfUserGroup(ctx, &organization.UpdatePolicyMappingOfUserGroupInput{
		UserGroupId: spotinst.String("ugr-17ae43d9"),
		Policies: []*organization.UserPolicy{
			&organization.UserPolicy{
				PolicyId: spotinst.String("pol-b236db1f"),
				AccountIds: []string{
					"act-abcd1234",
				},
			},
			&organization.UserPolicy{
				PolicyId: spotinst.String("pol-08715c90"),
				AccountIds: []string{
					"act-abcd1234",
				},
			},
			&organization.UserPolicy{
				PolicyId: spotinst.String("3"),
				AccountIds: []string{
					"act-abcd1234",
				},
			},
			&organization.UserPolicy{
				PolicyId: spotinst.String("pol-c75d8c06"),
				AccountIds: []string{
					"act-abcd1234",
				},
			},
		},
	})

	if err != nil {
		log.Fatalf("spotinst: failed to update user group: %v", err)
	}

}
