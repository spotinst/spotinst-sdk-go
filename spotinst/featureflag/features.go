package featureflag

// Default features.
var (
	// Toggle the usage of merging credentials in chain provider.
	//
	// This feature allows users to configure their credentials using multiple
	// providers. For example, a token can be statically configured using a file,
	// while the account can be dynamically configured via environment variables.
	MergeCredentialsChain = New("MergeCredentialsChain", false)
)

func init() {
	SetFromEnv()
}
