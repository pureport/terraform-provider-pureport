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

import (
	"time"
)

// AuditEntry struct for AuditEntry
type AuditEntry struct {
	Account       Link               `json:"account,omitempty"`
	Changes       []ChangeObject     `json:"changes,omitempty"`
	CorrelationId string             `json:"correlationId,omitempty"`
	EventType     string             `json:"eventType,omitempty"`
	IpAddress     string             `json:"ipAddress,omitempty"`
	Principal     Link               `json:"principal,omitempty"`
	Request       AuditEntryRequest  `json:"request,omitempty"`
	Response      AuditEntryResponse `json:"response,omitempty"`
	Result        string             `json:"result,omitempty"`
	Source        string             `json:"source,omitempty"`
	Subject       Link               `json:"subject,omitempty"`
	SubjectType   string             `json:"subjectType,omitempty"`
	Timestamp     time.Time          `json:"timestamp,omitempty"`
	UserAgent     string             `json:"userAgent,omitempty"`
}