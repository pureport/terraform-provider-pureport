/*
 * Pureport Control Plane
 *
 * Pureport API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

import (
	"time"
)

type PublicPeeringBlock struct {
	CidrBlock string    `json:"cidrBlock"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	Href      string    `json:"href,omitempty"`
	Id        string    `json:"id,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}