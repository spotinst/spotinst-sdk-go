# Spotinst SDK Go

The official Spotinst SDK for the Go programming language.

## Table of Contents

- [Installation](#installation)
- [Authentication](#authentication)
- [Complete SDK Example](#complete-sdk-example)
- [Documentation](#documentation)
- [Examples](#examples)
- [Getting Help](#getting-help)
- [Community](#community)
- [Contributing](#contributing)
- [License](#license)

## Installation

The best way to get started working with the SDK is to use go get to add the SDK to your Go application using Go modules.

```
go get -u github.com/spotinst/spotinst-sdk-go
```

Without Go Modules, or in a GOPATH with Go 1.11 or 1.12 use the /... suffix on the go get to retrieve all of the SDK's dependencies.

```
go get -u github.com/spotinst/spotinst-sdk-go/...
```

## Authentication

Set a `ChainProvider` that will search for a provider which returns credentials.

The `ChainProvider` provides a way of chaining multiple providers together
which will pick the first available using priority order of the Providers
in the list. If none of the Providers retrieve valid credentials, `ChainProvider`'s
`Retrieve()` will return the error `ErrNoValidProvidersFoundInChain`. If a Provider
is found which returns valid credentials `ChainProvider` will cache that Provider
for all calls until `Retrieve` is called again.

Example of `ChainProvider` to be used with an `EnvCredentialsProvider` and
`FileCredentialsProvider`. In this example `EnvProvider` will first check if
any credentials are available via the SPOTINST_TOKEN and SPOTINST_ACCOUNT environment variables. If there are
none `ChainProvider` will check the next `Provider` in the list, `FileProvider`
in this case. If `FileCredentialsProvider` does not return any credentials
`ChainProvider` will return the error `ErrNoValidProvidersFoundInChain`.

```go
// Initial credentials loaded from SDK's default credential chain. Such as
// the environment, shared credentials (~/.spotinst/credentials), etc.
sess := session.New()

// Create the chain credentials.
creds := credentials.NewChainCredentials(
    new(credentials.FileProvider),
    new(credentials.EnvProvider),
)

// Create service client value configured for credentials
// from the chain.
svc := elastigroup.New(sess, &spotinst.Config{Credentials: creds})
```

## Complete SDK Example

```go
package main

import (
	"context"
	"log"

	"github.com/spotinst/spotinst-sdk-go/service/elastigroup"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
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
	svc := elastigroup.New(sess)

	// Create a new context.
	ctx := context.Background()

	// Read group configuration.
	out, err := svc.CloudProviderAWS().Read(ctx, &aws.ReadGroupInput{
		GroupID: spotinst.String("sig-12345"),
	})
	if err != nil {
		log.Fatalf("spotinst: failed to read group: %v", err)
	}

	// Output.
	if out.Group != nil {
		log.Printf("Group %q: %s",
			spotinst.StringValue(out.Group.ID),
			stringutil.Stringify(out.Group))
	}
}
```

## Documentation

For a comprehensive documentation, check out the [Spot Documentation](https://docs.spot.io/) website.

## Examples

For a list of examples, check out the [examples](/examples) directory.

## Getting Help

We use GitHub issues for tracking bugs and feature requests. Please use these community resources for getting help:

- Ask a question on [Stack Overflow](https://stackoverflow.com/) and tag it with [spotinst-sdk-go](https://stackoverflow.com/questions/tagged/spotinst-sdk-go/).
- Join our Spotinst community on [Slack](http://slack.spot.io/).
- Open an [issue](https://github.com/spotinst/spotinst-sdk-go/issues/new/choose/).

## Community

- [Slack](http://slack.spot.io/)
- [Twitter](https://twitter.com/spot_hq/)

## Contributing

Please see the [contribution guidelines](.github/CONTRIBUTING.md).

## License

Code is licensed under the [Apache License 2.0](LICENSE). See [NOTICE.md](NOTICE.md) for complete details, including software and third-party licenses and permissions.
