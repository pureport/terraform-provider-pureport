package databricks

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"github.com/satori/go.uuid"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/databricks/mgmt/2018-04-01/databricks"

// CustomParameterType enumerates the values for custom parameter type.
type CustomParameterType string

const (
	// Bool ...
	Bool CustomParameterType = "Bool"
	// Object ...
	Object CustomParameterType = "Object"
	// String ...
	String CustomParameterType = "String"
)

// PossibleCustomParameterTypeValues returns an array of possible values for the CustomParameterType const type.
func PossibleCustomParameterTypeValues() []CustomParameterType {
	return []CustomParameterType{Bool, Object, String}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Accepted ...
	Accepted ProvisioningState = "Accepted"
	// Canceled ...
	Canceled ProvisioningState = "Canceled"
	// Created ...
	Created ProvisioningState = "Created"
	// Creating ...
	Creating ProvisioningState = "Creating"
	// Deleted ...
	Deleted ProvisioningState = "Deleted"
	// Deleting ...
	Deleting ProvisioningState = "Deleting"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Ready ...
	Ready ProvisioningState = "Ready"
	// Running ...
	Running ProvisioningState = "Running"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
	// Updating ...
	Updating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Accepted, Canceled, Created, Creating, Deleted, Deleting, Failed, Ready, Running, Succeeded, Updating}
}

// ErrorDetail ...
type ErrorDetail struct {
	// Code - The error's code.
	Code *string `json:"code,omitempty"`
	// Message - A human readable error message.
	Message *string `json:"message,omitempty"`
	// Target - Indicates which property in the request is responsible for the error.
	Target *string `json:"target,omitempty"`
}

// ErrorInfo ...
type ErrorInfo struct {
	// Code - A machine readable error code.
	Code *string `json:"code,omitempty"`
	// Message - A human readable error message.
	Message *string `json:"message,omitempty"`
	// Details - error details.
	Details *[]ErrorDetail `json:"details,omitempty"`
	// Innererror - Inner error details if they exist.
	Innererror *string `json:"innererror,omitempty"`
}

// ErrorResponse contains details when the response code indicates an error.
type ErrorResponse struct {
	// Error - The error details.
	Error *ErrorInfo `json:"error,omitempty"`
}

// Operation REST API operation
type Operation struct {
	// Name - Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
	// Display - The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay the object that represents the operation.
type OperationDisplay struct {
	// Provider - Service provider: Microsoft.ResourceProvider
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed.
	Resource *string `json:"resource,omitempty"`
	// Operation - Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`
}

// OperationListResult result of the request to list Resource Provider operations. It contains a list of
// operations and a URL link to get the next set of results.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - List of Resource Provider operations supported by the Resource Provider resource provider.
	Value *[]Operation `json:"value,omitempty"`
	// NextLink - URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`
}

// OperationListResultIterator provides access to a complete listing of Operation values.
type OperationListResultIterator struct {
	i    int
	page OperationListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OperationListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *OperationListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OperationListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OperationListResultIterator) Response() OperationListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OperationListResultIterator) Value() Operation {
	if !iter.page.NotDone() {
		return Operation{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the OperationListResultIterator type.
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return OperationListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (olr OperationListResult) IsEmpty() bool {
	return olr.Value == nil || len(*olr.Value) == 0
}

// operationListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (olr OperationListResult) operationListResultPreparer(ctx context.Context) (*http.Request, error) {
	if olr.NextLink == nil || len(to.String(olr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(olr.NextLink)))
}

// OperationListResultPage contains a page of Operation values.
type OperationListResultPage struct {
	fn  func(context.Context, OperationListResult) (OperationListResult, error)
	olr OperationListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OperationListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.olr)
	if err != nil {
		return err
	}
	page.olr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *OperationListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OperationListResultPage) NotDone() bool {
	return !page.olr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OperationListResultPage) Response() OperationListResult {
	return page.olr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OperationListResultPage) Values() []Operation {
	if page.olr.IsEmpty() {
		return nil
	}
	return *page.olr.Value
}

// Creates a new instance of the OperationListResultPage type.
func NewOperationListResultPage(getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return OperationListResultPage{fn: getNextPage}
}

// Resource the core properties of ARM resources
type Resource struct {
	// ID - READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
}

// Sku SKU for the resource.
type Sku struct {
	// Name - The SKU name.
	Name *string `json:"name,omitempty"`
	// Tier - The SKU tier.
	Tier *string `json:"tier,omitempty"`
}

// TrackedResource the resource model definition for a ARM tracked top level resource
type TrackedResource struct {
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	// ID - READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for TrackedResource.
func (tr TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if tr.Tags != nil {
		objectMap["tags"] = tr.Tags
	}
	if tr.Location != nil {
		objectMap["location"] = tr.Location
	}
	return json.Marshal(objectMap)
}

// Workspace information about workspace.
type Workspace struct {
	autorest.Response `json:"-"`
	// WorkspaceProperties - The workspace properties.
	*WorkspaceProperties `json:"properties,omitempty"`
	// Sku - The SKU of the resource.
	Sku *Sku `json:"sku,omitempty"`
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
	// Location - The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	// ID - READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for Workspace.
func (w Workspace) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if w.WorkspaceProperties != nil {
		objectMap["properties"] = w.WorkspaceProperties
	}
	if w.Sku != nil {
		objectMap["sku"] = w.Sku
	}
	if w.Tags != nil {
		objectMap["tags"] = w.Tags
	}
	if w.Location != nil {
		objectMap["location"] = w.Location
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Workspace struct.
func (w *Workspace) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var workspaceProperties WorkspaceProperties
				err = json.Unmarshal(*v, &workspaceProperties)
				if err != nil {
					return err
				}
				w.WorkspaceProperties = &workspaceProperties
			}
		case "sku":
			if v != nil {
				var sku Sku
				err = json.Unmarshal(*v, &sku)
				if err != nil {
					return err
				}
				w.Sku = &sku
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				w.Tags = tags
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				w.Location = &location
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				w.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				w.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				w.Type = &typeVar
			}
		}
	}

	return nil
}

// WorkspaceCustomBooleanParameter the value which should be used for this field.
type WorkspaceCustomBooleanParameter struct {
	// Type - The type of variable that this is. Possible values include: 'Bool', 'Object', 'String'
	Type CustomParameterType `json:"type,omitempty"`
	// Value - The value which should be used for this field.
	Value *bool `json:"value,omitempty"`
}

// WorkspaceCustomObjectParameter the value which should be used for this field.
type WorkspaceCustomObjectParameter struct {
	// Type - The type of variable that this is. Possible values include: 'Bool', 'Object', 'String'
	Type CustomParameterType `json:"type,omitempty"`
	// Value - The value which should be used for this field.
	Value interface{} `json:"value,omitempty"`
}

// WorkspaceCustomParameters custom Parameters used for Cluster Creation.
type WorkspaceCustomParameters struct {
	// AmlWorkspaceID - The Workspace ID of an Azure Machine Learning Workspace
	AmlWorkspaceID *WorkspaceCustomStringParameter `json:"amlWorkspaceId,omitempty"`
	// CustomVirtualNetworkID - The ID of a Virtual Network where this Databricks Cluster should be created
	CustomVirtualNetworkID *WorkspaceCustomStringParameter `json:"customVirtualNetworkId,omitempty"`
	// CustomPublicSubnetName - The name of a Public Subnet within the Virtual Network
	CustomPublicSubnetName *WorkspaceCustomStringParameter `json:"customPublicSubnetName,omitempty"`
	// CustomPrivateSubnetName - The name of the Private Subnet within the Virtual Network
	CustomPrivateSubnetName *WorkspaceCustomStringParameter `json:"customPrivateSubnetName,omitempty"`
	// EnableNoPublicIP - Should the Public IP be Disabled?
	EnableNoPublicIP *WorkspaceCustomBooleanParameter `json:"enableNoPublicIp,omitempty"`
	// LoadBalancerBackendPoolName - The name of a Backend Address Pool within an Azure Load Balancer
	LoadBalancerBackendPoolName *WorkspaceCustomStringParameter `json:"loadBalancerBackendPoolName,omitempty"`
	// LoadBalancerID - The Resource ID of an Azure Load Balancer
	LoadBalancerID *WorkspaceCustomStringParameter `json:"loadBalancerId,omitempty"`
	// RelayNamespaceName - The name of an Azure Relay Namespace
	RelayNamespaceName *WorkspaceCustomStringParameter `json:"relayNamespaceName,omitempty"`
	// StorageAccountName - The name which should be used for the Storage Account
	StorageAccountName *WorkspaceCustomStringParameter `json:"storageAccountName,omitempty"`
	// StorageAccountSkuName - The SKU which should be used for this Storage Account
	StorageAccountSkuName *WorkspaceCustomStringParameter `json:"storageAccountSkuName,omitempty"`
	// ResourceTags - A map of Tags which should be applied to the resources used by this Databricks Cluster.
	ResourceTags *WorkspaceCustomObjectParameter `json:"resourceTags,omitempty"`
	// VnetAddressPrefix - The first 2 octets of the virtual network /16 address range (e.g., '10.139' for the address range 10.139.0.0/16).
	VnetAddressPrefix *WorkspaceCustomStringParameter `json:"vnetAddressPrefix,omitempty"`
}

// WorkspaceCustomStringParameter the Value.
type WorkspaceCustomStringParameter struct {
	// Type - The type of variable that this is. Possible values include: 'Bool', 'Object', 'String'
	Type CustomParameterType `json:"type,omitempty"`
	// Value - The value which should be used for this field.
	Value *string `json:"value,omitempty"`
}

// WorkspaceListResult list of workspaces.
type WorkspaceListResult struct {
	autorest.Response `json:"-"`
	// Value - The array of workspaces.
	Value *[]Workspace `json:"value,omitempty"`
	// NextLink - The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// WorkspaceListResultIterator provides access to a complete listing of Workspace values.
type WorkspaceListResultIterator struct {
	i    int
	page WorkspaceListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *WorkspaceListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkspaceListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *WorkspaceListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter WorkspaceListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter WorkspaceListResultIterator) Response() WorkspaceListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter WorkspaceListResultIterator) Value() Workspace {
	if !iter.page.NotDone() {
		return Workspace{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the WorkspaceListResultIterator type.
func NewWorkspaceListResultIterator(page WorkspaceListResultPage) WorkspaceListResultIterator {
	return WorkspaceListResultIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (wlr WorkspaceListResult) IsEmpty() bool {
	return wlr.Value == nil || len(*wlr.Value) == 0
}

// workspaceListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (wlr WorkspaceListResult) workspaceListResultPreparer(ctx context.Context) (*http.Request, error) {
	if wlr.NextLink == nil || len(to.String(wlr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(wlr.NextLink)))
}

// WorkspaceListResultPage contains a page of Workspace values.
type WorkspaceListResultPage struct {
	fn  func(context.Context, WorkspaceListResult) (WorkspaceListResult, error)
	wlr WorkspaceListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *WorkspaceListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/WorkspaceListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.wlr)
	if err != nil {
		return err
	}
	page.wlr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *WorkspaceListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page WorkspaceListResultPage) NotDone() bool {
	return !page.wlr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page WorkspaceListResultPage) Response() WorkspaceListResult {
	return page.wlr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page WorkspaceListResultPage) Values() []Workspace {
	if page.wlr.IsEmpty() {
		return nil
	}
	return *page.wlr.Value
}

// Creates a new instance of the WorkspaceListResultPage type.
func NewWorkspaceListResultPage(getNextPage func(context.Context, WorkspaceListResult) (WorkspaceListResult, error)) WorkspaceListResultPage {
	return WorkspaceListResultPage{fn: getNextPage}
}

// WorkspaceProperties the workspace properties.
type WorkspaceProperties struct {
	// ManagedResourceGroupID - The managed resource group Id.
	ManagedResourceGroupID *string `json:"managedResourceGroupId,omitempty"`
	// Parameters - The workspace's custom parameters.
	Parameters *WorkspaceCustomParameters `json:"parameters,omitempty"`
	// ProvisioningState - READ-ONLY; The workspace provisioning state. Possible values include: 'Accepted', 'Running', 'Ready', 'Creating', 'Created', 'Deleting', 'Deleted', 'Canceled', 'Failed', 'Succeeded', 'Updating'
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
	// UIDefinitionURI - The blob URI where the UI definition file is located.
	UIDefinitionURI *string `json:"uiDefinitionUri,omitempty"`
	// Authorizations - The workspace provider authorizations.
	Authorizations *[]WorkspaceProviderAuthorization `json:"authorizations,omitempty"`
}

// WorkspaceProviderAuthorization the workspace provider authorization.
type WorkspaceProviderAuthorization struct {
	// PrincipalID - The provider's principal identifier. This is the identity that the provider will use to call ARM to manage the workspace resources.
	PrincipalID *uuid.UUID `json:"principalId,omitempty"`
	// RoleDefinitionID - The provider's role definition identifier. This role will define all the permissions that the provider must have on the workspace's container resource group. This role definition cannot have permission to delete the resource group.
	RoleDefinitionID *uuid.UUID `json:"roleDefinitionId,omitempty"`
}

// WorkspacesCreateOrUpdateFuture an abstraction for monitoring and retrieving the results of a
// long-running operation.
type WorkspacesCreateOrUpdateFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *WorkspacesCreateOrUpdateFuture) Result(client WorkspacesClient) (w Workspace, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databricks.WorkspacesCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("databricks.WorkspacesCreateOrUpdateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if w.Response.Response, err = future.GetResult(sender); err == nil && w.Response.Response.StatusCode != http.StatusNoContent {
		w, err = client.CreateOrUpdateResponder(w.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "databricks.WorkspacesCreateOrUpdateFuture", "Result", w.Response.Response, "Failure responding to request")
		}
	}
	return
}

// WorkspacesDeleteFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type WorkspacesDeleteFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *WorkspacesDeleteFuture) Result(client WorkspacesClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databricks.WorkspacesDeleteFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("databricks.WorkspacesDeleteFuture")
		return
	}
	ar.Response = future.Response()
	return
}

// WorkspacesUpdateFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type WorkspacesUpdateFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *WorkspacesUpdateFuture) Result(client WorkspacesClient) (w Workspace, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "databricks.WorkspacesUpdateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("databricks.WorkspacesUpdateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if w.Response.Response, err = future.GetResult(sender); err == nil && w.Response.Response.StatusCode != http.StatusNoContent {
		w, err = client.UpdateResponder(w.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "databricks.WorkspacesUpdateFuture", "Result", w.Response.Response, "Failure responding to request")
		}
	}
	return
}

// WorkspaceUpdate an update to a workspace.
type WorkspaceUpdate struct {
	// Tags - Resource tags.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for WorkspaceUpdate.
func (wu WorkspaceUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if wu.Tags != nil {
		objectMap["tags"] = wu.Tags
	}
	return json.Marshal(objectMap)
}
