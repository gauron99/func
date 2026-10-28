/*
Package deployers provides canonical short names for the function
deployer implementations.
*/
package deployers

import "fmt"

const (
	Knative    = "knative"
	Kubernetes = "raw"
	Keda       = "keda"

	// Default deployer absent any other configuration.
	Default = Knative
)

// Resolve returns the deployer a function deploys with: the one it asks for
// (intent), else the one it was last deployed with (deployed), else Default.
// Every place that picks a deployer for a function uses it, so they all agree
// on a function that names none.
func Resolve(intent, deployed string) string {
	if intent != "" {
		return intent
	}
	if deployed != "" {
		return deployed
	}
	return Default
}

// ValidateSwitch reports an error if an already-deployed function is being
// redeployed with a different deployer. 'from' is the already deployed deployer
// and 'to' is the new deployer "to deploy with"
func ValidateSwitch(from, to string) error {
	if from == "" || to == "" {
		return nil
	}
	if from == to {
		return nil
	}
	return fmt.Errorf("function was deployed with the %q deployer; redeploying with %q is not supported. Run func delete first, then redeploy", from, to)
}
