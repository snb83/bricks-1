// Copyright © 2019 by PACE Telematics GmbH. All rights reserved.
// Created at 2019/03/11 by Florian Hübsch

package transport

// NewDefaultTransportChain returns a transport chain with retry, jaeger and logging support.
// If not explicitly finalized via `Final` it uses `http.DefaultTransport` as finalizer.
func NewDefaultTransportChain() *RoundTripperChain {
	return Chain(
		&ExternalDependencyRoundTripper{},
		NewDefaultRetryRoundTripper(),
		&JaegerRoundTripper{},
		NewDumpRoundTripperEnv(),
		&LoggingRoundTripper{},
		&LocaleRoundTripper{},
		&RequestIDRoundTripper{},
	)
}

// NewDefaultTransportChain returns a transport chain with retry, jaeger and logging support.
// If not explicitly finalized via `Final` it uses `http.DefaultTransport` as finalizer.
// The passed name is recorded as external dependency
func NewDefaultTransportChainWithExternalName(name string) *RoundTripperChain {
	return Chain(
		&ExternalDependencyRoundTripper{name: name},
		NewDefaultRetryRoundTripper(),
		&JaegerRoundTripper{},
		NewDumpRoundTripperEnv(),
		&LoggingRoundTripper{},
		&LocaleRoundTripper{},
		&RequestIDRoundTripper{},
	)
}
