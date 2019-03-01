package k8s

import (
	"github.com/linkerd/linkerd2/pkg/k8s"
)

// NewFakeAPI provides a mock Kubernetes API for testing.
func NewFakeAPI(configs ...string) (*API, error) {
	clientSet, spClientSet := k8s.NewFakeClientSets(configs...)

	return NewAPI(
		clientSet,
		spClientSet,
		CM,
		Deploy,
		DS,
		SS,
		Endpoint,
		Pod,
		RC,
		RS,
		Svc,
		SP,
		MWC,
	), nil
}
