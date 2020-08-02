// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// NDESConnector Entity which represents an OnPrem Ndes connector.
type NDESConnector struct {
	// Entity is the base model of NDESConnector
	Entity
	// LastConnectionDateTime Last connection time for the Ndes Connector
	LastConnectionDateTime *time.Time `json:"lastConnectionDateTime,omitempty"`
	// State Ndes Connector Status
	State *NDESConnectorState `json:"state,omitempty"`
	// DisplayName The friendly name of the Ndes Connector.
	DisplayName *string `json:"displayName,omitempty"`
}