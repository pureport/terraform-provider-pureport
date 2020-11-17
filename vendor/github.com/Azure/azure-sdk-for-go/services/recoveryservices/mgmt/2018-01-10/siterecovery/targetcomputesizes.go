package siterecovery

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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// TargetComputeSizesClient is the client for the TargetComputeSizes methods of the Siterecovery service.
type TargetComputeSizesClient struct {
	BaseClient
}

// NewTargetComputeSizesClient creates an instance of the TargetComputeSizesClient client.
func NewTargetComputeSizesClient(subscriptionID string, resourceGroupName string, resourceName string) TargetComputeSizesClient {
	return NewTargetComputeSizesClientWithBaseURI(DefaultBaseURI, subscriptionID, resourceGroupName, resourceName)
}

// NewTargetComputeSizesClientWithBaseURI creates an instance of the TargetComputeSizesClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewTargetComputeSizesClientWithBaseURI(baseURI string, subscriptionID string, resourceGroupName string, resourceName string) TargetComputeSizesClient {
	return TargetComputeSizesClient{NewWithBaseURI(baseURI, subscriptionID, resourceGroupName, resourceName)}
}

// ListByReplicationProtectedItems lists the available target compute sizes for a replication protected item.
// Parameters:
// fabricName - fabric name.
// protectionContainerName - protection container name.
// replicatedProtectedItemName - replication protected item name.
func (client TargetComputeSizesClient) ListByReplicationProtectedItems(ctx context.Context, fabricName string, protectionContainerName string, replicatedProtectedItemName string) (result TargetComputeSizeCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TargetComputeSizesClient.ListByReplicationProtectedItems")
		defer func() {
			sc := -1
			if result.tcsc.Response.Response != nil {
				sc = result.tcsc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByReplicationProtectedItemsNextResults
	req, err := client.ListByReplicationProtectedItemsPreparer(ctx, fabricName, protectionContainerName, replicatedProtectedItemName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.TargetComputeSizesClient", "ListByReplicationProtectedItems", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByReplicationProtectedItemsSender(req)
	if err != nil {
		result.tcsc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "siterecovery.TargetComputeSizesClient", "ListByReplicationProtectedItems", resp, "Failure sending request")
		return
	}

	result.tcsc, err = client.ListByReplicationProtectedItemsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.TargetComputeSizesClient", "ListByReplicationProtectedItems", resp, "Failure responding to request")
	}
	if result.tcsc.hasNextLink() && result.tcsc.IsEmpty() {
		err = result.NextWithContext(ctx)
	}

	return
}

// ListByReplicationProtectedItemsPreparer prepares the ListByReplicationProtectedItems request.
func (client TargetComputeSizesClient) ListByReplicationProtectedItemsPreparer(ctx context.Context, fabricName string, protectionContainerName string, replicatedProtectedItemName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"fabricName":                  autorest.Encode("path", fabricName),
		"protectionContainerName":     autorest.Encode("path", protectionContainerName),
		"replicatedProtectedItemName": autorest.Encode("path", replicatedProtectedItemName),
		"resourceGroupName":           autorest.Encode("path", client.ResourceGroupName),
		"resourceName":                autorest.Encode("path", client.ResourceName),
		"subscriptionId":              autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-10"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectedItems/{replicatedProtectedItemName}/targetComputeSizes", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByReplicationProtectedItemsSender sends the ListByReplicationProtectedItems request. The method will close the
// http.Response Body if it receives an error.
func (client TargetComputeSizesClient) ListByReplicationProtectedItemsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByReplicationProtectedItemsResponder handles the response to the ListByReplicationProtectedItems request. The method always
// closes the http.Response Body.
func (client TargetComputeSizesClient) ListByReplicationProtectedItemsResponder(resp *http.Response) (result TargetComputeSizeCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByReplicationProtectedItemsNextResults retrieves the next set of results, if any.
func (client TargetComputeSizesClient) listByReplicationProtectedItemsNextResults(ctx context.Context, lastResults TargetComputeSizeCollection) (result TargetComputeSizeCollection, err error) {
	req, err := lastResults.targetComputeSizeCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "siterecovery.TargetComputeSizesClient", "listByReplicationProtectedItemsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByReplicationProtectedItemsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "siterecovery.TargetComputeSizesClient", "listByReplicationProtectedItemsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByReplicationProtectedItemsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "siterecovery.TargetComputeSizesClient", "listByReplicationProtectedItemsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByReplicationProtectedItemsComplete enumerates all values, automatically crossing page boundaries as required.
func (client TargetComputeSizesClient) ListByReplicationProtectedItemsComplete(ctx context.Context, fabricName string, protectionContainerName string, replicatedProtectedItemName string) (result TargetComputeSizeCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TargetComputeSizesClient.ListByReplicationProtectedItems")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByReplicationProtectedItems(ctx, fabricName, protectionContainerName, replicatedProtectedItemName)
	return
}
