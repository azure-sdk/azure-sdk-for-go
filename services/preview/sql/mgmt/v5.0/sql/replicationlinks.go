package sql

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
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

// ReplicationLinksClient is the the Azure SQL Database management API provides a RESTful set of web services that
// interact with Azure SQL Database services to manage your databases. The API enables you to create, retrieve, update,
// and delete databases.
type ReplicationLinksClient struct {
	BaseClient
}

// NewReplicationLinksClient creates an instance of the ReplicationLinksClient client.
func NewReplicationLinksClient(subscriptionID string) ReplicationLinksClient {
	return NewReplicationLinksClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewReplicationLinksClientWithBaseURI creates an instance of the ReplicationLinksClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewReplicationLinksClientWithBaseURI(baseURI string, subscriptionID string) ReplicationLinksClient {
	return ReplicationLinksClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Delete deletes the replication link.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// databaseName - the name of the database.
func (client ReplicationLinksClient) Delete(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationLinksClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, serverName, databaseName, linkID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ReplicationLinksClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "sql.ReplicationLinksClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ReplicationLinksClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ReplicationLinksClient) DeletePreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":      autorest.Encode("path", databaseName),
		"linkId":            autorest.Encode("path", linkID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/replicationLinks/{linkId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationLinksClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ReplicationLinksClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Failover fails over from the current primary server to this server.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// databaseName - the name of the database.
// linkID - the name of the replication link.
func (client ReplicationLinksClient) Failover(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string) (result ReplicationLinksFailoverFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationLinksClient.Failover")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.FailoverPreparer(ctx, resourceGroupName, serverName, databaseName, linkID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ReplicationLinksClient", "Failover", nil, "Failure preparing request")
		return
	}

	result, err = client.FailoverSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ReplicationLinksClient", "Failover", result.Response(), "Failure sending request")
		return
	}

	return
}

// FailoverPreparer prepares the Failover request.
func (client ReplicationLinksClient) FailoverPreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":      autorest.Encode("path", databaseName),
		"linkId":            autorest.Encode("path", linkID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/replicationLinks/{linkId}/failover", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// FailoverSender sends the Failover request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationLinksClient) FailoverSender(req *http.Request) (future ReplicationLinksFailoverFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// FailoverResponder handles the response to the Failover request. The method always
// closes the http.Response Body.
func (client ReplicationLinksClient) FailoverResponder(resp *http.Response) (result ReplicationLink, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// FailoverAllowDataLoss fails over from the current primary server to this server allowing data loss.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// databaseName - the name of the database.
// linkID - the name of the replication link.
func (client ReplicationLinksClient) FailoverAllowDataLoss(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string) (result ReplicationLinksFailoverAllowDataLossFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationLinksClient.FailoverAllowDataLoss")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.FailoverAllowDataLossPreparer(ctx, resourceGroupName, serverName, databaseName, linkID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ReplicationLinksClient", "FailoverAllowDataLoss", nil, "Failure preparing request")
		return
	}

	result, err = client.FailoverAllowDataLossSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ReplicationLinksClient", "FailoverAllowDataLoss", result.Response(), "Failure sending request")
		return
	}

	return
}

// FailoverAllowDataLossPreparer prepares the FailoverAllowDataLoss request.
func (client ReplicationLinksClient) FailoverAllowDataLossPreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":      autorest.Encode("path", databaseName),
		"linkId":            autorest.Encode("path", linkID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/replicationLinks/{linkId}/forceFailoverAllowDataLoss", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// FailoverAllowDataLossSender sends the FailoverAllowDataLoss request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationLinksClient) FailoverAllowDataLossSender(req *http.Request) (future ReplicationLinksFailoverAllowDataLossFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// FailoverAllowDataLossResponder handles the response to the FailoverAllowDataLoss request. The method always
// closes the http.Response Body.
func (client ReplicationLinksClient) FailoverAllowDataLossResponder(resp *http.Response) (result ReplicationLink, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets a replication link.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// databaseName - the name of the database.
// linkID - the name of the replication link.
func (client ReplicationLinksClient) Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string) (result ReplicationLink, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationLinksClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, serverName, databaseName, linkID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ReplicationLinksClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ReplicationLinksClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ReplicationLinksClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ReplicationLinksClient) GetPreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":      autorest.Encode("path", databaseName),
		"linkId":            autorest.Encode("path", linkID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/replicationLinks/{linkId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationLinksClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ReplicationLinksClient) GetResponder(resp *http.Response) (result ReplicationLink, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByDatabase gets a list of replication links on database.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
// databaseName - the name of the database.
func (client ReplicationLinksClient) ListByDatabase(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result ReplicationLinkListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationLinksClient.ListByDatabase")
		defer func() {
			sc := -1
			if result.rllr.Response.Response != nil {
				sc = result.rllr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByDatabaseNextResults
	req, err := client.ListByDatabasePreparer(ctx, resourceGroupName, serverName, databaseName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ReplicationLinksClient", "ListByDatabase", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByDatabaseSender(req)
	if err != nil {
		result.rllr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ReplicationLinksClient", "ListByDatabase", resp, "Failure sending request")
		return
	}

	result.rllr, err = client.ListByDatabaseResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ReplicationLinksClient", "ListByDatabase", resp, "Failure responding to request")
		return
	}
	if result.rllr.hasNextLink() && result.rllr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByDatabasePreparer prepares the ListByDatabase request.
func (client ReplicationLinksClient) ListByDatabasePreparer(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/replicationLinks", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByDatabaseSender sends the ListByDatabase request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationLinksClient) ListByDatabaseSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByDatabaseResponder handles the response to the ListByDatabase request. The method always
// closes the http.Response Body.
func (client ReplicationLinksClient) ListByDatabaseResponder(resp *http.Response) (result ReplicationLinkListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByDatabaseNextResults retrieves the next set of results, if any.
func (client ReplicationLinksClient) listByDatabaseNextResults(ctx context.Context, lastResults ReplicationLinkListResult) (result ReplicationLinkListResult, err error) {
	req, err := lastResults.replicationLinkListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.ReplicationLinksClient", "listByDatabaseNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByDatabaseSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.ReplicationLinksClient", "listByDatabaseNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByDatabaseResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ReplicationLinksClient", "listByDatabaseNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByDatabaseComplete enumerates all values, automatically crossing page boundaries as required.
func (client ReplicationLinksClient) ListByDatabaseComplete(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result ReplicationLinkListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationLinksClient.ListByDatabase")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByDatabase(ctx, resourceGroupName, serverName, databaseName)
	return
}

// ListByServer gets a list of replication links.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// serverName - the name of the server.
func (client ReplicationLinksClient) ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result ReplicationLinkListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationLinksClient.ListByServer")
		defer func() {
			sc := -1
			if result.rllr.Response.Response != nil {
				sc = result.rllr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByServerNextResults
	req, err := client.ListByServerPreparer(ctx, resourceGroupName, serverName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ReplicationLinksClient", "ListByServer", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByServerSender(req)
	if err != nil {
		result.rllr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "sql.ReplicationLinksClient", "ListByServer", resp, "Failure sending request")
		return
	}

	result.rllr, err = client.ListByServerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ReplicationLinksClient", "ListByServer", resp, "Failure responding to request")
		return
	}
	if result.rllr.hasNextLink() && result.rllr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByServerPreparer prepares the ListByServer request.
func (client ReplicationLinksClient) ListByServerPreparer(ctx context.Context, resourceGroupName string, serverName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serverName":        autorest.Encode("path", serverName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/replicationLinks", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByServerSender sends the ListByServer request. The method will close the
// http.Response Body if it receives an error.
func (client ReplicationLinksClient) ListByServerSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByServerResponder handles the response to the ListByServer request. The method always
// closes the http.Response Body.
func (client ReplicationLinksClient) ListByServerResponder(resp *http.Response) (result ReplicationLinkListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByServerNextResults retrieves the next set of results, if any.
func (client ReplicationLinksClient) listByServerNextResults(ctx context.Context, lastResults ReplicationLinkListResult) (result ReplicationLinkListResult, err error) {
	req, err := lastResults.replicationLinkListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "sql.ReplicationLinksClient", "listByServerNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByServerSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "sql.ReplicationLinksClient", "listByServerNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByServerResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sql.ReplicationLinksClient", "listByServerNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByServerComplete enumerates all values, automatically crossing page boundaries as required.
func (client ReplicationLinksClient) ListByServerComplete(ctx context.Context, resourceGroupName string, serverName string) (result ReplicationLinkListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ReplicationLinksClient.ListByServer")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByServer(ctx, resourceGroupName, serverName)
	return
}
