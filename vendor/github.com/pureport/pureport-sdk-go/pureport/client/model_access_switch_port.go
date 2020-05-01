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

// AccessSwitchPort struct for AccessSwitchPort
type AccessSwitchPort struct {
	AccessSwitch          Link    `json:"accessSwitch"`
	ConnectorType         string  `json:"connectorType"`
	Href                  string  `json:"href,omitempty"`
	Id                    string  `json:"id"`
	MediaType             string  `json:"mediaType"`
	Name                  string  `json:"name"`
	PatchPanelId          string  `json:"patchPanelId"`
	PatchPanelPortNumbers []int32 `json:"patchPanelPortNumbers"`
	Speed                 int32   `json:"speed"`
}