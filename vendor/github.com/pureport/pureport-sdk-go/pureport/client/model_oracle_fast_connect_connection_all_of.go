/*
 * Pureport Control Plane
 *
 * Pureport API
 *
 * API version: 1.0.0
 * Contact: support@pureport.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// OracleFastConnectConnectionAllOf struct for OracleFastConnectConnectionAllOf
type OracleFastConnectConnectionAllOf struct {
	CloudRegion Link                  `json:"cloudRegion,omitempty"`
	Peering     *PeeringConfiguration `json:"peering,omitempty"`
	// The primary Oracle Cloud ID (OCID) for the Oracle Fast Connect.
	PrimaryOcid string `json:"primaryOcid,omitempty"`
	// The secondary Oracle Cloud ID (OCID) for the Oracle Fast Connect.
	SecondaryOcid string `json:"secondaryOcid,omitempty"`
}