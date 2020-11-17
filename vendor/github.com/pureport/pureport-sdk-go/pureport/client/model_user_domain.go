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

// UserDomain struct for UserDomain
type UserDomain struct {
	EmailDomains []string `json:"emailDomains"`
	Href         string   `json:"href,omitempty"`
	Id           string   `json:"id,omitempty"`
	Provider     string   `json:"provider"`
}