package documentdb

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// CollectionPartitionClient is the azure Cosmos DB Database Service Resource Provider REST API
type CollectionPartitionClient struct {
	BaseClient
}

// NewCollectionPartitionClient creates an instance of the CollectionPartitionClient client.
func NewCollectionPartitionClient(subscriptionID string) CollectionPartitionClient {
	return NewCollectionPartitionClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCollectionPartitionClientWithBaseURI creates an instance of the CollectionPartitionClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewCollectionPartitionClientWithBaseURI(baseURI string, subscriptionID string) CollectionPartitionClient {
	return CollectionPartitionClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListMetrics retrieves the metrics determined by the given filter for the given collection, split by partition.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// accountName - cosmos DB database account name.
// databaseRid - cosmos DB database rid.
// collectionRid - cosmos DB collection rid.
// filter - an OData filter expression that describes a subset of metrics to return. The parameters that can be
// filtered are name.value (name of the metric, can have an or of multiple names), startTime, endTime, and
// timeGrain. The supported operator is eq.
func (client CollectionPartitionClient) ListMetrics(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string, filter string) (result PartitionMetricListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CollectionPartitionClient.ListMetrics")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[a-z0-9]+(-[a-z0-9]+)*`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("documentdb.CollectionPartitionClient", "ListMetrics", err.Error())
	}

	req, err := client.ListMetricsPreparer(ctx, resourceGroupName, accountName, databaseRid, collectionRid, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.CollectionPartitionClient", "ListMetrics", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListMetricsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "documentdb.CollectionPartitionClient", "ListMetrics", resp, "Failure sending request")
		return
	}

	result, err = client.ListMetricsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.CollectionPartitionClient", "ListMetrics", resp, "Failure responding to request")
		return
	}

	return
}

// ListMetricsPreparer prepares the ListMetrics request.
func (client CollectionPartitionClient) ListMetricsPreparer(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"collectionRid":     autorest.Encode("path", collectionRid),
		"databaseRid":       autorest.Encode("path", databaseRid),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01"
	queryParameters := map[string]interface{}{
		"$filter":     autorest.Encode("query", filter),
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/collections/{collectionRid}/partitions/metrics", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListMetricsSender sends the ListMetrics request. The method will close the
// http.Response Body if it receives an error.
func (client CollectionPartitionClient) ListMetricsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListMetricsResponder handles the response to the ListMetrics request. The method always
// closes the http.Response Body.
func (client CollectionPartitionClient) ListMetricsResponder(resp *http.Response) (result PartitionMetricListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListUsages retrieves the usages (most recent storage data) for the given collection, split by partition.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// accountName - cosmos DB database account name.
// databaseRid - cosmos DB database rid.
// collectionRid - cosmos DB collection rid.
// filter - an OData filter expression that describes a subset of usages to return. The supported parameter is
// name.value (name of the metric, can have an or of multiple names).
func (client CollectionPartitionClient) ListUsages(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string, filter string) (result PartitionUsagesResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CollectionPartitionClient.ListUsages")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[a-z0-9]+(-[a-z0-9]+)*`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("documentdb.CollectionPartitionClient", "ListUsages", err.Error())
	}

	req, err := client.ListUsagesPreparer(ctx, resourceGroupName, accountName, databaseRid, collectionRid, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.CollectionPartitionClient", "ListUsages", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListUsagesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "documentdb.CollectionPartitionClient", "ListUsages", resp, "Failure sending request")
		return
	}

	result, err = client.ListUsagesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.CollectionPartitionClient", "ListUsages", resp, "Failure responding to request")
		return
	}

	return
}

// ListUsagesPreparer prepares the ListUsages request.
func (client CollectionPartitionClient) ListUsagesPreparer(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"collectionRid":     autorest.Encode("path", collectionRid),
		"databaseRid":       autorest.Encode("path", databaseRid),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/collections/{collectionRid}/partitions/usages", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListUsagesSender sends the ListUsages request. The method will close the
// http.Response Body if it receives an error.
func (client CollectionPartitionClient) ListUsagesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListUsagesResponder handles the response to the ListUsages request. The method always
// closes the http.Response Body.
func (client CollectionPartitionClient) ListUsagesResponder(resp *http.Response) (result PartitionUsagesResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
