// A dummy main to help with releasing the Spotinst SDK Go module.
package main

import (
	"fmt"
)

// TODO(liran): delete this when we find a better way to generate release notes.
func main() {
	fmt.Println(`
This 'main' exists only to make goreleaser create release notes for the SDK.
See: https://github.com/goreleaser/goreleaser/issues/981`)
}
