// Package constant holds HTTP header names, service identifiers, and other
// string literals shared across every service.
package constant

const (
	HEADER_USER_ID   = "X-User-ID"
	HEADER_CLIENT_ID = "X-Client-ID"
	HEADER_MODULE    = "X-Module"

	// HEADER_SERVICE_TOKEN carries the shared secret that authenticates one SERVICE to another on
	// the /internal endpoints. The public gateway must never forward it from a client.
	HEADER_SERVICE_TOKEN = "X-Service-Token"
)
