package main_test

import (
	"fmt"
	"os"

	"github.com/spotinst/spotinst-sdk-go/spotinst/credentials"
	"github.com/spotinst/spotinst-sdk-go/spotinst/featureflag"
)

func Example_static() {
	// Initialize a new static credentials provider.
	provider := credentials.NewStaticCredentials("secret", "acc-12345")

	value, err := provider.Get()
	if err != nil {
		panic(err)
	}

	fmt.Println(value.Token, value.Account)
	// Output: secret acc-12345
}

func Example_env() {
	// Set both token and account.
	//
	// Can be set using an environment variables as well, for example:
	// export SPOTINST_TOKEN=secret
	// export SPOTINST_ACCOUNT=acc-12345
	os.Setenv(credentials.EnvCredentialsVarToken, "secret")
	os.Setenv(credentials.EnvCredentialsVarAccount, "acc-12345")

	// Unset.
	defer func() {
		os.Unsetenv(credentials.EnvCredentialsVarToken)
		os.Unsetenv(credentials.EnvCredentialsVarAccount)
	}()

	// Initialize a new env credentials provider.
	provider := credentials.NewEnvCredentials()

	value, err := provider.Get()
	if err != nil {
		panic(err)
	}

	fmt.Println(value.Token, value.Account)
	// Output: secret acc-12345
}

func Example_chainAllowPartial() {
	// Set both token and account.
	os.Setenv(credentials.EnvCredentialsVarToken, "secret")

	// Unset.
	defer func() {
		os.Unsetenv(credentials.EnvCredentialsVarToken)
	}()

	// Disable the usage of merging credentials in chain provider.
	//
	// Can be set using an environment variable as well, for example:
	// export SPOTINST_FEATURE_FLAGS=MergeCredentialsChain=false
	featureflag.Set("MergeCredentialsChain=false")

	// Initialize a new chain credentials provider.
	provider := credentials.NewChainCredentials(
		&credentials.EnvProvider{},
		&credentials.StaticProvider{
			Value: credentials.Value{
				Account: "acc-12345",
			},
		},
	)

	value, err := provider.Get()
	if err != nil {
		panic(err)
	}

	fmt.Println(value.Token, value.Account)
	// Output: secret
}

func Example_chainAllowMerge() {
	os.Setenv(credentials.EnvCredentialsVarToken, "secret")
	defer func() { os.Unsetenv(credentials.EnvCredentialsVarToken) }()

	// Enable the usage of merging credentials in chain provider.
	//
	// Can be set using an environment variable as well, for example:
	// export SPOTINST_FEATURE_FLAGS=MergeCredentialsChain=true
	featureflag.Set("MergeCredentialsChain=true")

	// Initialize a new chain credentials provider.
	provider := credentials.NewChainCredentials(
		&credentials.EnvProvider{},
		&credentials.StaticProvider{
			Value: credentials.Value{
				Account: "acc-12345",
			},
		},
	)

	value, err := provider.Get()
	if err != nil {
		panic(err)
	}

	fmt.Println(value.Token, value.Account)
	// Output: secret acc-12345
}
