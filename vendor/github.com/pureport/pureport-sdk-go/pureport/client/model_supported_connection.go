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

// SupportedConnection struct for SupportedConnection
type SupportedConnection struct {
	BillingPlans          []BillingPlan `json:"billingPlans,omitempty"`
	BillingProductId      string        `json:"billingProductId"`
	Groups                []Link        `json:"groups,omitempty"`
	HighAvailability      bool          `json:"highAvailability,omitempty"`
	Href                  string        `json:"href,omitempty"`
	Id                    string        `json:"id,omitempty"`
	Location              Link          `json:"location"`
	PeeringType           string        `json:"peeringType"`
	Pending               bool          `json:"pending,omitempty"`
	ReachableCloudRegions []Link        `json:"reachableCloudRegions,omitempty"`
	Speed                 int32         `json:"speed"`
	Type                  string        `json:"type"`
}
