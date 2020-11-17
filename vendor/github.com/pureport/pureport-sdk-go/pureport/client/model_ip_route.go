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

// IpRoute struct for IpRoute
type IpRoute struct {
	DirectlyConnected bool   `json:"directlyConnected,omitempty"`
	Distance          int64  `json:"distance,omitempty"`
	Fib               bool   `json:"fib,omitempty"`
	InterfaceName     string `json:"interfaceName,omitempty"`
	Metric            int64  `json:"metric,omitempty"`
	NextHop           string `json:"nextHop,omitempty"`
	NextHopConnection Link   `json:"nextHopConnection,omitempty"`
	NextHopGateway    Link   `json:"nextHopGateway,omitempty"`
	Prefix            string `json:"prefix,omitempty"`
	Protocol          string `json:"protocol,omitempty"`
	Selected          bool   `json:"selected,omitempty"`
	Uptime            string `json:"uptime,omitempty"`
}
